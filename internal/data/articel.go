package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"real_world/internal/biz"
	myerror "real_world/pkg/error"
)

type Article struct {
	gorm.Model
	Slug           string `gorm:"not null;type:varchar(255);index"`
	Title          string `gorm:"not null;"`
	Description    string `gorm:"not null;"`
	Body           string `gorm:"not null;"`
	FavoritesCount int    `gorm:"not null;default:0"`
	AuthorName     string `gorm:"not null;type:varchar(255)"`
}

type ArticleTag struct {
	ArticleId int
	TagId     int
}

type ArticleFavorited struct {
	gorm.Model
	Username    string `gorm:"index:idx_username_articleslug;not null:type:varchar(255)"`
	ArticleSlug string `gorm:"index:idx_username_articleslug;not null;type:varchar(255)"`
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

func (ar articleRepo) CreateArticle(ctx context.Context, a *biz.Article, username string) (*biz.Article, error) {
	article := bizToDataArticle(a)
	article.AuthorName = username
	result := ar.data.db.Create(&article)
	if result.Error != nil {
		return nil, myerror.HttpBadRequest("article", "create fail")
	}
	return dataToBizArticle(article), nil
}

func (ar articleRepo) GetArticleBySlug(ctx context.Context, slug string) (*biz.Article, error) {
	var article Article
	result := ar.data.db.Where("slug = ?", slug).First(&article)

	if result.Error != nil {
		return nil, myerror.HttpBadRequest("article", "get fail")
	}
	return dataToBizArticle(&article), nil
}

func (ar articleRepo) UpdateArticle(ctx context.Context, a *biz.Article) (*biz.Article, error) {
	article := bizToDataArticle(a)
	result := ar.data.db.Where(&Article{Slug: a.Slug}).Updates(article)
	if result.Error != nil {
		return nil, myerror.HttpBadRequest("article", "update fail")
	}
	return dataToBizArticle(article), nil
}

func (ar articleRepo) Favorited(ctx context.Context, username, slug string) bool {
	f := ArticleFavorited{
		Username:    username,
		ArticleSlug: slug,
	}
	result := ar.data.db.First(&f)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

func bizToDataArticle(a *biz.Article) *Article {
	return &Article{
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
	}
}

func dataToBizArticle(a *Article) *biz.Article {
	return &biz.Article{
		Slug:           a.Slug,
		Title:          a.Title,
		Description:    a.Description,
		Body:           a.Body,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
		FavoritesCount: a.FavoritesCount,
	}
}
