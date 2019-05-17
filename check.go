package check

import (
	"fmt"
	"log"
	"os"
)

// Error prints the error to stderr if it's not nil
func Error(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

// Log prints the error via log if not nil
func Log(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Panic will panic the error if not nil
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

// Print does a normal print of the error if not nil
func Print(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// String returns a string for the error. An empty string if nil.
func String(err error) string {
	var result string
	if err != nil {
		result = fmt.Sprintf("%s", err)
	}
	return result
}
