package service

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "real_world/api/real_world/v1"
	"real_world/internal/biz"
)

// 描述: 文章,评论,tag相关api
// 作者: hgy
// 创建日期: 2022/11/26

func (s *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return &v1.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (*v1.MultipleArticlesReply, error) {
	return &v1.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) GetArticle(ctx context.Context, req *v1.GetArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) CreateArticle(ctx context.Context, req *v1.CreateArticleRequest) (*v1.SingleArticleReply, error) {
	var a biz.Article
	_ = copier.Copy(&a, req.Article)
	article, err := s.auc.CreateArticle(ctx, a)
	if err != nil {
		return nil, err
	}

	author, err := s.uuc.GetArticleAuthor(ctx, article.Slug)
	if err != nil {
		return nil, err
	}
	article.Author = *author

	return &v1.SingleArticleReply{
		Article: &v1.SingleArticleReply_Article{
			Slug:           article.Slug,
			Title:          article.Title,
			Description:    article.Description,
			Body:           article.Body,
			TagList:        article.TagList,
			CreatedAt:      article.CreatedAt.String(),
			UpdatedAt:      article.UpdatedAt.String(),
			Favorited:      article.Favorited,
			FavoritesCount: uint32(article.FavoritesCount),
			Author: &v1.SingleArticleReply_Author{
				Username:  article.Author.Username,
				Bio:       article.Author.Bio,
				Image:     article.Author.Image,
				Following: article.Author.Following,
			},
		},
	}, nil
}
func (s *RealWorldService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleRequest) (*v1.SingleArticleReply, error) {
	return &v1.SingleArticleReply{}, nil
}
func (s *RealWorldService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleRequest) (*v1.DeleteArticleReply, error) {
	return &v1.DeleteArticleReply{}, nil
}
func (s *RealWorldService) AddComment(ctx context.Context, req *v1.AddCommentRequest) (*v1.SingleCommentReply, error) {
	return &v1.SingleCommentReply{}, nil
}
func (s *RealWorldService) GetComments(ctx context.Context, req *v1.GetCommentsRequest) (*v1.MultipleCommentsReply, error) {
	return &v1.MultipleCommentsReply{}, nil
}
func (s *RealWorldService) DeleteComments(ctx context.Context, req *v1.DeleteCommentsRequest) (*v1.DeleteCommentsReply, error) {
	return &v1.DeleteCommentsReply{}, nil
}
func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (*v1.ListTagsReply, error) {
	return &v1.ListTagsReply{}, nil
}
