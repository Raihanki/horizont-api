package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(plainText string) (string, error) {
	hashedText, errHash := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if errHash != nil {
		return "", errHash
	}

	return string(hashedText), nil
}

func Verify(hashedText string, plainText string) (bool, error) {
	errVerify := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))
	if errVerify != nil {
		return false, errVerify
	}

	return true, nil
}
