package data

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// 描述: user redis
// 作者: hgy
// 创建日期: 2022/11/28

var (
	jwtTimeOut = time.Hour * 72
)

// SaveToken 保存token到redis
func (r userRepo) SaveToken(ctx context.Context, token, email, username string) {
	key := buildTokenKey(email, username)
	if _, err := r.data.rd.Set(ctx, key, token, jwtTimeOut).Result(); err != nil {
		r.log.Error("jwt token cache fail")
	}
}

func buildTokenKey(email string, username string) string {
	var sb strings.Builder
	sb.WriteString(email)
	sb.WriteString("::")
	sb.WriteString(username)
	return sb.String()
}

func (r userRepo) GetToken(ctx context.Context, email, username string) (bool, string) {
	key := buildTokenKey(email, username)
	result, err := r.data.rd.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, ""
		}
		return false, ""
	}

	return true, result
}

func (r userRepo) RemoveToken(ctx context.Context, email, username string) {
	key := buildTokenKey(email, username)
	r.data.rd.Del(ctx, key)
}
