// 自动生成模板TeacherRecruit
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// TeacherRecruit 结构体
// 如果含有time.Time 请自行import time包
type TeacherRecruit struct {
	global.GVA_MODEL
	TId       *int      `json:"tId" form:"tId" gorm:"column:t_id;comment:老师id;size:10;"`
	Picture   string    `json:"picture" form:"picture" gorm:"column:picture;comment:首图;size:255;"`
	Introduce string    `json:"introduce" form:"introduce" gorm:"column:introduce;comment:介绍;size:255;"`
	Num       *int      `json:"num" form:"num" gorm:"column:num;comment:人数;size:10;"`
	Need      string    `json:"need" form:"need" gorm:"column:need;comment:要求;size:255;"`
	Start     time.Time `json:"start" form:"start" gorm:"column:start;comment:开始;"`
	End       time.Time `json:"end" form:"end" gorm:"column:end;comment:结束;"`
}

// TableName TeacherRecruit 表名
func (TeacherRecruit) TableName() string {
	return "teacher_recruit"
}
