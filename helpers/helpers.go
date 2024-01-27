package helpers

import (
	"fmt"
	"os"
	"strings"
)

var PasswordMap = make(map[string]string)

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
	return PasswordMap[key]
}

func UpdateAndSavePasswords(b []byte, site string, password string) {
	allPasswordsRawString := string(b[:])
	fmt.Println(allPasswordsRawString, GetUserChosenPassword(site), site, len(site))
	if strings.Contains(allPasswordsRawString, site) {
		updatedContent := strings.Replace(allPasswordsRawString, GetUserChosenPassword(site), password, 1)
		os.WriteFile("./.vault-password/password.txt", []byte(updatedContent), 0755)
	} else {
		os.WriteFile("./.vault-password/password.txt", []byte(allPasswordsRawString+site+"="+password+","), 0755)
	}
}

func PrintHelpMenu() {
	fmt.Println("To save or update a password use the following command\n save site password\nTo get a particular password enter the command\n get site\nTo get all use getall command")
}
