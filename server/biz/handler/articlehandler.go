package handler

import (
	"context"

	"github.com/moumou/server/biz/model"

	"github.com/moumou/server/biz/conv"
	"github.com/moumou/server/biz/service"
	pb "github.com/moumou/server/gen/proto"
)

type ArticleHandlerService struct {
	svc       *service.Service
	converter conv.IConverter
}

func NewArticleHandlerService(svc *service.Service, converter conv.IConverter) pb.ArticleHandlerHTTPServer {
	return &ArticleHandlerService{svc: svc, converter: converter}
}

func (s *ArticleHandlerService) GetArticleList(ctx context.Context, req *pb.GetArticleListRequest) (*pb.GetArticleListResponse, error) {
	articleList, total, err := s.svc.Dao.ArticleDao(ctx).Find()
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleListResponse{
		Data: &pb.GetArticleListResponseData{
			List:  s.converter.ConvertArticleListToVO(articleList),
			Total: total,
		},
	}, nil
}
func (s *ArticleHandlerService) GetArticleInfo(ctx context.Context, req *pb.GetArticleInfoRequest) (*pb.GetArticleInfoResponse, error) {
	article, err := s.svc.Dao.ArticleDao(ctx).Preload("ArticleContent").GetByID(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetArticleInfoResponse{
		Data: s.converter.ConvertArticleToVO(article),
	}, nil
}
func (s *ArticleHandlerService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	article := &model.Article{}
	s.converter.ConvertCreateArticleRequestDataToBO(req.GetArticle(), article)
	err := s.svc.Dao.ArticleDao(ctx).Create(article)
	if err != nil {
		return nil, err
	}
	return &pb.CreateArticleResponse{
		Data: &pb.CreateArticleResponseData{
			Id: article.Id,
		},
	}, nil
}
func (s *ArticleHandlerService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	return &pb.UpdateArticleResponse{}, nil
}
func (s *ArticleHandlerService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	return &pb.DeleteArticleResponse{}, nil
}
