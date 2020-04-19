package handlers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// UserInput : take a prompt string and return the user's reply
func UserInput(prompt string) string {
	// Generate a new reader
	reader := bufio.NewReader(os.Stdin)
	// Prompt
	fmt.Print(prompt)
	// Read input using newline as the delimiter
	inputRaw, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Remove the final character (delimeter)
	return inputRaw[:len(inputRaw)-1]
}

// GenerateFile : takes filename and contents and creates file
func GenerateFile(filename string, fileContent []byte) {
	// Write the file with rw-rw-rw- premissions
	err := ioutil.WriteFile(filename, fileContent, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadFile : takes a filename and returns the content
func ReadFile(filename string) []byte {
	// Read the file at the filename
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Return the content as octets
	return content
}
