package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TeamMemberService struct {
}

// CreateTeamMember 创建TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) CreateTeamMember(teamMember autocode.TeamMember) (err error) {
	err = global.GVA_DB.Create(&teamMember).Error
	return err
}

// DeleteTeamMember 删除TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) DeleteTeamMember(teamMember autocode.TeamMember) (err error) {
	err = global.GVA_DB.Delete(&teamMember).Error
	return err
}

// DeleteTeamMemberByIds 批量删除TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) DeleteTeamMemberByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.TeamMember{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeamMember 更新TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) UpdateTeamMember(teamMember autocode.TeamMember) (err error) {
	err = global.GVA_DB.Save(&teamMember).Error
	return err
}

// GetTeamMember 根据id获取TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) GetTeamMember(id uint) (err error, teamMember autocode.TeamMember) {
	err = global.GVA_DB.Where("id = ?", id).First(&teamMember).Error
	return
}

// GetTeamMemberInfoList 分页获取TeamMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamMemberService *TeamMemberService) GetTeamMemberInfoList(info autoCodeReq.TeamMemberSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.TeamMember{})
	var teamMembers []autocode.TeamMember
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("team_id=?", info.TeamId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&teamMembers).Error
	return err, teamMembers, total
}
