package pkg

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 描述: 密码加密
// 作者: hgy
// 创建日期: 2022/11/26

func GeneratePasswordHash(pwd string) string {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", pwdHash)
	return string(pwdHash)
}

func CompareHashAndPassword(hashed, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd)); err != nil {
		return false
	} else {
		return true
	}
}
