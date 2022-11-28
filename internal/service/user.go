package service

import (
	"context"
	"net/http"
	"real_world/internal/biz"

	v1 "real_world/api/real_world/v1"
	myerror "real_world/pkg/error"
)

// 描述: 用户相关api
// 作者: hgy
// 创建日期: 2022/11/26

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.UserReply, error) {
	if len(req.User.Email) == 0 {
		return nil, myerror.NewHttpError(http.StatusUnprocessableEntity, "email", "can be empty")
	}

	ul, err := s.uuc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return userLoginToUserReply(ul), nil
}
func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	ul, err := s.uuc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return userLoginToUserReply(ul), nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	ul, err := s.uuc.GetCurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	return userLoginToUserReply(ul), nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}
func (s *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) UnFollowUser(ctx context.Context, req *v1.UnFollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, req *v1.UnFavoriteArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}

func userLoginToUserReply(ul *biz.UserLogin) *v1.UserReply {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Email:    ul.Email,
			Token:    ul.Token,
			Username: ul.Username,
			Bio:      ul.Bio,
			Image:    ul.Image,
		},
	}
}
