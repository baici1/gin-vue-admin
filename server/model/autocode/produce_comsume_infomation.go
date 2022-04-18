// 自动生成模板ProduceComsumeInfomation
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProduceComsumeInfomation 结构体
// 如果含有time.Time 请自行import time包
type ProduceComsumeInfomation struct {
	global.GVA_MODEL
	Producer int  `json:"producer" form:"producer" gorm:"column:producer;comment:响应者;size:10;"`
	Comsumer int  `json:"comsumer" form:"comsumer" gorm:"column:comsumer;comment:发布者;size:10;"`
	FormId   int  `json:"formId" form:"formId" gorm:"column:form_id;comment:参赛表;size:10;"`
	Check    bool `json:"check" form:"check" gorm:"column:check;comment:是否查看;"`
}

// TableName ProduceComsumeInfomation 表名
func (ProduceComsumeInfomation) TableName() string {
	return "produce_comsume_infomation"
}
