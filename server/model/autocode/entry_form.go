// 自动生成模板EntryForm
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EntryForm 结构体
// 如果含有time.Time 请自行import time包
type EntryForm struct {
	global.GVA_MODEL
	Name  string `json:"name" form:"name" gorm:"column:name;comment:参赛表名;size:255;"`
	CmpId *int   `json:"cmpId" form:"cmpId" gorm:"column:cmp_id;comment:竞赛编号;size:10;"`
	PId   *int   `json:"pId" form:"pId" gorm:"column:p_id;comment:项目编号，0代表没有;size:10;"`
}

// TableName EntryForm 表名
func (EntryForm) TableName() string {
	return "entry_form"
}
