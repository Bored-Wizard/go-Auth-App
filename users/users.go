package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	PassWord string
}

type authUser struct {
	email        string
	passwordHash string
}

var authUserDB = map[string]authUser{}

var DefaultUserService userService

type userService struct {
}

func (userService) VerifyUser(user User) bool {
	authUser, ok := authUserDB[user.Email]
	if !ok {
		return false
	}
	passed := bcrypt.CompareHashAndPassword([]byte(authUser.passwordHash), []byte(user.PassWord))
	return passed == nil
}

func (userService) CreateUser(newUser User) error {
	_, ok := authUserDB[newUser.Email]
	if ok {
		return errors.New("User Already Exists")
	}
	passwordHash, err := getPasswordHash(newUser.PassWord)
	if err != nil {
		return err
	}
	newAuthUser := authUser{
		email:        newUser.Email,
		passwordHash: passwordHash,
	}
	authUserDB[newAuthUser.email] = newAuthUser
	return nil
}

func getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err
}
