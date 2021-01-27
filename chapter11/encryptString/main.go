package main

import (
	"log"

	"github.com/4L3X4NND3RR/chapter11/encryptString/utils"
)

func main() {
	key := "111023043350789514532147"
	message := "I am A Message"
	log.Println("Original message: ", message)
	encryptedString := utils.EncryptString(key, message)
	log.Println("Encrypted message: ", encryptedString)
	decryptedString := utils.DecryptString(key, encryptedString)
	log.Println("Decrypted message: ", decryptedString)
}
