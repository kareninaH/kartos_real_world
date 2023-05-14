package data

import (
	"context"
	"real_world/internal/biz"
	myerror "real_world/pkg/error"

	"gorm.io/gorm"

	"github.com/go-kratos/kratos/v2/log"
)

type Tag struct {
	gorm.Model
	TagName     string `gorm:"not null:type:varchar(255)"`
	ArticleSlug string `gorm:"not null:type:varchar(255)"`
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

func (tr tagRepo) SaveTags(ctx context.Context, tags []string, slug string) error {
	for _, tag := range tags {
		result := tr.data.db.Create(&Tag{
			TagName:     tag,
			ArticleSlug: slug,
		})
		if result.Error != nil {
			return myerror.HttpBadRequest("tag", "save fail")
		}
	}
	return nil
}
