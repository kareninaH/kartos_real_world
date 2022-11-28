package data

import (
	"real_world/internal/biz"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	gorm.Model
	Slug           string `gorm:"not null;"`
	Title          string `gorm:"not null;"`
	Description    string `gorm:"not null;"`
	Body           string `gorm:"not null;"`
	FavoritesCount int    `gorm:"not null;default:0"`
	AuthorId       int    `gorm:"not null;"`
}

type ArticleTag struct {
	ArticleId int
	TagId     int
}

type ArticleFavorited struct {
	gorm.Model
	UserId    int `gorm:"index:idx_userid_articleid;not null"`
	ArticleId int `gorm:"index:idx_userid_articleid;not null"`
}

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
