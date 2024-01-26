package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to Your Personal Cli Vault")
	// Create an hidden folder
	if err := os.Mkdir(".vault-password", 0755); err != nil {
		log.Fatal(err)
	}
	// create the file which will contain passwords
	f, err := os.OpenFile("./.vault-password/password.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil || f.Close() != nil {
		log.Fatal(err)
	}
	// Read the file for all the passwords
	b, err := os.ReadFile("./.vault-password/password.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
