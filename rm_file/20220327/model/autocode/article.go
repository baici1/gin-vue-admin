// 自动生成模板Article
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Article 结构体
// 如果含有time.Time 请自行import time包
type Article struct {
	global.GVA_MODEL
	Commentabled  *bool      `json:"commentabled" form:"commentabled" gorm:"column:commentabled;comment:是否可评论;"`
	Published     *bool      `json:"published" form:"published" gorm:"column:published;comment:是否发表;"`
	PublishedTime *time.Time `json:"publishedTime" form:"publishedTime" gorm:"column:published_time;comment:发布时间;"`
	Title         string     `json:"title" form:"title" gorm:"column:title;comment:标题;size:255;"`
	Description   string     `json:"description" form:"description" gorm:"column:description;comment:引用;size:255;"`
	Content       string     `json:"content" form:"content" gorm:"column:content;comment:内容;size:4294967295;"`
	Views         *int       `json:"views" form:"views" gorm:"column:views;comment:浏览量;size:10;"`
	Author        string     `json:"author" form:"author" gorm:"column:author;comment:作者;size:8;"`
	Type          *int       `json:"type" form:"type" gorm:"column:type;comment:文章类型;size:10;"`
	OrderNum      *bool      `json:"orderNum" form:"orderNum" gorm:"column:order_num;comment:置顶;"`
}

// TableName Article 表名
func (Article) TableName() string {
	return "article"
}
