package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

type EnteryFormDetail struct {
	global.GVA_MODEL
	Name        string                   `json:"name" form:"name" gorm:"column:name;comment:参赛表名;size:255;"`
	CmpId       *int                     `json:"cmpId" form:"cmpId" gorm:"column:cmp_id;comment:竞赛编号;size:10;"`
	PId         *int                     `json:"pId" form:"pId" gorm:"column:p_id;comment:项目编号，0代表没有;size:10;"`
	Rank        *int                     `json:"rank" form:"rank" gorm:"column:rank;comment:获奖级别;size:10;"`
	AchName     string                   `json:"achName" form:"achName" gorm:"column:ach_name;comment:获奖名称;size:255;"`
	Members     []autocode.EntryMember   `json:"members" gorm:"foreignkey:FormId;references:ID"`
	Project     autocode.ProjectInfo     `json:"project,omitempty" gorm:"foreignkey:ID;references:PId"`
	Competition autocode.CompetitionSche `json:"competition" gorm:"foreignkey:ID;references:CmpId"`
}
type EnterysFormInfo struct {
	global.GVA_MODEL
	Name        string                   `json:"name" form:"name" gorm:"column:name;comment:参赛表名;size:255;"`
	CmpId       *int                     `json:"cmpId" form:"cmpId" gorm:"column:cmp_id;comment:竞赛编号;size:10;"`
	PId         *int                     `json:"pId" form:"pId" gorm:"column:p_id;comment:项目编号，0代表没有;size:10;"`
	Rank        *int                     `json:"rank" form:"rank" gorm:"column:rank;comment:获奖级别;size:10;"`
	AchName     string                   `json:"achName" form:"achName" gorm:"column:ach_name;comment:获奖名称;size:255;"`
	Member      autocode.EntryMember     `json:"member" gorm:"foreignkey:FormId;references:ID"`
	Project     autocode.ProjectInfo     `json:"project,omitempty" gorm:"foreignkey:ID;references:PId"`
	Competition autocode.CompetitionSche `json:"competition" gorm:"foreignkey:ID;references:CmpId"`
}
