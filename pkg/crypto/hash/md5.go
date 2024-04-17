package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(src []byte) string {
	h := md5.New()
	h.Write(src)
	return hex.EncodeToString(h.Sum(nil))
}
