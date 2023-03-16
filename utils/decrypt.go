package utils

import "golang.org/x/crypto/bcrypt"

func PasswordCompare(hashedPw string, plainPw []byte) error {
	byteHash := []byte(hashedPw)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPw)
	if err != nil {
		return err
	}
	return nil
}

func HasAndSalt(p string) (string, error) {
	pwd := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
