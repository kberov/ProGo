package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func init() {
	http.HandleFunc("/cookie",
		func(writer http.ResponseWriter, request *http.Request) {
			counterVal := 1
			counterCookie, err := request.Cookie("counter")
			if err == nil {
				counterVal, _ = strconv.Atoi(counterCookie.Value)
				counterVal++
			}
			http.SetCookie(writer, &http.Cookie{
				Name: "counter", Value: strconv.Itoa(counterVal),
			})
			// 679 Chapter 25 ■ Creating HTTP Clients
			if len(request.Cookies()) > 0 {
				for _, c := range request.Cookies() {
					fmt.Fprintf(writer, "Cookie Name: %v, Value: %v\n",
						c.Name, c.Value)
				}
			} else {
				fmt.Fprintln(writer, "Request contains no cookies")
			}
		})
}
