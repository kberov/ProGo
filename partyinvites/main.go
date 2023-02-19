package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var responses = make([]*Rsvp, 10)
var templates = make(map[string]*template.Template, 3)

func main() {
	loadTemplates()
	http.HandleFunc("/", welcomeH)
	http.HandleFunc("/list", listH)
	http.HandleFunc("/form", formH)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func loadTemplates() {
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", 1+index, name)
		} else {
			panic(err)
		}
	}
}

// Model
type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

type formData struct {
	*Rsvp
	Errors []string
}

// Controller

// The second argument is a pointer to an instance of the Request struct, defined
// in the net/http package, which describes the request being processed. The first
// argument is an example of an interface, which is why it isn’t defined as a
// pointer.
// handler functions - actions
func welcomeH(res http.ResponseWriter, req *http.Request) {
	templates["welcome"].Execute(res, nil)
}

func listH(res http.ResponseWriter, req *http.Request) {
	templates["list"].Execute(res, responses)
}

func formH(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		/* The second parameter creates a formData struct by creating a new
		* instance of the Rsvp struct and creating a string slice that contains
		* no values. The ampersand (the & character) creates a pointer to a
		* value. The formData struct has been defined to expect a pointer to an
		* Rsvp value, which the ampersand allows me to create. */
		templates["form"].Execute(res, formData{Rsvp: &Rsvp{}, Errors: []string{}})
		return
	}

	if req.Method == http.MethodPost {
		req.ParseForm()
		// a struct is instantiated with values for its fields, as opposed to
		// the default values that were used in Listing 1-19
		resData := Rsvp{
			// HTML forms can include multiple values with the same name, so
			// the form data is presented as a slice of values.
			Name:       req.Form["name"][0],
			Email:      req.Form["email"][0],
			Phone:      req.Form["phone"][0],
			WillAttend: req.Form["willattend"][0] == "true",
		}
		errors := validate(&resData)
		if len(errors) > 0 {
			templates["form"].Execute(res, formData{Rsvp: &resData, Errors: errors})
			return
		}

		// The append function is used to append a value to a slice. Notice
		// that I use the ampersand to create a pointer to the Rsvp value I
		// created. If I had not used a pointer, then my Rsvp value would be
		// duplicated when it is added to the slice.
		responses = append(responses, &resData)
		if resData.WillAttend {
			templates["thanks"].Execute(res, resData.Name)
		} else {
			templates["sorry"].Execute(res, resData.Name)

		}
	}
}

func validate(resData *Rsvp) []string {
	errors := []string{}
	if resData.Name == "" {
		errors = append(errors, "Please enter your name!")
	}
	if resData.Email == "" {
		errors = append(errors, "Please enter your email!")
	}
	if resData.Phone == "" {
		errors = append(errors, "Please enter your phone!")
	}
	return errors
}
