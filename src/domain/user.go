package domain

import (
	"math/rand"
)

type User struct {
	Id         int
	Nombre     string
	Email      string
	Nick       string
	Contraseña string
}

func NewUser(nombre string, email string, nick string, contraseña string) *User {
	user := User{rand.Int(), nombre, email, nick, contraseña}

	return &user
}
