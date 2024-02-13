package models

import "golang.org/x/oauth2"

type User struct {
	Username string       `json:"username"`
	Email    string       `json:"email"`
	ID       string       `json:"id"`
	Token    oauth2.Token `json:"token"`
}
