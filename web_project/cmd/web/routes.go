package main

import (
	"net/http"
	"web_project/controllers"
	"web_project/templates"
	"web_project/views"

	"github.com/go-chi/chi/v5"
)

func setupRoutes(mw appMiddleware, c appControllers) *chi.Mux {
	r := chi.NewRouter()
	r.Use(mw.CSRFMiddleware)
	r.Use(mw.UserMiddleware.SetUser)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "index.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.Get("/signup", c.Users.New)
	r.Get("/signin", c.Users.SignIn)
	r.Post("/users", c.Users.Create)
	r.Post("/signin", c.Users.ProcessSignIn)
	r.Post("/signout", c.Users.ProcessSignOut)
	r.Get("/forgot-pw", c.Users.ForgotPassword)
	r.Post("/forgot-pw", c.Users.ProcessForgotPassword)
	r.Get("/reset-pw", c.Users.ResetPassword)
	r.Post("/reset-pw", c.Users.ProcessResetPassword)

	r.Route("/users/me", func(r chi.Router) {
		r.Use(mw.UserMiddleware.RequireUser)
		r.Get("/", c.Users.CurrentUser)
	})

	r.Route("/galleries", func(r chi.Router) {
		r.Get("/{id}", c.Galleries.Show)
		r.Get("/{id}/images/{filename}", c.Galleries.Image)
		r.Group(func(r chi.Router) {
			r.Use(mw.UserMiddleware.RequireUser)
			r.Get("/", c.Galleries.Index)
			r.Get("/new", c.Galleries.New)
			r.Post("/", c.Galleries.Create)
			r.Get("/{id}/edit", c.Galleries.Edit)
			r.Post("/{id}", c.Galleries.Update)
			r.Post("/{id}/delete", c.Galleries.Delete)
			r.Post("/{id}/images", c.Galleries.UploadImage)
			r.Post("/{id}/images/{filename}/delete", c.Galleries.DeleteImage)
		})
	})

	assetsHandler := http.FileServer(http.Dir("assets"))
	r.Get("/assets/*", http.StripPrefix("/assets", assetsHandler).ServeHTTP)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	return r
}