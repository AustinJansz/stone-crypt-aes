package main

import (
	"fmt"
	"log"
	"os"

	"github.com/theuhrmacher/stone-crypt-aes-public/stone-crypt-aes/handlers"
)

// Globally available variables
var (
	keyRaw       string
	keyLengthRaw string
	filename     string
	key          []byte
	secret       []byte
	cipher       []byte
)

// keyUI : takes a raw key and key length from the user and formats key
func keyUI() {
	// User input as string
	keyRaw = handlers.UserInput("Enter key: ")
	keyLengthRaw = handlers.UserInput("Enter a key length [32]: ")
	// Handler function to return a key as octets
	key = handlers.KeyFormatter(keyRaw, keyLengthRaw)
}

// encryptUI : main user interface for the encryption function
func encryptUI() {
	// Check with the user to see if they would like to use a file as input
	file := handlers.UserInput("Would you like to read from a file [y|n]? ")
	if file == "y" {
		// Get the name of the file
		inputFile := handlers.UserInput("Filename to read from: ")
		// Pass the filename to the file reader
		secret = handlers.ReadFile(inputFile)
	} else {
		// Ask the user to enter a message to encrypt
		secret = []byte(handlers.UserInput("Enter the text to encrypt: "))
		// Get a filename to save the encrypted message
		filename = handlers.UserInput("Filename to write to: Encrypted_")
	}
	// Call on the keyUI to create a valid key
	keyUI()
	// Generate the cipher using the encryption handler
	cipher = handlers.OctetEncryptor(secret, key)
	// Generate the file with the encrypted contents
	handlers.GenerateFile(("Encrypted_" + filename), cipher)
}

// decryptUI : main user interface for the decryption function
func decryptUI() {
	// Ask the user for the file where the decrypted contents are stored
	filename = handlers.UserInput("Filename to read from: ")
	// Capture the cipher stored in the file
	cipher = handlers.ReadFile(filename)
	// Call on the keyUI to create a valid key
	keyUI()
	// Decrypt the content of the file using the handler function
	secret = handlers.OctetDecryptor(cipher, key)
	// Generate the file with the decrypted secret
	handlers.GenerateFile(("Decrypted_" + filename), secret)
}

// main : Main entrypoint for the program
func main() {
	// Splash screen for the program
	fmt.Printf(`
=======================================================
   ______________  _  ______  __________  _____  ______
  / __/_  __/ __ \/ |/ / __/ / ___/ _ \ \/ / _ \/_  __/
 _\ \  / / / /_/ /    / _/  / /__/ , _/\  / ___/ / /   
/___/ /_/  \____/_/|_/___/  \___/_/|_| /_/_/    /_/ 

					AES Edition
=======================================================

> Constructor: TheUhrMacher (Austin Jansz)
> The super simple crypto-tool
> Algorimth: AES (GCM Mode)
> Key lengths:
>	16: 128-bit AES
>	24: 192-bit AES
>	32: 256-bit AES

Please select a mode:
	[1] Encrypt File or Message to File
	[2] Decrypt from File
	[3] Exit

`)
	// Mode switcher to handle the direction of the program
	switch mode := handlers.UserInput("Mode: "); mode {
	case "1":
		encryptUI()
	case "2":
		decryptUI()
	case "3":
		os.Exit(1)
	default:
		log.Fatal("Select a valid mode.")
	}
}
