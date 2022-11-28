package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

type ArticleUsecase struct {
	ar  ArticleRepo
	cr  CommentRepo
	tr  TagRepo
	log *log.Helper
}

func NewArticleUsecase(ar ArticleRepo,
	cr CommentRepo,
	tr TagRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{ar: ar, cr: cr, tr: tr, log: log.NewHelper(logger)}
}
