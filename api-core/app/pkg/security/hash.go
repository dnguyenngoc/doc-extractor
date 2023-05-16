package security

import "golang.org/x/crypto/bcrypt"

func VerifyHash(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
