package security

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Return generated password end salt
func GenerateHashAndSalt(plainPassword string, cost int) (string, string, error) {
	salt := uuid.NewV4().String()

	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword+salt), cost)
	if err != nil {
		return "", "", err
	}

	return string(hash), string(salt), nil
}

func CompareHashAndSalt(plainPassword string, salt string, hash string) bool {
	saltedPassword := plainPassword + salt

	err := bcrypt.CompareHashAndPassword(([]byte(hash)), []byte(saltedPassword))
	if err != nil {
		return false
	}

	return true
}
