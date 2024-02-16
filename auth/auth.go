package auth

import (
	"github.com/asawo/api/db/model"
	"golang.org/x/crypto/bcrypt"
)

// AuthenticateUser checks if the provided email and password match the user's credentials
func AuthenticateUser(user *model.User, password string) bool {
	return user != nil && checkPassword(user, password)
}

// CheckPassword verifies if the provided password matches the user's hashed password
func checkPassword(user *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// HashPassword hashes the password before storing in db
func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}
