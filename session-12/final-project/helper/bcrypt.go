package helper

import "golang.org/x/crypto/bcrypt"

func Hash(password string) string {
	salt := 8
	pass := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pass, salt)

	return string(hash)
}

func Compare(hashedPassword, password []byte) bool {
	hash, pass := []byte(hashedPassword), []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
