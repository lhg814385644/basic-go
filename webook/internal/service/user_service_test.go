package service

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPasswordEncrypt(t *testing.T) {
	pwd := []byte("123456#123456")
	encrypt, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	// 比较(将Hash值与可能的明文进行比较，返回nil表示匹配成功)
	err = bcrypt.CompareHashAndPassword(encrypt, pwd)
	if err == nil {
		println("Compare:TRUE")
	} else {
		println("Compare:FALSE")
	}
}
