// 自动生成模板Achievement
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Achievement 结构体
// 如果含有time.Time 请自行import time包
type Achievement struct {
      global.GVA_MODEL
      FormId  *int `json:"formId" form:"formId" gorm:"column:form_id;comment:队伍id;size:10;"`
      Match  *int `json:"match" form:"match" gorm:"column:match;comment:赛事;size:10;"`
      Rank  *int `json:"rank" form:"rank" gorm:"column:rank;comment:获奖级别;size:10;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:成果附件;size:255;"`
}


// TableName Achievement 表名
func (Achievement) TableName() string {
  return "achievement"
}

