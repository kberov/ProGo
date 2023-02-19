package main

import (
	"fmt"
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/form",
		func(writer http.ResponseWriter, request *http.Request) {
			err := request.ParseMultipartForm(90 * 1024 * 1024) // 90MB - maxMemory bytes
			if err == nil {
				for name, vals := range request.MultipartForm.Value {
					fmt.Fprintf(writer, "Field %v: %v\n", name, vals)
				}
				for name, files := range request.MultipartForm.File {
					for _, file := range files {
						// 687 Chapter 25 ■ Creating HTTP Clients
						fmt.Fprintf(writer, "File %v: %v\n", name, file.Filename)
						if f, err := file.Open(); err == nil {
							defer func() { _ = f.Close() }()
							_, _ = io.Copy(writer, f)
						}
					}
				}
			} else {
				fmt.Fprintf(writer, "Cannot parse form %v", err.Error())
			}
		})
}
