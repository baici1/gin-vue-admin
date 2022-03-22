// 自动生成模板TeacherInfo
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TeacherInfo 结构体
// 如果含有time.Time 请自行import time包
type TeacherInfo struct {
	global.GVA_MODEL
	PersonnelId    string `json:"personnelId" form:"personnelId" gorm:"column:personnel_id;comment:人事编号;size:10;"`
	OfficeId       string `json:"officeId" form:"officeId" gorm:"column:office_id;comment:教务编号;size:10;"`
	FinancialId    string `json:"financialId" form:"financialId" gorm:"column:financial_id;comment:财务编号;size:10;"`
	UId            int    `json:"uId" form:"uId" gorm:"column:u_id;comment:用户编号;size:10;"`
	Nickname       string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称;size:255;"`
	Email          string `json:"email" form:"email" gorm:"column:email;comment:邮箱;size:255;"`
	Avatar         string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;size:255;"`
	RealName       string `json:"realName" form:"realName" gorm:"column:real_name;comment:姓名;size:255;"`
	Gender         *int   `json:"gender" form:"gender" gorm:"column:gender;comment:性别;size:10;"`
	Department     string `json:"department" form:"department" gorm:"column:department;comment:学院;size:20;"`
	Major          string `json:"major" form:"major" gorm:"column:major;comment:专业方向;size:255;"`
	Position       string `json:"position" form:"position" gorm:"column:position;comment:职称;size:10;"`
	Specialty      string `json:"specialty" form:"specialty" gorm:"column:specialty;comment:特长;size:255;"`
	Degree         string `json:"degree" form:"degree" gorm:"column:degree;comment:学历;size:10;"`
	BankName       string `json:"bankName" form:"bankName" gorm:"column:bank_name;comment:银行名称;size:30;"`
	BankCardNumber string `json:"bankCardNumber" form:"bankCardNumber" gorm:"column:bank_card_number;comment:银行卡号;size:30;"`
	Introduction   string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:介绍;size:255;"`
}

// TableName TeacherInfo 表名
func (TeacherInfo) TableName() string {
	return "teacher_info"
}
