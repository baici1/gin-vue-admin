package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserInfoSearch struct {
	autocode.UserInfo
	request.PageInfo
}

type UserRegister struct {
	Password      string `json:"password" form:"password" gorm:"column:password;comment:密码;size:255;" binding:"required"`
	AgainPassword string `json:"againPassword" form:"againPassword" binding:"required,eqfield=Password"`
	Phone         string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;" binding:"required"`
	Identity      int    `json:"identity" form:"identity" gorm:"column:identity;comment:身份 0是老师 1是学生;size:10;"`
}

type UserLogin struct {
	Phone     string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;" binding:"required"`
	Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;size:255;" binding:"required"`
	Captcha   string `json:"captcha" form:"captcha" binding:"required"`     // 验证码
	CaptchaId string `json:"captchaId" form:"captchaId" binding:"required"` // 验证码ID
}
