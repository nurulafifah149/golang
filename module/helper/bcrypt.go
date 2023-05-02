package helper

import "golang.org/x/crypto/bcrypt"

func HashPass(passIn string) string {
	password := []byte(passIn)
	passHashes, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return string(passHashes)
}

func ComparePass(plainPass, hashPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(plainPass))
	return err
}
