package main

import (
	// . "fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func GetCategories(products []Product) (categories []template.HTML) {
	catMap := map[string]string{}
	for _, p := range products {
		if catMap[p.Category] == "" {
			catMap[p.Category] = p.Category
			// println(p.Category)
			/* Unlike other languages, Go doesn’t support implicit type
			conversion. Although when dividing numbers, implicit conversions
			happen depending on the scenario. So we need to be very careful
			about what type to use where.
			So here we had to cast the string value to template.HTML allthough
			it is just an alias for string.
			*/
			categories = append(categories, template.HTML("<p>"+p.Category+"</p>"))
		}
	}
	return
}
func main() {

	var templateFuncs = map[string]any{
		"getCats": GetCategories,
		"lower": func(s template.HTML) template.HTML {
			return template.HTML(strings.ToLower(string(s)))
		},
	}
	Printfln("\nUsing HTML and Text Templates")

	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: $%.2f",
			p.Name, p.Category, p.Price)
	}

	t, err := template.ParseFiles("templates/first_template.html")
	if err == nil {
		// stream output to STDOUT
		er := t.Execute(os.Stdout, &Kayak)
		if er != nil {
			Printfln("Template error: %s", er.Error())
		}
		// 599Chapter 23 ■ Using HTML and Text Templates
	} else {
		Printfln("Error: %v", err.Error())
	}
	Printfln("\nLoading Multiple Templates")
	_, _ = os.Stdout.WriteString("--------------\n")

	t2, _ := template.ParseFiles("templates/extras.html")
	er := t2.Execute(os.Stdout, &Products[1])
	if er != nil {
		Printfln("Error: %v", er.Error())
	}
	_, _ = os.Stdout.WriteString("-----Glob templates/*html---\n")

	templateFiles, _ := filepath.Glob(filepath.Join("templates", "*.html"))
	all := template.New("all")
	all.Funcs(templateFuncs)
	all, err = all.ParseFiles(templateFiles...)
	if err == nil {
		for i, t := range templateFiles {
			t = strings.Replace(t, "templates/", "", 1)
			Printfln("---- %s", t)
			er = all.ExecuteTemplate(os.Stdout, t, &Products[i])
			if er != nil {
				Printfln("Error: %v", er.Error())
			}

		}
	} else {
		Printfln("Error: %v", err.Error())
	}
	Printfln("Enumerating Loaded Templates")
	allTemplates := template.New("allTemplates")
	allTemplates.Funcs(templateFuncs)
	// we could use ParseGlob above as we do here
	allTemplates, err = allTemplates.ParseGlob("templates/*.html")
	if err == nil {
		for _, t := range allTemplates.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error:  %v", err.Error())
		os.Exit(0)
	}

	Printfln("Looking Up a Specific Template")
	selectedTemplate := allTemplates.Lookup("template.html")
	err = Exec(selectedTemplate, &Products[3])
	if err != nil {
		Printfln("Error: %v", err.Error())
	}

	Printfln("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s",
		"\nUnderstanding Template Actions",
		"\n Inserting Data Values",
		"\n Formatting Data Values",
		"\n Chaining and Parenthesizing Template Expressions",
		"\n Trimming Whitespace",
		"\n Using Slices in Templates",
		"\n Using the Built-in Slice Functions",
		"\n Conditionally Executing Template Content",
		"\n Using the Optional Conditional Actions",
		"\n Creating Named Nested Templates",
		"\n Defining Template Blocks",
		"\n Defining Template Functions",
		"\n Disabling Function Result Encoding",
		"\n Defining Template Variables",
		"\n")
	// I renamed range.html to _range.html so it is always loaded first and
	// list.html can redefine the `body` block.
	// all, err = template.ParseFiles("templates/range.html", "templates/list.html")
	// range action
	e := all.Lookup("mainRange").Execute(os.Stdout, Products)
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
	if e != nil {
		Printfln("Error: %v", e.Error())
	}
	_, _ = os.Stdout.WriteString("----- calling defined function getCats---\n")

	e = all.Lookup("mainTemplate").Execute(os.Stdout, Products)
	if e != nil {
		Printfln("Error: %v", e.Error())
	}

}

func Exec(t *template.Template, p *Product) error {
	return t.Execute(os.Stdout, p)
}
