package main

import (
	"fmt"
	"net/http"
	"web_project/controllers"
	"web_project/templates"
	"web_project/views"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "index.gohtml", "layout-parts.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "layout-parts.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "layout-parts.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on: 3500...")
	http.ListenAndServe(":3500", r)
}
