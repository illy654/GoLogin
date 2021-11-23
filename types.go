package main

type Main struct {
	Type string
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
