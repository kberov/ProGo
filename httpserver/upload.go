package main

import (
	"fmt"
	"io"
	"net/http"
)

func HandleMultipartForm(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseMultipartForm(10000000)
	fmt.Fprintf(writer, "Name: %v, City: %v\n",
		request.MultipartForm.Value["name"][0],
		request.MultipartForm.Value["city"][0])
	// 657 Chapter 24 ■ Creating HTTP Servers
	fmt.Fprintln(writer, "------")
	for _, header := range request.MultipartForm.File["files"] {
		fmt.Fprintf(writer, "Name: %v, Size: %v\n", header.Filename, header.Size)
		file, err := header.Open()
		if err == nil {
			defer func() { _ = file.Close() }()
			fmt.Fprintln(writer, "------")
			_, _ = io.Copy(writer, file)
		} else {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func init() {
	http.HandleFunc("/forms/upload", HandleMultipartForm)
}
