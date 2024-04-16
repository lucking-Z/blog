package aes

import (
	"bytes"
	"errors"
	"fmt"
)

type IPadding interface {
	Padding(src []byte, blockSize int) ([]byte, error)
	UnPadding(src []byte, blockSize int) ([]byte, error)
}

type ZeroPadding struct {
}

func (ZeroPadding) Padding(src []byte, blockSize int) ([]byte, error) {
	paddingLen := blockSize - (len(src) % blockSize)
	if paddingLen == 0 {
		paddingLen = blockSize
	}
	pad := bytes.Repeat([]byte{byte(0)}, paddingLen)
	return append(src, pad...), nil
}

func (ZeroPadding) UnPadding(src []byte, blockSize int) ([]byte, error) {
	return bytes.TrimRightFunc(src, func(r rune) bool {
		return r == 0
	}), nil
}

type Pkcs7Padding struct {
}

func (Pkcs7Padding) UnPadding(src []byte, blockSize int) ([]byte, error) {
	if blockSize < 1 {
		return nil, fmt.Errorf("invalid blocklen %d", blockSize)
	}
	if len(src)%blockSize != 0 || len(src) == 0 {
		return nil, fmt.Errorf("invalid data len %d", len(src))
	}

	// the last byte is the length of padding
	paddingLen := int(src[len(src)-1])

	// check padding integrity, all bytes should be the same
	pad := src[len(src)-paddingLen:]
	for _, padbyte := range pad {
		if padbyte != byte(paddingLen) {
			return nil, errors.New("invalid padding")
		}
	}

	return src[:len(src)-paddingLen], nil
}

func (Pkcs7Padding) Padding(src []byte, blockSize int) ([]byte, error) {
	paddingLen := blockSize - (len(src) % blockSize)
	if paddingLen == 0 {
		paddingLen = blockSize
	}
	pad := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	return append(src, pad...), nil
}
