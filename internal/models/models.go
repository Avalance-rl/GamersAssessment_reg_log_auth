package models

import "time"

type User struct {
	Email            string `bson:"email,omitempty"`
	UserName         string `bson:"user_name,omitempty"`
	Password         string `bson:"password,omitempty"`
	RegistrationTime time.Time
}

type RegisterRequest struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
