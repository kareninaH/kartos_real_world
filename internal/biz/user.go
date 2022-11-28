package biz

import (
	"context"
	"net/http"
	"real_world/internal/conf"
	"real_world/pkg"
	myerror "real_world/pkg/error"
	"real_world/pkg/middleware/auth"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"github.com/jinzhu/copier"

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
	SaveToken(ctx context.Context, token, email, username string)
	GetToken(ctx context.Context, email, username string) (bool, string)
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

func (uc UserUsecase) generateToken(username, email string) string {
	return auth.GenerateToken(uc.jwt.Secret, username, email)
}

func (uc UserUsecase) Register(ctx context.Context, username, email, pwd string) (*UserLogin, error) {
	u := &User{
		Email:          email,
		Username:       username,
		PasswordHashed: pkg.GeneratePasswordHash(pwd),
	}

	if err := uc.ur.Create(ctx, u); err != nil {
		return nil, err
	}

	token := uc.isTokenActivate(ctx, u)

	return &UserLogin{
		Email:    email,
		Username: username,
		Token:    token,
	}, nil
}

func (uc UserUsecase) Login(ctx context.Context, email, pwd string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !pkg.CompareHashAndPassword(u.PasswordHashed, pwd) {
		return nil, myerror.HttpUnauthorized("password", "密码错误")
	}

	return uc.userToUserLogin(ctx, u)
}

// isTokenActivate token是否存在, 不存在就创建并保存到redis, 返回token
func (uc UserUsecase) isTokenActivate(ctx context.Context, u *User) string {
	flag, token := uc.ur.GetToken(ctx, u.Email, u.Username)
	if flag {
		return token
	} else {
		var sb strings.Builder
		sb.WriteString("Token ")
		sb.WriteString(uc.generateToken(u.Username, u.Email))
		token = sb.String()
		uc.ur.SaveToken(ctx, token, u.Email, u.Username)
		return token
	}
}

func (uc UserUsecase) GetCurrentUser(ctx context.Context) (*UserLogin, error) {
	claims, ok := auth.FromContext(ctx)
	if !ok {
		return nil, myerror.HttpUnauthorized("jwt", "get fail")
	}
	email := claims.(jwt.MapClaims)["email"].(string)
	user, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return uc.userToUserLogin(ctx, user)
}

func (uc UserUsecase) userToUserLogin(ctx context.Context, u *User) (*UserLogin, error) {
	var ul UserLogin
	err := copier.Copy(&ul, u)
	if err != nil {
		return nil, myerror.NewHttpError(http.StatusInternalServerError, "copier", "copy fail")
	}

	ul.Token = uc.isTokenActivate(ctx, u)
	return &ul, nil
}
