package data

import (
	"real_world/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
