package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kisstc/image_uploader/controllers"
	"github.com/kisstc/image_uploader/templates"
	"github.com/kisstc/image_uploader/views"
)

func main() {
	r := chi.NewRouter()
	// middlewares
	r.Use(middleware.Logger)

	// parse the templates before starting html
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.html", "tailwind.html",
	))))

	//r.Get("/", homeHandler)
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"contact.html", "tailwind.html",
	))))

	tpl := views.Must(views.ParseFS(templates.FS, "faq.html", "tailwind.html"))
	r.Get("/faq", controllers.FAQ(tpl))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.html", "tailwind.html"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	// r.Get("/galleries/{id}", getGalleryHandler)
	// r.NotFound(func(w http.ResponseWriter, r *http.Request) {
	// 	http.Error(w, "page not found", http.StatusNotFound)
	// })
	fmt.Println("starting server on port 3000...")
	http.ListenAndServe(":3000", r)
}
