package md5

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	fmt.Println(Md5([]byte("test")))
}

func TestMd5File(t *testing.T) {
	_, err := Md5File("./md5.go")
	if err != nil {
		t.Error(err)
	}
}
