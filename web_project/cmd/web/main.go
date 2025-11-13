package main

import (
	"fmt"
	"net/http"
)

func main() {
	cfg, err := loadEnvConfig()
	if err != nil {
		panic(err)
	}

	db, _, middleware, controllers := setupApp(cfg)
	defer db.Close()

	r := setupRoutes(middleware, controllers)

	fmt.Printf("Starting the server on %s ...\n", cfg.Server.Address)
	err = http.ListenAndServe(cfg.Server.Address, r)
	if err != nil {
		panic(err)
	}
}