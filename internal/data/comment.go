package data

import (
	"real_world/internal/biz"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type Comment struct {
	gorm.Model
	Body      string `gorm:"not null"`
	ArticleId int    `gorm:"not null"`
}

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
