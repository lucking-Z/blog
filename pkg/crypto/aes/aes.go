package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

const (
	Zero  string = "Zero"
	PKCS5 string = "PKCS5"
	PKCS7 string = "PKCS7"
)

/*
 *plantText: plant text
  *key: size 16,24, 32 match AES-128, AES-192, AES-256
   *padding: one of Zero,PKCS5,PKCS7,default PKCS5
*/
func CBCEncrypt(plantText, key []byte, padding string) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	switch padding {
	case Zero:
		plantText = ZeroPadding(plantText, block.BlockSize())
	case PKCS7:
		plantText = PKCS7Padding(plantText, block.BlockSize())
	default:
		plantText = PKCS5Padding(plantText, block.BlockSize())
	}

	blockModel := cipher.NewCBCEncrypter(block, key)

	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

/*
 *cipherText: cipher text
  *key: size 16,24, 32 match AES-128, AES-192, AES-256
   *unpadding: one of Zero,PKCS5,PKCS7,default PKCS5
*/
func CBCDecrypt(ciphertext, key []byte, unpadding string) ([]byte, error) {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, keyBytes)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)

	fmt.Println(plantText)
	switch unpadding {
	case Zero:
		plantText = ZeroUnPadding(plantText)
	case PKCS7:
		plantText = PKCS7UnPadding(plantText)
	default:
		plantText = PKCS5UnPadding(plantText)
	}
	return plantText, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
