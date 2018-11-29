package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Welcome holds information that we'll display in our HTML file
type Welcome struct {
	Name string
}

func main() {
	// We'll get the name of the user as a query param from the URL
	welcome := Welcome{"Anon"}

	// Using Must here, which will handle any errors and halt on fatals
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	// Create a handle that looks in the static dir, and uses that as a URL that our
	// HTML can refer to when looking for our CSS.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Take in the URL path, "/", and a function that takes in a response writer an an HTTP request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Grab the name form the query params
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		// Return an Internal Server Error on error
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Now we start the webserver! Localhost is assumed when there's no path provided.
	fmt.Println("listening")
	// The handler is nil, which means to use DefaultServeMux.
	fmt.Println(http.ListenAndServe(":8080", nil))
}
