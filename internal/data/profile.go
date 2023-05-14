package data

import (
	"context"
	"real_world/internal/biz"
	myerror "real_world/pkg/error"

	"github.com/go-kratos/kratos/v2/log"
)

type profileRepo struct {
	data *Data
	log  *log.Helper
}

func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (pr profileRepo) Followed(ctx context.Context, username, currentUser string) (bool, error) {
	var f UserFollow
	result := pr.data.db.Where(&UserFollow{
		UserName:       currentUser,
		FollowUserName: username,
	}).First(&f)

	if result.RowsAffected > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (pr profileRepo) FollowUser(ctx context.Context, followUser, currentUser string) error {
	f := UserFollow{
		UserName:       currentUser,
		FollowUserName: followUser,
	}
	result := pr.data.db.Create(&f)
	if result.Error != nil {
		return myerror.HttpBadRequest("follow", "follow fail")
	}

	return nil
}

func (pr profileRepo) UnFollowUser(ctx context.Context, followUser, currentUser string) error {
	var f UserFollow
	result := pr.data.db.Where(&UserFollow{
		UserName:       currentUser,
		FollowUserName: followUser,
	}).Delete(&f)
	if result.Error != nil {
		return myerror.HttpBadRequest("follow", "unfollow fail")
	}

	return nil
}

func (pr profileRepo) GetArticleAuthor(ctx context.Context, slug, username string) (*biz.Author, error) {
	var article Article
	result := pr.data.db.Select("author_name").Where("slug = ?", slug).First(&article)

	if result.Error != nil {
		return nil, myerror.HttpBadRequest("article", "get fail")
	}

	var user User
	pr.data.db.Where(&User{Username: article.AuthorName}).First(&user)

	followed, err := pr.Followed(ctx, user.Username, username)
	if err != nil {
		return nil, err
	}

	r := biz.Author{
		Username:  user.Username,
		Bio:       *user.Bio,
		Image:     *user.Image,
		Following: followed,
	}

	return &r, nil
}
