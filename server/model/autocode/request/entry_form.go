package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EntryFormSearch struct {
	autocode.EntryForm
	request.PageInfo
}

type EntryFormAllSearch struct {
	UId int `json:"u_id" form:"u_id"`
}

type CreateEntryForm struct {
	global.GVA_MODEL
	Name    string                `json:"name" form:"name" gorm:"column:name;comment:参赛表名;size:255;" binding:"required"`
	CmpId   *int                  `json:"cmpId" form:"cmpId" gorm:"column:cmp_id;comment:竞赛编号;size:10;" binding:"required"`
	PId     int                   `json:"pId" form:"pId" gorm:"column:p_id;comment:项目编号，0代表没有;size:10;"`
	Status  autocode.EntryMember  `json:"status"`
	Project *autocode.ProjectInfo `json:"project,omitempty" `
}

type UpdateEntryForm struct {
	global.GVA_MODEL
	Name    string                 `json:"name" form:"name" gorm:"column:name;comment:参赛表名;size:255;" binding:"required"`
	CmpId   *int                   `json:"cmpId" form:"cmpId" gorm:"column:cmp_id;comment:竞赛编号;size:10;" binding:"required"`
	PId     int                    `json:"pId" form:"pId" gorm:"column:p_id;comment:项目编号，0代表没有;size:10;"`
	Members []autocode.EntryMember `json:"status"`
	Project *autocode.ProjectInfo  `json:"project,omitempty" `
}
