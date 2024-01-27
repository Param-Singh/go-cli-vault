package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Param-Singh/go-cli-vault/helpers"
)

func main() {
	fmt.Println("Welcome to Your Personal Cli Vault")
	if !helpers.DoesDirExist(".vault-password") {
		// Create an hidden folder
		if err := os.Mkdir(".vault-password", 0755); err != nil {
			log.Fatal(err)
		}
	}
	args := os.Args[1:]
	if args[0] == "help" {
		helpers.PrintHelpMenu()
		return
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
	helpers.GetAllPasswords(b)
	if args[0] == "save" {
		helpers.UpdateAndSavePasswords(b, args[1], args[2])
		return
	}

	// getting password of the one selected by user
	if args[0] == "get" {
		key := args[1]
		password := helpers.GetUserChosenPassword(key)
		fmt.Println(password)
	}
}
