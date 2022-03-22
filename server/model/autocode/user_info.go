// 自动生成模板UserInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserInfo 结构体
// 如果含有time.Time 请自行import time包
type UserInfo struct {
	global.GVA_MODEL
	Password string `json:"password" form:"password" gorm:"column:password;comment:密码;size:255;"`
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;"`
	Slat     string `json:"slat" form:"slat" gorm:"column:slat;comment:加盐;size:255;"`
	Identity int    `json:"identity" form:"identity" gorm:"column:identity;comment:身份 1是老师 0是学生;size:10;"`
	Check    int    `json:"check" form:"check" gorm:"column:check;comment:审核老师;size:10;"`
}

// TableName UserInfo 表名
func (UserInfo) TableName() string {
	return "user_info"
}
