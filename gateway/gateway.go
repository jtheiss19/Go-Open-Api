package gateway

import (
	"log"
	"net"
	"os"
	"time"
)

var timeoutDuration = 10 * time.Second

var power bool = false
var backendServers map[string]string = make(map[string]string)
var shutdownchan chan string = make(chan string)

//Stop is the function that should be called to properly close the gateway
func Stop() {
	shutdownchan <- "User Chose To Shutdown"
}

//Start begins the hosting process for the
//client to server application
func Start(port string) {
	f, _ := os.Create("./gateway_log")
	log.SetOutput(f)

	log.Output(0, "Launching Software Defined Network...")

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Output(0, err.Error())
		shutdownchan <- "Could Not Listen on Port" + ":" + port
		return
	}

	log.Output(0, "Online - Now Listening On Port: "+port)
	power = true
	go shutdownLoop()

	ConnSignal := make(chan string)

	for power {
		go session(ln, ConnSignal, port)
		log.Output(0, <-ConnSignal)
	}

	log.Output(0, "...Shut Down")
}

func shutdownLoop() {
	log.Output(0, <-shutdownchan)
	log.Output(0, "Shutting Down Now...")
	power = false
}

//session creates a new seesion listening on a port. This
//session handles all interactions with the connected
//client
func session(ln net.Listener, ConnSignal chan string, port string) {
	conn, err := ln.Accept()
	if err != nil {
		ConnSignal <- "Could not establish a connection with client"
		conn.Close()
		return
	}
	defer conn.Close()

	ConnSignal <- "New Connection \n"

	var serverConn net.Conn = nil

	//Open a new connection to the server
	serverConn, err = net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Output(0, "Signal could not be relized")
		return
	}

	shutdownSession := make(chan string)
	go sessionListener(conn, serverConn) // For the incomming client message
	go sessionListener(serverConn, conn) // For the Incomming server message

	<-shutdownSession
}

//SessionListener listens for connections noise and sends it to the writer
func sessionListener(InuputConnection net.Conn, OutputConnection net.Conn) {
	var temp []byte
	for {
		buf := make([]byte, 1024)

		//Used for the timeout function incase of bad internet or extrenious wait times.
		InuputConnection.SetReadDeadline(time.Now().Add(timeoutDuration))

		//Read without error from the inputconnection
		bytes, err := InuputConnection.Read(buf)
		if err != nil {
			log.Output(0, "read Error")
			break
		}

		//If we have the full message we can now send it
		if bytes != 1024 {
			log.Output(0, "Full Message Recieved")
			buf = buf[0 : bytes+1] //Allows for the endcharacter to be captured, not outputed as a byte count number normally
			temp = append(temp, buf...)
			break
		}

		//If we are still adding to the message, we append and try and read more from buffer.
		temp = append(temp, buf...)

	}
	log.Output(0, string(temp))
	OutputConnection.Write(temp)
}
