// 自动生成模板StudentInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// StudentInfo 结构体
// 如果含有time.Time 请自行import time包
type StudentInfo struct {
	global.GVA_MODEL
	StudentId      string `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学号;size:20;"`
	Phone          string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;size:255;"`
	Password       string `json:"-" form:"password" gorm:"column:password;comment:密码;size:255;"`
	AuthorityId    string `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:用户身份;size:191;"`
	Nickname       string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称;size:20;"`
	Email          string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:255;"`
	Avatar         string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;size:255;"`
	RealName       string `json:"realName" form:"realName" gorm:"column:real_name;comment:姓名;size:255;"`
	Gender         *int   `json:"gender" form:"gender" gorm:"column:gender;comment:性别;size:10;"`
	Degree         string `json:"degree" form:"degree" gorm:"column:degree;comment:学历;size:6;"`
	Grade          string `json:"grade" form:"grade" gorm:"column:grade;comment:年级;size:6;"`
	Department     string `json:"department" form:"department" gorm:"column:department;comment:学院;size:20;"`
	Major          string `json:"major" form:"major" gorm:"column:major;comment:专业;size:20;"`
	ClassNum       string `json:"classNum" form:"classNum" gorm:"column:class_num;comment:班级;size:8;"`
	Specialty      string `json:"specialty" form:"specialty" gorm:"column:specialty;comment:特长;size:255;"`
	QQ             string `json:"QQ" form:"QQ" gorm:"column:QQ;comment:QQ号;size:12;"`
	Wechat         string `json:"wechat" form:"wechat" gorm:"column:wechat;comment:微信号;size:255;"`
	BankName       string `json:"bankName" form:"bankName" gorm:"column:bank_name;comment:银行名称;size:30;"`
	BankCardNumber string `json:"bankCardNumber" form:"bankCardNumber" gorm:"column:bank_card_number;comment:银行卡号;size:30;"`
	Introduction   string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:介绍;size:255;"`
	//用户权限信息
	Authority system.SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
}

// TableName StudentInfo 表名
func (StudentInfo) TableName() string {
	return "student_info"
}
