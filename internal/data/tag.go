package data

import (
	"real_world/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
