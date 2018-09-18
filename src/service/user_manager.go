package service

import (
	"errors"
	"github.com/tweet_manager/src/domain"
)

var userManager UserManager

type UserManager struct {
	Users []*domain.User
}

func NewUserManager() *UserManager {
	userManager = UserManager{}
	userManager.Users = make([]*domain.User, 0)

	return &userManager
}

// func Login(nickname string, password string) (domain.User, error) {
// 	return nil, nil
// }

func (manager UserManager) Register_user(nombre string, email string, nick string, contraseña string) (*domain.User, error) {
	var user *domain.User

	if !userManager.is_registered_user(nick) {
		user = domain.NewUser(nombre, email, nick, contraseña)
		manager.Users = append(manager.Users, user)
	} else {
		return nil, errors.New("User already registered")
	}

	return user, nil
}

func (manager UserManager) is_registered_user(nick string) bool {
	for _, registered_user := range manager.Users {
		if nick == registered_user.Nick {
			return true
		}
	}

	return false
}
