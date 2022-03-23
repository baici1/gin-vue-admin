// 自动生成模板EntryMember
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EntryMember 结构体
// 如果含有time.Time 请自行import time包
type EntryMember struct {
	global.GVA_MODEL
	FormId   int  `json:"formId" form:"formId" gorm:"column:form_id;comment:参赛表编号;size:10;"`
	UId      *int `json:"uId" form:"uId" gorm:"column:u_id;comment:用户编号;size:10;"`
	Identify *int `json:"identify" form:"identify" gorm:"column:identify;comment:身份;size:10;"`
	Order    *int `json:"order" form:"order" gorm:"column:order;comment:排序级别;size:10;"`
}

// TableName EntryMember 表名
func (EntryMember) TableName() string {
	return "entry_member"
}
