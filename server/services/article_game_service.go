package services

import (

	"github.com/mlogclub/simple/sqls"

	"bbs-go/model"
	"bbs-go/repositories"
)

// ArticleGameService 桥接 articleGameRepository，提供具体的业务方法

var ArticleGameService = newArticleGameService()

func newArticleGameService() *articleGameService{
	return &articleGameService{}
}

type articleGameService struct {
}

// Get 根据 id 取出一条数据
func (s *articleGameService) Get(articleId int64) *model.ArticleGame {
	return repositories.ArticleGameRepository.Get(sqls.DB(), articleId)
}