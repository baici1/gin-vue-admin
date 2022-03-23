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
	CId        int       `json:"cId" form:"cId" gorm:"column:c_id;comment:竞赛编号;size:10;"`
	Level      *int      `json:"level" form:"level" gorm:"column:level;comment:竞赛级别;size:10;"`
	Version    *int      `json:"version" form:"version" gorm:"column:version;comment:比赛的届数;size:10;"`
	StartTime  time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:报名开始日期;"`
	EndTime    time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:报名截止日期;"`
	RStartTime time.Time `json:"rStartTime" form:"rStartTime" gorm:"column:r_start_time;comment:比赛开始时间;"`
	REndTime   time.Time `json:"rEndTime" form:"rEndTime" gorm:"column:r_end_time;comment:比赛结束时间;"`
}

// TableName CompetitionSche 表名
func (CompetitionSche) TableName() string {
	return "competition_sche"
}
