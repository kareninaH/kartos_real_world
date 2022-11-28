package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/davecgh/go-spew/spew"
)

// 描述: pwd_hash_test
// 作者: hgy
// 创建日期: 2022/11/26

func TestGeneratePasswordHash(t *testing.T) {
	hash := GeneratePasswordHash("123456")
	spew.Dump(hash)
}

func TestCompareHashAndPassword(t *testing.T) {
	a := assert.New(t)
	a.True(true,
		CompareHashAndPassword("$2a$10$pgEASJk2jiVIfKEjYMaOAuj/tXdUkir5vAOxGkVCEbMule8qnFoRe", "123456"))
}
