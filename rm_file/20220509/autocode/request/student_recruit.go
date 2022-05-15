package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type StudentRecruitSearch struct {
	autocode.StudentRecruit
	request.PageInfo
}

type StudentRecruitCreate struct {
	autocode.StudentInfo    `json:"student_info"`
	autocode.StudentRecruit `json:"student_recruit"`
}
