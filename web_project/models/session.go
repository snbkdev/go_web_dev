package models

import "database/sql"

type Session struct {
	ID int
	UserID int
	// token is only set when creating a new session
	Token string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// create the session token
	
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}