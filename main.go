package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// first we parse the template
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template %v", err)
		http.Error(w, "there was an error parsing the template", http.StatusInternalServerError)
		return
	}

	// execute when parsing correctly
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template %v", err)
		http.Error(w, "there was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.html")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.html")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
	<li>
	<b>Is there a free version?</b>
	Yes! We offer a free trial for 30 days on any paid plans.
	</li>
	<li>
	<b>What are your support hours?</b>
	We have support staff answering emails 24/7, though response
	times may be a bit slower on weekends.
	</li>
	<li>
	<b>How do I contact support?</b>
	Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
	</li>
	</ul>
	`)
}

func getGalleryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	id := chi.URLParam(r, "id")
	fmt.Fprint(w, "<h2>galerias de "+id+"</h2>")
}

func main() {
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/galleries/{id}", getGalleryHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("starting server on port 3000...")
	http.ListenAndServe(":3000", r)
}
