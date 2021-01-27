package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// EncryptString encrypts the string with given key
func EncryptString(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	cipherText := make([]byte, len(plaintext))
	cfb.XORKeyStream(cipherText, plaintext)
	return base64.StdEncoding.EncodeToString(cipherText)
}

// DecryptString decrypts the encrypted string to original
func DecryptString(key, text string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	cipherText, _ := base64.StdEncoding.DecodeString(text)
	cfb := cipher.NewCFBDecrypter(block, initVector)
	plaintext := make([]byte, len(cipherText))
	cfb.XORKeyStream(plaintext, cipherText)
	return string(plaintext)
} 
