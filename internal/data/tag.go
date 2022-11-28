package data

import (
	"real_world/internal/biz"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type Tag struct {
	gorm.Model
	TagName string `gorm:"not null"`
}

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
