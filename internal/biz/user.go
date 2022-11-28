package biz

import (
	"context"
	"real_world/internal/conf"
	"real_world/pkg"
	"real_world/pkg/middleware/auth"
	"strings"

	"github.com/jinzhu/copier"

	"github.com/go-kratos/kratos/v2/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Email          string
	Username       string
	Bio            string
	Image          string
	PasswordHashed string
}

type UserLogin struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

type UserRepo interface {
	Create(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur  UserRepo
	jwt *conf.JWT
	pr  ProfileRepo
	log *log.Helper
}

func NewUserUsecase(ur UserRepo, jwt *conf.JWT,
	pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, jwt: jwt, pr: pr, log: log.NewHelper(logger)}
}

func (uc UserUsecase) generateToken(username string) string {
	return auth.GenerateToken(uc.jwt.Secret, username)
}

func (uc UserUsecase) Register(ctx context.Context, username, email, pwd string) (*UserLogin, error) {
	u := &User{
		Email:          email,
		Username:       username,
		PasswordHashed: pkg.GeneratePasswordHash(pwd),
	}
	if err := uc.ur.Create(ctx, u); err != nil {
		return nil, err
	} else {
		// jwt token
		var sb strings.Builder
		sb.WriteString("Token ")
		sb.WriteString(uc.generateToken(u.Username))
		return &UserLogin{
			Email:    email,
			Username: username,
			Token:    sb.String(),
		}, nil
	}
}

func (uc UserUsecase) Login(ctx context.Context, email, pwd string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !pkg.CompareHashAndPassword(u.PasswordHashed, pwd) {
		return nil, errors.New(401, "password error", "login error")
	}
	var ul UserLogin
	err = copier.Copy(ul, u)
	if err != nil {
		return nil, err
	}
	var sb strings.Builder
	sb.WriteString("Token ")
	sb.WriteString(uc.generateToken(u.Username))
	ul.Token = sb.String()
	return &ul, nil
}
