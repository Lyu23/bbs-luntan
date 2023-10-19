package repositories

import (
	"gorm.io/gorm"
	"fmt"
	"bbs-go/model"
)




// ArticleGameRepository article_game 表仓库

var	ArticleGameRepository = newArticleGameRepository()

func newArticleGameRepository() *articleGameRepository {
	return &articleGameRepository{}
}

type articleGameRepository struct{}

// Get 根据 ID 取出一条数据
func (r *articleGameRepository) Get(db *gorm.DB, articleId int64) *model.ArticleGame {
	ret := &model.ArticleGame{}
	fmt.Print("nogood")
	fmt.Print(articleId)
	if err := db.First(ret, "article_id = ?", articleId).Error; err != nil {
		fmt.Print(err)
		return nil
	}
	fmt.Print("good")
	return ret
}