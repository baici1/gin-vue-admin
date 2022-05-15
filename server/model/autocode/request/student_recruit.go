package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type StudentRecruitSearch struct {
	autocode.StudentRecruit
	request.PageInfo
}

type StudentRequestInfo struct {
	StudentId    string `json:"studentId" form:"studentId" gorm:"column:student_id;comment:学号;size:20;"`
	RealName     string `json:"realName" form:"realName" gorm:"column:real_name;comment:姓名;size:255;"`
	Gender       *int   `json:"gender" form:"gender" gorm:"column:gender;comment:性别;size:10;"`
	Degree       string `json:"degree" form:"degree" gorm:"column:degree;comment:学历;size:6;"`
	Grade        string `json:"grade" form:"grade" gorm:"column:grade;comment:年级;size:6;"`
	Department   string `json:"department" form:"department" gorm:"column:department;comment:学院;size:20;"`
	Major        string `json:"major" form:"major" gorm:"column:major;comment:专业;size:20;"`
	ClassNum     string `json:"classNum" form:"classNum" gorm:"column:class_num;comment:班级;size:8;"`
	QQ           string `json:"QQ" form:"QQ" gorm:"column:QQ;comment:QQ号;size:12;"`
	Wechat       string `json:"wechat" form:"wechat" gorm:"column:wechat;comment:微信号;size:255;"`
	Introduction string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:介绍;size:255;"`
}

type StudentRecruitToRabbitmq struct {
	Producer StudentRequestInfo      `json:"producer"`
	Comsumer autocode.StudentRecruit `json:"comsumer"`
}
