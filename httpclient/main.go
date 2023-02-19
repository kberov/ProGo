package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	Printfln("Starting HTTP Server")
	go func() { _ = http.ListenAndServe(":5000", nil) }()
	time.Sleep(time.Second)

	get_url := "http://localhost:5000/ХТМЛ"
	Printfln("\nMaking a GET request to %s\n", get_url)

	response, err := http.Get(get_url)
	if err == nil {
		//	_ = response.Write(os.Stdout)
	} else {
		Printfln("Error: %v", err.Error())
	}

	Printfln("But a realistic scenario would be the following code:")

	if err == nil && response.StatusCode == http.StatusOK {

		data, err := io.ReadAll(response.Body)
		if err == nil {
			defer func() { _ = response.Body.Close() }()
			Printfln("Status: %v %v\n\nHeaders:", response.Proto, response.Status)
			for k, v := range response.Header {
				Printfln("%v: %v", k, v)
			}
			Printfln("Body:")
			_, _ = os.Stdout.Write(data)
		}
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	}

	get_url = "http://localhost:5000/json"
	Printfln("\nMaking a GET request to %s\n", get_url)
	response, err = http.Get("http://localhost:5000/json")
	if err == nil && response.StatusCode == http.StatusOK {
		defer func() { _ = response.Body.Close() }()
		data := []Product{}
		body, _ := io.ReadAll(response.Body)
		// we need to convert the []byte slice to a string, then return a new reader for NewDecoder
		err = json.NewDecoder(strings.NewReader(string(body))).Decode(&data)
		if err == nil {
			Printfln("JSON: %s\nDecoded data:", string(body))
			for _, p := range data {
				Printfln("Name: %v, Price: $%.2f", p.Name, p.Price)
			}
		} else {
			Printfln("Decode error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	}

	Printfln("\nSending POST Requests")

	formData := map[string][]string{
		"name":     {"Kayak "},
		"category": {"Watersports"},
		"price":    {"279"},
	}
	post_url := "http://localhost:5000/echo"
	Printfln("Sending a POST request with form data: %#v\n to url %s", formData, post_url)
	response, err = http.PostForm(post_url, formData)
	if err == nil && response.StatusCode == http.StatusOK {
		Printfln("response.Body:")
		_, _ = io.Copy(os.Stdout, response.Body)
		defer func() { _ = response.Body.Close() }()
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	}

	Printfln("\n\nPosting a Form Using a Reader")
	/* The Post function sends a POST request to the server and creates the
	* request body by reading content from a Reader, as shown in Listing 25-12.
	* Unlike the PostForm function, the data doesn’t have to be encoded as a
	* form. */
	var builder strings.Builder
	err = json.NewEncoder(&builder).Encode(Products[0])
	if err == nil {
		response, err := http.Post("http://localhost:5000/echo",
			"application/json",
			strings.NewReader(builder.String()))
		if err == nil && response.StatusCode == http.StatusOK {
			_, _ = io.Copy(os.Stdout, response.Body)
			defer func() { _ = response.Body.Close() }()
		} else {
			Printfln("Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}

	Printfln("\nConfiguring HTTP Client Requests")
	// reset the strings Builder to be empty and reuse it below.
	builder.Reset()
	err = json.NewEncoder(&builder).Encode(Products[0])
	if err == nil {
		reqURL, err := url.Parse("http://localhost:5000/echo")
		if err == nil {
			req := http.Request{
				Method: http.MethodPost,
				URL:    reqURL,
				Header: map[string][]string{
					"Content-Type":    {"application.json"},
					"X-MyHeader":      {"1"},
					"Accept-Encoding": {"deflate"},
				},
				Body: io.NopCloser(strings.NewReader(builder.String())),
			}
			response, err := http.DefaultClient.Do(&req)
			if err == nil && response.StatusCode == http.StatusOK {
				_, _ = io.Copy(os.Stdout, response.Body)
				defer func() { _ = response.Body.Close() }()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
		} else {
			Printfln("Parse Error: %v", err.Error())
		}
	} else {
		Printfln("Encoder Error: %v", err.Error())
	}

	Printfln("\nUsing the Convenience Functions to Create a Request")

	// url is automatically parsed. req.Header is initialized
	req, err := http.NewRequest(http.MethodPost, post_url,
		io.NopCloser(strings.NewReader(builder.String())))
	if err == nil {
		req.Header["Content-Type"] = []string{"application/json"}
		response, err := http.DefaultClient.Do(req)
		if err == nil && response.StatusCode == http.StatusOK {
			_, _ = io.Copy(os.Stdout, response.Body)
			defer func() { _ = response.Body.Close() }()
		} else {
			Printfln("Request Error: %v", err.Error())
		}
		// 678 Chapter 25 ■ Creating HTTP Clients
	} else {
		Printfln("Encoder Error: %v", err.Error())
	}

	Printfln("\nWorking with Cookies")
	// prepare a jar for cookies in the client so it manages cookies automatically
	jar, err := cookiejar.New(nil)
	if err == nil {
		http.DefaultClient.Jar = jar
	}
	for i := 0; i < 3; i++ {
		req, err := http.NewRequest(http.MethodGet,
			"http://localhost:5000/cookie", nil)
		if err == nil {
			response, err := http.DefaultClient.Do(req)
			if err == nil && response.StatusCode == http.StatusOK {
				Printfln("response.Header:\n%v", response.Header)
				_, _ = io.Copy(os.Stdout, response.Body)
				//defer func() { _ = response.Body.Close() }()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
			// 680 Chapter 25 ■ Creating HTTP Clients
		} else {
			Printfln("Request Init Error: %v", err.Error())
		}
	}

	Printfln("\nCreating Separate Clients and Cookie Jars")
	/* If multiple Client values are required but cookies should be shared,
	* then a single CookieJar can be used, as shown in Listing 25-18. */
	jar, err = cookiejar.New(nil)
	clients := make([]http.Client, 3)
	for index, client := range clients {
		//jar, err := cookiejar.New(nil)
		if err == nil {
			client.Jar = jar
		}
		for i := 0; i < 3; i++ {
			req, err := http.NewRequest(http.MethodGet,
				"http://localhost:5000/cookie", nil)
			if err == nil {
				response, err := client.Do(req)
				if err == nil && response.StatusCode == http.StatusOK {
					fmt.Fprintf(os.Stdout, "Client %v: ", index)
					_, _ = io.Copy(os.Stdout, response.Body)
					// defer response.Body.Close()
				} else {
					Printfln("Request Error: %v", err.Error())
				}
			} else {
				Printfln("Request Init Error: %v", err.Error())
			}
		}
	}

	Printfln("\nManaging Redirections\n go to index.html after 3 redirections")
	// A custom policy is defined by assigning a function to the
	// Client.CheckRedirect field
	http.DefaultClient.CheckRedirect = func(req *http.Request,
		previous []*http.Request) error {
		// Custom policy: go to /ХТМЛ if we were redirected already 3 times
		println("...previous rediections:", len(previous))
		if len(previous) == 3 {
			url, _ := url.Parse("http://localhost:5000/ХТМЛ")
			req.URL = url
		}
		return nil
	}
	req, err = http.NewRequest(http.MethodGet,
		"http://localhost:5000/redirect1", nil)
	if err == nil {
		var response *http.Response
		response, err = http.DefaultClient.Do(req)
		if err == nil {
			_, _ = io.Copy(os.Stdout, response.Body)
		} else {
			Printfln("Request Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}

	Printfln("\nCreating Multipart Forms")

	// A specific sequence is required to create a form. First, call the
	// NewWriter function to get a multipart.Writer:
	var buffer bytes.Buffer
	formWriter := multipart.NewWriter(&buffer)

	/* A Reader is required to use the form data as the body for an HTTP
	* request, but a Writer is required to create the form. This is an ideal
	* situation for a bytes.Buffer struct, which provides an in-memory
	* implementation of both the Reader and Writer interfaces. Once the
	* multipart.Writer has been created, the CreateFormField and CreateFormFile
	* methods are used to add fields and files to the form: */
	fieldWriter, err := formWriter.CreateFormField("name")
	if err == nil {
		_, _ = io.WriteString(fieldWriter, "Alice")
	}
	fieldWriter, err = formWriter.CreateFormField("city")
	if err == nil {
		_, _ = io.WriteString(fieldWriter, "New York")
	}
	fileWriter, err := formWriter.CreateFormFile("codeFile", "printer.go")
	if err == nil {
		fileData, err := os.ReadFile("./printer.go")
		if err == nil {
			_, _ = fileWriter.Write(fileData)
		}
	}
	// Both of these methods return a Writer that is used to write content to
	// the form. See above.

	/* Caution: Don’t use the defer keyword on the call to the Close method;
	 * otherwise, the final boundary string won’t be added to the form until after
	 * the request will be sent, producing a form that not all servers will process.
	 * It is important to call the Close method before sending the request. */
	_ = formWriter.Close() // We are done with the form. Close it.
	req, err = http.NewRequest(http.MethodPost,
		"http://localhost:5000/form", &buffer)

	/* Once the fields and files have been added, the next step is to set the
	* Content-Type header, using the result from the FormDataContentType
	* method: */
	req.Header["Content-Type"] = []string{formWriter.FormDataContentType()}
	println("form content type and parts bundary:", formWriter.FormDataContentType())
	// 689 Chapter 25 ■ Creating HTTP Clients
	if err == nil {
		var response *http.Response
		response, err = http.DefaultClient.Do(req)
		if err == nil {
			_, _ = io.Copy(os.Stdout, response.Body)
		} else {
			Printfln("Request Error: %v", err.Error())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}

}
