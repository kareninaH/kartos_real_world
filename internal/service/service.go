package service

import (
	v1 "real_world/api/real_world/v1"
	"real_world/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uuc *biz.UserUsecase
	auc *biz.ArticleUsecase
}

func NewRealWorldService(uuc *biz.UserUsecase, auc *biz.ArticleUsecase) *RealWorldService {
	return &RealWorldService{
		uuc: uuc,
		auc: auc,
	}
}
