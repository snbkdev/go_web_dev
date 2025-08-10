package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int
	Email string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	passwordHash := string(hashedBytes)
	
	user := User{
		Email: email,
		PasswordHash: passwordHash,
	}

	row := us.DB.QueryRow(`insert into users(email, password_hash) values($1, $2) returning id`, email, passwordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("Create user: %w", err)
	}

	return &user, nil
}