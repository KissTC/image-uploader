package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/kisstc/image_uploader/controllers"
	"github.com/kisstc/image_uploader/models"
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

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.html", "tailwind.html"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.html", "tailwind.html"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)

	fmt.Println("starting server on port 3000...")
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		csrf.Secure(false)) //TODO: Fix this before deploying
	http.ListenAndServe(":3000", csrfMw(r))
}
