package helpers

import (
	"fmt"
	"os"
	"strings"

	encryption "github.com/Param-Singh/go-cli-vault/encyption"
)

var PasswordMap = make(map[string]string)

const secretkey string = "1234567890qwertyuiopasdf"

func DoesDirExist(path string) bool {
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}

func GetAllPasswords(b []byte) {
	allPasswordsRawString := string(b[:])
	passowrdKeyValues := strings.Split(allPasswordsRawString, ",")
	for i := 0; i < len(passowrdKeyValues)-1; i++ {
		PasswordMap[strings.Split(passowrdKeyValues[i], "=")[0]] = strings.Split(passowrdKeyValues[i], "=")[1]
	}
}

func GetUserChosenPassword(key string) string {
	password, err := encryption.Decrypt(PasswordMap[key], secretkey)
	if err != nil {
		return ""
	}
	return password

}

func UpdateAndSavePasswords(b []byte, site string, password string) (err error) {
	allPasswordsRawString := string(b[:])
	encryptedPassword, err := encryption.Encrypt(password, secretkey)
	if err != nil {
		return err
	}
	if strings.Contains(allPasswordsRawString, site) {
		updatedContent := strings.Replace(allPasswordsRawString, GetUserChosenPassword(site), encryptedPassword, 1)
		os.WriteFile("./.vault-password/password.txt", []byte(updatedContent), 0755)
	} else {
		os.WriteFile("./.vault-password/password.txt", []byte(allPasswordsRawString+site+"="+encryptedPassword+","), 0755)
	}
	return nil
}

func PrintHelpMenu() {
	fmt.Println("                ==> Cli Vault <==")
	fmt.Println(`========================================================
To save or update a password use the following command
set site password
========================================================
To get a particular password enter the command
get site
========================================================
To get all use getall command`)
}
