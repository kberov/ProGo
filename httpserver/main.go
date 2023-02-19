package main

import (
	. "fmt"
	"io"
	"net/http"
	"net/http/cgi"
	"os"
	"strings"
)

type StringHandler struct {
	message string
}

var port = 5000

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	Printfln("Method: %#v", request.Method)
	Printfln("URL: %#v", request.URL)
	Printfln("HTTP Version: %#v", request.Proto)
	Printfln("Host: %#v", request.Host)
	Printfln("RemoteAddr: %#v", request.RemoteAddr)
	for name, val := range request.Header {
		Printfln("Header: %v, Value: %#v", name, val)
	}
	Printfln("---")
	Printfln("Response headers...")
	Printfln("TODO---")
	Printfln("Request for %v", request.URL.Path)

	Printfln("Request for %v", request.URL.Path)
	n, e := io.WriteString(writer, sh.message+"\n")
	if e != nil {
		println("Error writing string, long", n, "bytes with writer:", e.Error())
	} else {
		println("Wrote ", n, "bytes to client", request.UserAgent())
	}
	//writer.WriteHeader(http.StatusForbidden) //404
}

func HTTPSRedirect(writer http.ResponseWriter,
	request *http.Request) {
	host_port := strings.Split(request.Host, ":")[0:2]
	Printfln("host_port: %v", host_port)
	target := "https://" + host_port[0] + ":5500" + request.URL.Path
	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}
	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
}

func main() {
	Printfln("\nCreating a Simple HTTP Server")
	Printfln("%s%s%s%s",
		"\n Creating a Static HTTP Server",
		"\n Using Templates to Generate Responses",
		"\n Reading Form Data from Requests",
		"\n Reading and Setting Cookies")
	/*
		for _, p := range Products {
			Printfln("Product: %v, Category: %v, Price: $%.2f",
				p.Name, p.Category, p.Price)
		}
	*/
	http.Handle("/message", StringHandler{"Hello, World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))
	/*
		NOTE!!!
		cgi.Serve() would accept and serve the same type as
		http.ListenAndServe() would do. It does so by invoking our
		(sh StringHandler)ServeHTTP(w,r). See also ServeHTTP(w,r)
		defined in go/src/net/http/cgi/host.go
		We just have to detect if the program is executed in a CGI environment.
		TODO: MAY BE write a CGIServer (and FCGIServer, and SCGIServer) extension for
		gobuffallo or other framework to detect CGI and handle the request.
		Publish it to contribute for glory and make ICDSoft accept my proposal
		to support one of these protocols to make other technologies also first
		class citizens beside PHP on their Shared Hosting servers... See:
		https://en.wikipedia.org/wiki/Simple_Common_Gateway_Interface
		http://python.ca/scgi/protocol.txt
		https://en.wikipedia.org/wiki/Common_Gateway_Interface
		https://www.rfc-editor.org/rfc/rfc3875
		NOTE!!!! 2023-02-06
		cgi.Serve(nil)
		works just like
		http.ListenAndServeTLS(Sprintf(":%d", port), "localhost.crt",
					"localhost.key", nil)
					and
		http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
		It serves the url defined in PATH_INFO environment variable depending
		on the defined url handlers via http.Handle(...)
		Templates handling is defined in dynamic.go JSON responses - in json.go
	*/
	if len(os.Args) > 1 && strings.Contains(os.Args[1], "cgi") {
		Printfln("will serve via CGI with args:%#v-------------------", os.Args)
		cgiErr := cgi.Serve(nil)
		if cgiErr != nil {
			Printfln("Error: %v", cgiErr.Error())
		}
		println("served via CGI!-------------------\n\n")
		os.Exit(0)
	}
	Println("Creating the HTTP Listener and Handler")
	go func() {
		port = 5500
		Printfln("will serve on localhost:%d", port)
		err := http.ListenAndServeTLS(Sprintf(":%d", port), "localhost.crt",
			"localhost.key", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()
	Printfln("will serve on localhost:%d", port)
	err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
	if err != nil {
		Printfln("Error: %v", err.Error())
	}

}
