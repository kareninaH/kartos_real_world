package data

import (
	"real_world/internal/conf"
	"testing"
)

// 描述: 数据库连接测试
// 作者: hgy
// 创建日期: 2022/11/26

func TestNewDB(t *testing.T) {
	NewDB(&conf.Data{}, nil)
}
