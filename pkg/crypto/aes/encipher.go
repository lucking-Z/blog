package aes

import (
	"crypto/cipher"
	"errors"
	"fmt"
)

type IEncipher interface {
	Encrypt(a *Aes, dst, src []byte) (err error)
	Decrypt(a *Aes, dst, src []byte) (err error)
}

type CBC struct {
}

func (c CBC) Encrypt(a *Aes, dst, src []byte) (err error) {
	defer func() {
		recoverErr := recover()
		if recoverErr != nil {
			err = fmt.Errorf("cbc encrypt is err:%s", err.Error())
		}
	}()
	cipher.NewCBCEncrypter(a.block, a.iv).CryptBlocks(dst, src)
	return nil
}

func (c CBC) Decrypt(a *Aes, dst, src []byte) (err error) {
	defer func() {
		recoverErr := recover()
		if recoverErr != nil {
			err = fmt.Errorf("cbc encrypt is err:%s", err.Error())
		}
	}()
	cipher.NewCBCDecrypter(a.block, a.iv).CryptBlocks(dst, src)
	return nil
}

type ECB struct {
}

func (ECB) Encrypt(a *Aes, dst, src []byte) (err error) {
	if len(src)%a.block.BlockSize() != 0 {
		return errors.New("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		return errors.New("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		a.block.Encrypt(dst, src[:a.block.BlockSize()])
		src = src[a.block.BlockSize():]
		dst = dst[a.block.BlockSize():]
	}
	return nil
}

func (a ECB) Decrypt(b *Aes, dst, src []byte) (err error) {
	if len(src)%b.block.BlockSize() != 0 {
		return errors.New("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		return errors.New("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		b.block.Decrypt(dst, src[:b.block.BlockSize()])
		src = src[b.block.BlockSize():]
		dst = dst[b.block.BlockSize():]
	}
	return nil
}
