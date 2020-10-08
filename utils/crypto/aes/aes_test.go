package aes

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	key := []byte("1111111111111111")
	var origin string = "this is plant text"

	ciphertext, err := CBCEncrypt([]byte(origin), key, PKCS5)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("ciphertext:%s\n", ciphertext)

	planttext, err := CBCDecrypt(ciphertext, key, PKCS5)
	fmt.Printf("planttext: %s\n", planttext)
	if string(origin) != string(planttext) {
		t.FailNow()
	}
}
