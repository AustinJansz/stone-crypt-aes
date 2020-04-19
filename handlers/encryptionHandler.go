package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"strconv"
)

// KeyFormatter : takes key as string and formats it to 32 octet array
func KeyFormatter(keyRaw, keyLengthRaw string) []byte {
	// Initialize variables for key and key length
	var err error
	keyLength := 0
	keyString := ""

	// Check to see if the user selected the key length default (32)
	if keyLengthRaw == "" {
		keyLength = 32
	} else if keyLengthRaw != "16" && keyLengthRaw != "24" && keyLengthRaw != "32" {
		log.Fatal("Length of key is not one of the valid options (16, 24, 32)")
	} else {
		// Check to see if the user input can be passed as an integer
		keyLength, err = strconv.Atoi(keyLengthRaw)
		if err != nil {
			log.Fatal("Number required for key length")
		}
	}
	// Check to see if key is too long and cut it to length
	if keyLength < len(keyRaw) {
		keyString = keyRaw[0 : keyLength-1]
	} else {
		// Increase the size of the key to the desired key length
		counter := 0
		// Key cycle until the length is the same as the desired key length
		for len(keyString) < keyLength {
			keyString += string(keyRaw[counter%len(keyRaw)])
			counter++
		}
	}
	// Format the key as bytes and return the key
	return []byte(keyString)
}

// OctetEncryptor : takes key and message as octets and returns encrypted octets
func OctetEncryptor(message, key []byte) []byte {
	// Generate new AES key
	aesKey, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	// Generate new GCM wrapped aes key
	gcm, err := cipher.NewGCM(aesKey)
	if err != nil {
		log.Fatal(err)
	}
	// Generate nonce based on GCM
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}
	// Return the encrypted message as octets
	return gcm.Seal(nonce, nonce, message, nil)
}

// OctetDecryptor : takes key and ciphertext as octets and returns decrypted octets
func OctetDecryptor(message, key []byte) []byte {
	// Generate new AES key
	aesKey, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	// Generate new GCM wrapped aes key
	gcm, err := cipher.NewGCM(aesKey)
	if err != nil {
		log.Fatal(err)
	}
	// Generate nonce based on GCM
	nonceSize := gcm.NonceSize()
	if len(message) < nonceSize {
		log.Fatal(err)
	}
	nonce, encryptedMessage := message[:nonceSize], message[nonceSize:]
	// Decrypted the secret message
	secret, err := gcm.Open(nil, nonce, encryptedMessage, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Return message as octets
	return secret
}
