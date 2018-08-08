package main

import "time"

type User struct {
    logins_count string
    user_id string
    last_login time.Time
    name string
    picture string
    email string
 }

type Users struct{
    start int
    limit int
    length int
    users []User
    total int
 }

 type Token struct {
	Access_token          string `json:"access_token"`
	Scope                 string    `json:"scope"`
	Expires_in            int    `json:"expires_in"`
	Token_type            string `json:"token_type"`
}

 