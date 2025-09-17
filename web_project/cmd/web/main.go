package main

import (
	"fmt"
	"net/http"
	"web_project/controllers"
	"web_project/migrations"
	"web_project/models"
	"web_project/templates"
	"web_project/views"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func main() {
	// Connect the database
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	// Setup services
	userService := models.UserService{
		DB: db,
	}

	sessionService := models.SessionService{
		DB: db,
	}

	// Setup middleware
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}

	csrfKey := []byte("32-byte-long-auth-key-1234567890abcd")
	csrfMw := csrf.Protect([]byte(csrfKey), csrf.Secure(false))

	// setup controllers
	usersC := controllers.Users{UserService: &userService, SessionService: &sessionService}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))

	// Setup router and routes
	r := chi.NewRouter()
	r.Use(csrfMw)
	r.Use(umw.SetUser)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "index.gohtml", "tailwind.gohtml"))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	r.Get("/signup", usersC.New)
	r.Get("/signin", usersC.SignIn)
	r.Post("/users", usersC.Create)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})
	//r.Get("/users/me", usersC.CurrentUser)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	// Start the server
	fmt.Println("Starting the server on: 3500...")
	http.ListenAndServe(":3500", r)
}

// func TimerMiddleware(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		h(w,r)
// 		fmt.Println("Request time:", time.Since(start))
// 	}
// }
