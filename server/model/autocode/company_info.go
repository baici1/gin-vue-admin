// 自动生成模板CompanyInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CompanyInfo 结构体
// 如果含有time.Time 请自行import time包
type CompanyInfo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:公司名称;size:255;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:公司地址;size:255;"`
      Introduction  string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:公司介绍;size:255;"`
      Check  *int `json:"check" form:"check" gorm:"column:check;comment:审核;size:10;"`
}


// TableName CompanyInfo 表名
func (CompanyInfo) TableName() string {
  return "company_info"
}

