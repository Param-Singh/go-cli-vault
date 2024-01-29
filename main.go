package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Param-Singh/go-cli-vault/helpers"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided\nuse help to see the help menu")
		return
	}
	if !helpers.DoesDirExist(".vault-password") {
		// Create an hidden folder
		if err := os.Mkdir(".vault-password", 0755); err != nil {
			log.Fatal(err)
		}
	}
	if args[0] == "help" {
		helpers.PrintHelpMenu()
		return
	}
	// Read Passwords stored in the hidden in file
	helpers.MakeDir()
	b := helpers.ReadFile()
	helpers.GetAllPasswords(b)
	if args[0] == "set" {
		err := helpers.UpdateAndSavePasswords(b, args[1], args[2])
		if err != nil {
			fmt.Println(err)
		}
		return
	}
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
