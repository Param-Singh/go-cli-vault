package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Param-Singh/go-cli-vault/helpers"
)

func main() {
	if !helpers.DoesDirExist(".vault-password") {
		// Create an hidden folder
		if err := os.Mkdir(".vault-password", 0755); err != nil {
			log.Fatal(err)
		}
	}
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided\nuse help to see the help menu")
		return
	}
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
	if args[0] == "set" {
		err := helpers.UpdateAndSavePasswords(b, args[1], args[2])
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	// getting password of the one selected by user
	if args[0] == "get" {
		key := args[1]
		password := helpers.GetUserChosenPassword(key)
		fmt.Println(password)
	}
	if args[0] == "getall" {
		for site, password := range helpers.PasswordMap {
			fmt.Println(site, "=> ", password)
		}
	}
}

