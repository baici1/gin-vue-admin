package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ArticleService struct {
}

// CreateArticle 创建Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) CreateArticle(article autocode.Article) (err error) {
	err = global.GVA_DB.Create(&article).Error
	return err
}

// DeleteArticle 删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)DeleteArticle(article autocode.Article) (err error) {
	err = global.GVA_DB.Delete(&article).Error
	return err
}

// DeleteArticleByIds 批量删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)DeleteArticleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Article{},"id in ?",ids.Ids).Error
	return err
}

// UpdateArticle 更新Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)UpdateArticle(article autocode.Article) (err error) {
	err = global.GVA_DB.Save(&article).Error
	return err
}

// GetArticle 根据id获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)GetArticle(id uint) (err error, article autocode.Article) {
	err = global.GVA_DB.Where("id = ?", id).First(&article).Error
	return
}

// GetArticleInfoList 分页获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)GetArticleInfoList(info autoCodeReq.ArticleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Article{})
    var articles []autocode.Article
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.PublishedTime != nil {
         db = db.Where("published_time = ?",info.PublishedTime)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Author != "" {
        db = db.Where("author LIKE ?","%"+ info.Author+"%")
    }
    if info.Type != nil {
        db = db.Where("type = ?",info.Type)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&articles).Error
	return err, articles, total
}
