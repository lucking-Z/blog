package aes

import (
	"encoding/base64"
	"encoding/hex"
)

type IEncoder interface {
	Encode(src []byte) ([]byte, error)
	Decode(src []byte) ([]byte, error)
}

type Base64 struct{}

func (Base64) Encode(src []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(buf, src)
	return buf, nil
}

func (Base64) Decode(src []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(src)))
	n, err := base64.StdEncoding.Decode(buf, src)
	return buf[:n], err
}

type Hex struct{}

func (Hex) Encode(src []byte) ([]byte, error) {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst, nil
}

func (Hex) Decode(src []byte) ([]byte, error) {
	buf := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(buf, src)
	return buf[:n], err
}
