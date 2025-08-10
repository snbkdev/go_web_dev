package models

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

	// insert some data
	// name := "Jack"
	// email := "test2@mail.ru"
	// row := db.QueryRow(`
	// 	insert into users(name, email)
	// 	values($1, $2) returning id;`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created. id =", id)

	// id := 7
	// row := db.QueryRow(`select name, email from users where id = $1;`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("Error No Rows!")
	// }
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User information: name=%s, email=%s\n", name, email)

	// userID := 1
	// for i := 0; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("fake order #%d", i)
	// 	_, err := db.Exec(`insert into orders(user_id, amount, description) values($1, $2, $3)`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created fake orders")

	type Order struct {
		ID int
		UserID int
		Amount int
		Description string
	}

	var orders []Order
	userID := 1
	rows, err := db.Query(`select id, amount, description from orders where user_id = $1`, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserID = userID
		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	err = rows.Err()

	if rows.Err() != nil {
		panic(err)
	}

	fmt.Println("Orders:", orders)
}