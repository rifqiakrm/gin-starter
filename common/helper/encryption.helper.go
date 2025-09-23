package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1" // #nosec
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 10
)

// Decrypt encrypts a string.
func Decrypt(encryptedString string, keyString string) string {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}

// SHAEncrypt encrypts a string
func SHAEncrypt(plainText string) string {
	sha := sha1.New() // #nosec
	sha.Write([]byte(plainText))
	encrypted := sha.Sum(nil)
	encryptedString := fmt.Sprintf("%x", encrypted)

	return encryptedString
}

// BcryptEncrypt encrypts a string.
func BcryptEncrypt(plainText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), cost)
	return string(hashed), err
}

// BcryptVerifyHash compares hashed and plain string.
func BcryptVerifyHash(encrypted, plain string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(plain)); err != nil {
		return false
	}
	return true
}
