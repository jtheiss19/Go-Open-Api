package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/jtheiss19/Go-Open-Api/communication"
)

var server http.Server
var power bool = false
var logger *log.Logger

var port = "8080"

func status() bool {
	return power
}

func stop() {
	if power {
		logger.Output(0, "User Has Chose to Shutdown")
		logger.Output(0, "Shutting Down...")
		server.Close()
		power = false
		logger.Output(0, "...Shutdown")
		os.Exit(3)
	}
}

func main() {
	power = true
	f, _ := os.Create("./logs/website_log")
	logger = log.New(f, "Website ", log.LstdFlags)

	mux := http.NewServeMux()
	mux.Handle("/public/", logging(public()))
	mux.Handle("/", logging(index()))

	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	logger.Output(0, "main: running simple server on port: "+port)

	go communication.Host(myHandler)

	if err := server.ListenAndServe(); err != nil {
		logger.Fatalf("main: couldn't start simple server: %v\n", err)
	}
}

// logging is middleware for wrapping any handler we want to track response
// times for and to see what resources are requested.
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		logger.Output(0, req)
		next.ServeHTTP(w, r)
		logger.Output(0, req+" completed in "+time.Now().Sub(start).String())
	})
}

// templates references the specified templates and caches the parsed results
// to help speed up response times.
var templates = template.Must(template.ParseFiles("./templates/base.html", "./templates/body.html"))

// index is the handler responsible for rending the index page for the site.
func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := struct {
			Title        template.HTML
			BusinessName string
			Slogan       string
		}{
			Title:        template.HTML("Business &verbar; Landing"),
			BusinessName: "Business,",
			Slogan:       "we get things done.",
		}
		err := templates.ExecuteTemplate(w, "base", &b)
		if err != nil {
			http.Error(w, fmt.Sprintf("index: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
	})
}

func testIndex(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("TestIndex: couldn't create HTTP GET request: %v", err)
	}

	rec := httptest.NewRecorder()

	index().ServeHTTP(rec, req)

	res := rec.Result()
	defer func() {
		err := res.Body.Close()
		if err != nil {
			t.Fatalf("TestIndex: couldn't close response body: %v", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		t.Errorf("TestIndex: got status code %v, but want: %v", res.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("TestIndex: could not read response body: %v", err)
	}

	if len(string(body)) == 0 {
		t.Errorf("TestIndex: unexpected empty response body")
	}
}

func public() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
}
