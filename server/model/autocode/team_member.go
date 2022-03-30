// 自动生成模板TeamMember
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeamMember 结构体
// 如果含有time.Time 请自行import time包
type TeamMember struct {
	global.GVA_MODEL
	TeamId   int  `json:"teamId" form:"teamId" gorm:"column:team_id;comment:团队编号;size:10;"`
	UId      int  `json:"uId" form:"uId" gorm:"column:u_id;comment:用户编号;size:10;"`
	Identify *int `json:"identify" form:"identify" gorm:"column:identify;comment:身份;size:10;"`
}

// TableName TeamMember 表名
func (TeamMember) TableName() string {
	return "team_member"
}
