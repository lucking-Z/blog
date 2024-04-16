package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type Aes struct {
	p     IPadding
	c     IEncipher
	enc   IEncoder
	block cipher.Block
	iv    []byte
}

type Options func(a *Aes)

func ZeroPaddingOp() Options {
	return func(a *Aes) {
		a.p = ZeroPadding{}
	}
}

func Pkcs7PaddingOp() Options {
	return func(a *Aes) {
		a.p = Pkcs7Padding{}
	}
}

func IvOp(iv []byte) Options {
	return func(a *Aes) {
		a.iv = iv
	}
}

func Base64Op() Options {
	return func(a *Aes) {
		a.enc = Base64{}
	}
}

func HexOp() Options {
	return func(a *Aes) {
		a.enc = Hex{}
	}
}

func NewAes(key []byte, c IEncipher, op ...Options) (*Aes, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	a := &Aes{
		block: block,
		c:     c,
	}
	for _, v := range op {
		v(a)
	}
	return a, nil
}

func (a *Aes) Decrypt(src []byte) (dst []byte, err error) {
	if a.enc != nil {
		src, err = a.enc.Decode(src)
		if err != nil {
			return nil, err
		}
	}
	if a.c == nil {
		return nil, errors.New("encrypt mode is nil")
	}
	dst = make([]byte, len(src))
	err = a.c.Decrypt(a, dst, src)
	if err != nil {
		return nil, err
	}
	if a.p != nil {
		return a.p.UnPadding(dst, a.block.BlockSize())
	}
	return dst, nil
}

func (a *Aes) Encrypt(src []byte) (dst []byte, err error) {
	if a.p != nil {
		src, err = a.p.Padding(src, a.block.BlockSize())
		if err != nil {
			return nil, err
		}
	}
	if a.c == nil {
		return nil, errors.New("encrypt mode is nil")
	}
	dst = make([]byte, len(src))
	err = a.c.Encrypt(a, dst, src)
	if err != nil {
		return nil, err
	}

	if a.enc != nil {
		return a.enc.Encode(dst)
	}
	return dst, nil
}
