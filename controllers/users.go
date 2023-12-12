package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// view to render
	u.Templates.New.Execute(w, nil)
}

// parsing the signup form, --> POST - respond when user submit the form
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "email: ", r.FormValue("email"))
	fmt.Fprint(w, "password: ", r.FormValue("password"))
}
