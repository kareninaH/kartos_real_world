package data

import (
	"context"
	"real_world/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r userRepo) Create(ctx context.Context, user *biz.User) error {
	//r.data.db.Create(user)
	return nil
}

func (r userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	//TODO implement me
	panic("implement me")
}
