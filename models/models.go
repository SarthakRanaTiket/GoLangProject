package models


//schema for users

type User struct {
	Id int `json:"id"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Email string `json:"email"`
	Age int `json:"age"`
}