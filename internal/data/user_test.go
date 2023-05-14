package data

import (
	"context"
	"testing"
)

// 描述: TODO
// 作者: hgy
// 创建日期: 2022/11/28

func (r userRepo) testRedis(t *testing.T) {
	ctx := context.Background()
	r.SaveToken(ctx, "123", "123@abc.com", "demo1")
}
