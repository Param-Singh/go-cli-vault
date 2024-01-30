package encryptiontest

import (
	"testing"

	encryption "github.com/Param-Singh/go-cli-vault/encyption"
)

func TestEncryption(t *testing.T) {
	text := "sjfiasojfsoidfjsoaidfjsdlfk"
	secretKey := "1234567890qwertyuiopasdf"
	result, err := encryption.Encrypt(text, secretKey)
	if err != nil {
		t.Error(err)
	}
	if result == "" {
		t.Errorf("The result string is empty")
	}
}

func TestDecryption(t *testing.T) {
	text := "�A�,�϶w��MGQ�_;�!����"
	secretKey := "1234567890qwertyuiopasdf"
	result, err := encryption.Decrypt(text, secretKey)
	if err != nil {
		t.Error(err)
	}
	if result == "" {
		t.Errorf("The result string is empty")
	}
}
