package handlers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

// UserInput : take a prompt string and return the user's reply
func UserInput(prompt string) string {
	// Variables for function use
	// Container to hold the user input
	var inputRaw string
	var err error
	// Retrieve the GOOS environment variable for compatibility
	const GOOS string = runtime.GOOS
	// Generate a new reader
	reader := bufio.NewReader(os.Stdin)
	// Prompt
	fmt.Print(prompt)
	// Read input using newline as the delimiter
	// Windows line endings: \r
	// Unix line endings: \n
	if GOOS == "windows" {
		inputRaw, err = reader.ReadString('\r')
	} else {
		inputRaw, err = reader.ReadString('\n')
	}
	if err != nil {
		log.Fatal(err)
	}
	// Remove the final character (delimeter)
	return inputRaw[:len(inputRaw)-1]
}

// GenerateFile : takes filename and contents and creates file
func GenerateFile(filename string, fileContent []byte) {
	// Variables for function use
	var err error

	// Write the file with rw-rw-rw- premissions
	err = ioutil.WriteFile(filename, fileContent, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadFile : takes a filename and returns the content
func ReadFile(filename string) []byte {
	// Variables for function use
	// Container for content
	var content []byte
	var err error

	// Read the file at the filename
	content, err = ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Return the content as octets
	return content
}
