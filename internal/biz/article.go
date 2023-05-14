package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Article struct {
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	TagList        []string  `json:"tagList"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Author         Author
}

type Author struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

type ArticleRepo interface {
	CreateArticle(ctx context.Context, a *Article, username string) (*Article, error)
	Favorited(ctx context.Context, username, slug string) bool
	GetArticleBySlug(ctx context.Context, slug string) (*Article, error)
	UpdateArticle(ctx context.Context, a *Article) (*Article, error)
}

type CommentRepo interface {
}

type TagRepo interface {
	SaveTags(ctx context.Context, tags []string, slug string) error
}

type ArticleUsecase struct {
	ar  ArticleRepo
	cr  CommentRepo
	tr  TagRepo
	ur  UserRepo
	log *log.Helper
}

func NewArticleUsecase(ar ArticleRepo,
	cr CommentRepo,
	tr TagRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{ar: ar, cr: cr, tr: tr, log: log.NewHelper(logger)}
}

func (uc ArticleUsecase) CreateArticle(ctx context.Context, a Article) (*Article, error) {
	claims, err := jwtParse(ctx)
	if err != nil {
		return nil, err
	}

	username := claims["username"].(string)

	ar, err := uc.ar.CreateArticle(ctx, &a, username)
	if err != nil {
		return nil, err
	}

	err = uc.tr.SaveTags(ctx, a.TagList, ar.Slug)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func (uc ArticleUsecase) UpdateArticle(ctx context.Context, a *Article) (*Article, error) {
	ar, err := uc.ar.UpdateArticle(ctx, a)
	if err != nil {
		return nil, err
	}

	return ar, nil
}

func profileToAuthor(p *Profile) *Author {
	return &Author{
		Username:  p.Username,
		Bio:       p.Bio,
		Image:     p.Image,
		Following: p.Following,
	}
}
