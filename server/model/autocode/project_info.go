// 自动生成模板ProjectInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProjectInfo 结构体
// 如果含有time.Time 请自行import time包
type ProjectInfo struct {
      global.GVA_MODEL
      ProjectCode  string `json:"projectCode" form:"projectCode" gorm:"column:project_code;comment:项目编号;size:64;"`
      ProjectName  string `json:"projectName" form:"projectName" gorm:"column:project_name;comment:项目名称;size:255;"`
      Introduction  string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:项目简介;size:255;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
}


// TableName ProjectInfo 表名
func (ProjectInfo) TableName() string {
  return "project_info"
}

