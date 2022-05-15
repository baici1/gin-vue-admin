// 自动生成模板StudentRecruit
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// StudentRecruit 结构体
// 如果含有time.Time 请自行import time包
type StudentRecruit struct {
	global.GVA_MODEL
	ComId     *int      `json:"comId" form:"comId" gorm:"column:com_id;comment:比赛id;size:10;"`
	EntryId   *int      `json:"entryId" form:"entryId" gorm:"column:entry_id;comment:参赛表对应id;size:10;"`
	UId       *int      `json:"uId" form:"uId" gorm:"column:u_id;comment:发布者id;size:10;"`
	Picture   string    `json:"picture" form:"picture" gorm:"column:picture;comment:首图;size:255;"`
	Introduce string    `json:"introduce" form:"introduce" gorm:"column:introduce;comment:介绍;size:255;"`
	Num       *int      `json:"num" form:"num" gorm:"column:num;comment:人数;size:10;"`
	Need      string    `json:"need" form:"need" gorm:"column:need;comment:要求;size:255;"`
	Start     time.Time `json:"start" form:"start" gorm:"column:start;comment:开始时间;"`
	End       time.Time `json:"end" form:"end" gorm:"column:end;comment:;"`
}

// TableName StudentRecruit 表名
func (StudentRecruit) TableName() string {
	return "student_recruit"
}
