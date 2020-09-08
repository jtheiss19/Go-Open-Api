package main

import "net/http"

func myHandler(headers http.Header) {
	for k, v := range headers {
		if k == "Status" {
			for _, value := range v {
				if value == "stop" {
					stop()
				}
			}
		}
	}
}
