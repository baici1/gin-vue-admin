// 自动生成模板EntryTeacher
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EntryTeacher 结构体
// 如果含有time.Time 请自行import time包
type EntryTeacher struct {
      global.GVA_MODEL
      FormId  *int `json:"formId" form:"formId" gorm:"column:form_id;comment:参赛表id;size:10;"`
      TId  *int `json:"tId" form:"tId" gorm:"column:t_id;comment:老师id;size:10;"`
      Order  *int `json:"order" form:"order" gorm:"column:order;comment:排序;size:10;"`
}


// TableName EntryTeacher 表名
func (EntryTeacher) TableName() string {
  return "entry_teacher"
}

