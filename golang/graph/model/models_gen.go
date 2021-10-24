// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
