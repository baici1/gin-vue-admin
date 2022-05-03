// 自动生成模板CompetitionSche
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// CompetitionSche 结构体
// 如果含有time.Time 请自行import time包
type CompetitionSche struct {
	global.GVA_MODEL
	CId       *int            `json:"cId" form:"cId" gorm:"column:c_id;comment:竞赛编号;size:10;"`
	Version   *int            `json:"version" form:"version" gorm:"column:version;comment:比赛的届数;size:10;"`
	StartTime time.Time       `json:"startTime" form:"startTime" gorm:"column:start_time;comment:报名开始日期;"`
	EndTime   time.Time       `json:"endTime" form:"endTime" gorm:"column:end_time;comment:报名截止日期;"`
	Year      *int            `json:"year" form:"year" gorm:"column:year;comment:年份;size:10;"`
	Publish   *bool           `json:"publish" form:"publish" gorm:"column:publish;comment:是否发布;"`
	BaseInfo  CompetitionInfo `json:"base_info" gorm:"foreignKey:ID;references:CId"`
}

// TableName CompetitionSche 表名
func (CompetitionSche) TableName() string {
	return "competition_sche"
}
