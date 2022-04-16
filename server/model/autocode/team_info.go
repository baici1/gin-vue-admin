// 自动生成模板TeamInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeamInfo 结构体
// 如果含有time.Time 请自行import time包
type TeamInfo struct {
	global.GVA_MODEL
	Name                 string `json:"name" form:"name" gorm:"column:name;comment:团队名称;size:255;"`
	UId                  *int   `json:"uId" form:"uId" gorm:"column:u_id;comment:负责人;size:10;"`
	CompanyId            *int   `json:"companyId" form:"companyId" gorm:"column:company_id;comment:注册公司,0表示未注册;size:10;" default:"0"`
	Introduction         string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:团队介绍;size:255;"`
	IntellectualProperty string `json:"intellectualProperty" form:"intellectualProperty" gorm:"column:intellectual_property;comment:知识产权;size:255;"`
	Remark               string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	Check                *int   `json:"check" form:"check" gorm:"column:check;comment:审核;size:10;" default:"0"`
}

// TableName TeamInfo 表名
func (TeamInfo) TableName() string {
	return "team_info"
}
