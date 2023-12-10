package controllers

import (
	"net/http"

	"github.com/kisstc/image_uploader/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
