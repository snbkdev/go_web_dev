package models

import (
	"database/sql"
	"fmt"
	"web_project/rand"
)

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
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	session := Session{
		UserID: userID,
		Token: token,

	}

	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}