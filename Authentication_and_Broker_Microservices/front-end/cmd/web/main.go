package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Handler to render the html test page for microservices, when '/' is accessed.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	// Setup an HHTP server that listen for incoming requests on port 80.
	fmt.Println("Starting front end service on port 80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	// Slice of strings containing all the html templates required to build the final html page.
	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t)) // Append the html test page template.

	// Loop and append the html template files.
	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	// Parse all the html template files.
	tmpl, err := template.ParseFiles(templateSlice...)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the final html page.
	// Type "http://localhost/"" in a browser to display the html page.
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
