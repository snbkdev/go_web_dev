package models_test

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)


type PostgresConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
	SSLMode string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func config() {
	cfg := PostgresConfig{
		Host: "localhost",
		Port: "5432",
		User: "baloo",
		Password: "junglebook",
		Database: "lenslock",
		SSLMode: "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	// Create a table
	_, err = db.Exec(`
		create table if not exists users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		create table if not exists orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created")
}