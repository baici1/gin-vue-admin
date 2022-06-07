package autocode

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/es"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"reflect"
	"strconv"
)

type ArticleService struct {
}

// CreateArticle 创建Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) CreateArticle(article autocode.Article) (err error) {

	err = global.GVA_DB.Create(&article).Error
	if err != nil {
		global.GVA_LOG.Error("global.GVA_DB.Create failed", zap.Error(err))
	}
	data, err := json.Marshal(&article)
	if err != nil {
		global.GVA_LOG.Error("CreateArticle article 序列化失败", zap.Error(err))
	}
	err = (&EsService{}).Create(es.ArticleIndex, strconv.Itoa(int(article.ID)), string(data))
	return err
}

// DeleteArticle 删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) DeleteArticle(article autocode.Article) (err error) {
	err = (&EsService{}).Delete(es.ArticleIndex, strconv.Itoa(int(article.ID)))
	if err != nil {
		global.GVA_LOG.Error("es Article Delete Failed", zap.Error(err))
	}
	err = global.GVA_DB.Delete(&article).Error
	return err
}

// DeleteArticleByIds 批量删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) DeleteArticleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Article{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateArticle 更新Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) UpdateArticle(article autocode.Article) (err error) {
	(&EsService{}).Update(es.ArticleIndex, strconv.Itoa(int(article.ID)), article)
	if err != nil {
		global.GVA_LOG.Error("es Article update Failed", zap.Error(err))
	}
	err = global.GVA_DB.Save(&article).Error
	return err
}

// GetArticle 根据id获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticle(id uint) (err error, article autocode.Article) {
	err = global.GVA_DB.Where("id = ?", id).First(&article).Error
	return
}

// GetArticleInfoList 分页获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) GetArticleInfoList(info autoCodeReq.ArticleSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Article{})
	var articles []autocode.Article
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("published_time > ?", info.PublishedTime)
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Author != "" {
		db = db.Where("author LIKE ?", "%"+info.Author+"%")
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Published != nil {
		db = db.Where("published = ?", info.Published)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&articles).Error
	return err, articles, total
}

func (articleService *ArticleService) EsGetArticleInfoByMatch(index, queryStr string) ([]autocode.Article, error) {
	ctx := context.Background()
	// 创建 bool 查询
	boolQuery := elastic.NewBoolQuery()
	// 短句匹配
	titleQuery := elastic.NewMatchPhrasePrefixQuery("title", queryStr).MaxExpansions(10)
	//titleQuery := elastic.NewMatchQuery("title", queryStr)
	descriptionQuery := elastic.NewMatchQuery("description", queryStr)
	// 定义highlight
	highlight := elastic.NewHighlight()
	// 指定需要高亮的字段
	highlight = highlight.Fields(elastic.NewHighlighterField("title"), elastic.NewHighlighterField("description"))
	// 指定高亮的返回逻辑 <span style='color: red;'>...msg...</span>
	highlight = highlight.PreTags("<span style='color: red;'>").PostTags("</span>")
	boolQuery.Should(titleQuery, descriptionQuery)
	boolQuery.MinimumNumberShouldMatch(1) // 至少满足 should 中的一个条件
	searchResult, err := global.GVA_ES.Search().Index(index).Query(boolQuery).
		Highlight(highlight).
		//Sort("create_time", true). // 设置排序字段，根据 create_time 字段升序排序
		From(0).  // 设置分页参数 - 起始偏移量，从第 0 行记录开始
		Size(10). // 设置分页参数 - 每页大小
		Do(ctx)   // 执行请求
	if err != nil {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		return nil, err
	}
	data := []autocode.Article{}
	for _, v := range searchResult.Each(reflect.TypeOf(autocode.Article{})) {
		data = append(data, v.(autocode.Article))
	}
	for index, highliter := range searchResult.Hits.Hits {
		if vv, ok := highliter.Highlight["title"]; ok {
			data[index].Title = vv[0]
		}
		if vv, ok := highliter.Highlight["description"]; ok {
			data[index].Description = vv[0]
		}
	}
	return data, nil
}
