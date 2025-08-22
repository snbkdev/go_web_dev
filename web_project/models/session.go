package models

import (
	"database/sql"
	"fmt"
	"web_project/rand"
)

const (
	MinBytesPerToken = 32
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
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	// create the session token
	token, err := rand.String(bytesPerToken)
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