// 自动生成模板CompetitionInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CompetitionInfo 结构体
// 如果含有time.Time 请自行import time包
type CompetitionInfo struct {
      global.GVA_MODEL
      CName  string `json:"cName" form:"cName" gorm:"column:c_name;comment:竞赛名称;size:255;"`
      CType  *int `json:"cType" form:"cType" gorm:"column:c_type;comment:竞赛类型;size:10;"`
      Organizer  string `json:"organizer" form:"organizer" gorm:"column:organizer;comment:主办单位;size:255;"`
      Introduction  string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:简介;size:255;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:竞赛官网;size:255;"`
}


// TableName CompetitionInfo 表名
func (CompetitionInfo) TableName() string {
  return "competition_info"
}

