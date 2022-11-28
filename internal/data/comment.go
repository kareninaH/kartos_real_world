package data

import (
	"real_world/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
