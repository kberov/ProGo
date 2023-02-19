package main

import "net/http"

func init() {
	// 684 Chapter 25 ■ Creating HTTP Clients
	http.HandleFunc("/redirect1",
		func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, "/redirect2",
				http.StatusTemporaryRedirect)
		})
	http.HandleFunc("/redirect2",
		func(writer http.ResponseWriter, request *http.Request) {
			http.Redirect(writer, request, "/redirect1",
				http.StatusTemporaryRedirect)
		})
}
