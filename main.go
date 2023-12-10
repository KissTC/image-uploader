package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kisstc/image_uploader/controllers"
	"github.com/kisstc/image_uploader/views"
)

func main() {
	r := chi.NewRouter()
	// middlewares
	r.Use(middleware.Logger)

	// parse the templates before starting html
	tpl := views.Must(views.Parse(filepath.Join("templates", "home.html")))
	r.Get("/", controllers.StaticHandler(tpl))

	//r.Get("/", homeHandler)
	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.html")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.html")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	// r.Get("/galleries/{id}", getGalleryHandler)
	// r.NotFound(func(w http.ResponseWriter, r *http.Request) {
	// 	http.Error(w, "page not found", http.StatusNotFound)
	// })
	fmt.Println("starting server on port 3000...")
	http.ListenAndServe(":3000", r)
}
