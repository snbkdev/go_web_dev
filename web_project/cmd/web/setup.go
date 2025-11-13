package main

import (
	"database/sql"
	"net/http"
	"web_project/controllers"
	"web_project/migrations"
	"web_project/models"
	"web_project/templates"
	"web_project/views"

	"github.com/gorilla/csrf"
)

type appServices struct {
	UserService    *models.UserService
	SessionService *models.SessionService
	PwResetService *models.PasswordResetService
	GalleryService *models.GalleryService
	EmailService   *models.EmailService
}

type appMiddleware struct {
	UserMiddleware controllers.UserMiddleware
	CSRFMiddleware func(http.Handler) http.Handler
}

type appControllers struct {
	Users     controllers.Users
	Galleries controllers.Galleries
}

func setupApp(cfg config) (*sql.DB, appServices, appMiddleware, appControllers) {
	// Подключаем базу данных
	db, err := models.Open(cfg.PSQL)
	if err != nil {
		panic(err)
	}

	// Применяем миграции
	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Сервисы
	services := appServices{
		UserService:    &models.UserService{DB: db},
		SessionService: &models.SessionService{DB: db},
		PwResetService: &models.PasswordResetService{DB: db},
		GalleryService: &models.GalleryService{DB: db},
		EmailService:   models.NewEmailService(cfg.SMTP),
	}

	// Middleware
	umw := controllers.UserMiddleware{
		SessionService: services.SessionService,
	}
	csrfMw := csrf.Protect([]byte(cfg.CSRF.Key), csrf.Secure(cfg.CSRF.Secure), csrf.Path("/"))

	middleware := appMiddleware{
		UserMiddleware: umw,
		CSRFMiddleware: csrfMw,
	}

	// Controllers
	usersC := controllers.Users{
		UserService:          services.UserService,
		SessionService:       services.SessionService,
		PasswordResetService: services.PwResetService,
		EmailService:         services.EmailService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml"))
	usersC.Templates.ForgotPassword = views.Must(views.ParseFS(templates.FS, "forgot-pw.gohtml", "tailwind.gohtml"))
	usersC.Templates.CheckYourEmail = views.Must(views.ParseFS(templates.FS, "check_your_email.gohtml", "tailwind.gohtml"))
	usersC.Templates.ResetPassword = views.Must(views.ParseFS(templates.FS, "reset-pw.gohtml", "tailwind.gohtml"))

	galleriesC := controllers.Galleries{
		GalleryService: services.GalleryService,
	}
	galleriesC.Templates.New = views.Must(views.ParseFS(templates.FS, "galleries/new.gohtml", "tailwind.gohtml"))
	galleriesC.Templates.Edit = views.Must(views.ParseFS(templates.FS, "galleries/edit.gohtml", "tailwind.gohtml"))
	galleriesC.Templates.Index = views.Must(views.ParseFS(templates.FS, "galleries/index.gohtml", "tailwind.gohtml"))
	galleriesC.Templates.Show = views.Must(views.ParseFS(templates.FS, "galleries/show.gohtml", "tailwind.gohtml"))

	controllersStruct := appControllers{
		Users:     usersC,
		Galleries: galleriesC,
	}

	return db, services, middleware, controllersStruct
}