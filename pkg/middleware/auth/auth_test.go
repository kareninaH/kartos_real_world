package auth

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// 描述: TODO
// 作者: hgy
// 创建日期: 2022/11/27

func TestJWTAuth(t *testing.T) {
	tkStr := GenerateToken("karenina", "2233")
	spew.Dump(tkStr)
}
