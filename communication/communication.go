package communication

import (
	"log"
	"net/http"
)

var comHandler func(http.Header)

func Host(myHandler func(http.Header)) {
	comHandler = myHandler
	SendReq()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	comHandler(r.Header)
}

func SendReq() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8000", nil)

	req.Header.Set("status", "stop")

	client.Do(req)

}
