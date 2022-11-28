package data

import (
	"context"
	"fmt"
	"testing"
)

// 描述: TODO
// 作者: hgy
// 创建日期: 2022/11/28

func (r userRepo) testRedis(t *testing.T) {
	ctx := context.Background()
	err := r.SaveToken(ctx, "123", "123@abc.com", "demo1")
	if err != nil {
		fmt.Println(err)
	}
}
