package aes

import (
	"errors"
	"testing"
)

type example struct {
	key        []byte
	origin     []byte
	ciphertext []byte
	iv         []byte
	mode       string
	enc        string
	padding    string
}

var ex []example = []example{
	{key: []byte("1111111111111111"), origin: []byte("this is plant te"), ciphertext: []byte("YfwwsuUz5UKnASLLhD1mhPMMacT5ReZU69SziLHI95A="), mode: "ecb", padding: "pkcs7", enc: "base64"},
	{key: []byte("1111111111111111"), iv: []byte("1111111111111111"), origin: []byte("1111234"), ciphertext: []byte("GkoX1Bv7meNfgsDSs96WXQ=="), mode: "cbc", padding: "pkcs7", enc: "base64"},
	{key: []byte("1111111111111111"), iv: []byte("1111111111111111"), origin: []byte("1111234"), ciphertext: []byte("GRioNT8zBsQGNcXWJVLttw=="), mode: "cbc", padding: "zero", enc: "base64"},
	{key: []byte("1111111111111111"), iv: []byte("1111111111111111"), origin: []byte("11111111111111111111111111111111111"), ciphertext: []byte("ad0110483d0559d12e8cb97203e6c908ad0110483d0559d12e8cb97203e6c908698f287f056e4b7f271c43db3cbd8160"), mode: "ecb", padding: "zero", enc: "hex"},
}

func TestAes(t *testing.T) {
	for _, v := range ex {
		opts := make([]Options, 0)
		if v.enc == "base64" {
			opts = append(opts, Base64Op())
		} else if v.enc == "hex" {
			opts = append(opts, HexOp())
		}
		if v.padding == "pkcs7" {
			opts = append(opts, Pkcs7PaddingOp())
		} else if v.padding == "zero" {
			opts = append(opts, ZeroPaddingOp())
		}
		var blockMode IEncipher
		if v.mode == "ecb" {
			blockMode = ECB{}
		} else {
			blockMode = CBC{}
		}
		aes, err := NewAes(v.key, blockMode, opts...)
		if err != nil {
			t.Fatal(err)
		}
		res, err := aes.Decrypt(v.ciphertext)
		if err != nil {
			t.Fatal(err)
		}
		if string(res) != string(v.origin) {
			t.Fatal(errors.New("解密失败"))
		}
		res1, err := aes.Encrypt(v.origin)
		if err != nil {
			t.Fatal(err)
		}
		if string(res1) != string(v.ciphertext) {
			t.Fatal(errors.New("加密失败"))
		}
	}
}
