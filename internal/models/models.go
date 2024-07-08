package models

import "time"

// todo: разобраться с omitempty
type User struct {
	Email            string `bson:"email,omitempty"`
	UserName         string `bson:"user_name,omitempty"`
	Password         string `bson:"password,omitempty"`
	RegistrationTime time.Time
}

type Token struct {
	Token string `bson:"token,omitempty"`
	Email string `bson:"email,omitempty"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
