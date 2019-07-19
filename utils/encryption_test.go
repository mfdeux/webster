package utils

import (
	"testing"
)

var testPassphrase = "pWcfnMQaU8gZsaxq7OrqxQg5762HXi2btJ1zOStwXVdbxa4K2i"

func TestEncryptionAES(t *testing.T) {
	payload := "This is a payload to pass."
	encrypted, err := EncryptAES([]byte(payload), testPassphrase)
	if err != nil {
		t.Error(err.Error())
		return
	}
	decrypted, err := DecryptAES(encrypted, testPassphrase)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if payload != string(decrypted) {
		t.Error("Payloads do not match")
		return
	}
}
