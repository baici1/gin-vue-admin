package autocode

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
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

func (teamMemberService *TeamMemberService) CreateOwnTeamMember(member autoCodeReq.OwnTeamMember) (err error) {
	//phone存在，去user表获取id
	var user autocode.StudentInfo
	//保证用户已注册账号
	if errors.Is(global.GVA_DB.Where("phone=?", member.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("当前用户未注册账号！无法添加！")
	}
	fmt.Println("u_id", user.ID)
	//保证用户没有参与当前团队或者用户没有加入过团队
	//查询当前用户是否已经加入过团队
	err = global.GVA_DB.Where("u_id=?", user.ID).First(&autocode.TeamMember{}).Error
	//有三种可能，一种是没有加入过团队 err!=nil，一种是已经加入过团队 err==nil，还有一种是加入了但是被删除了
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//如果之前已经加入后又删除了
			if err := global.GVA_DB.Unscoped().Where("u_id=? and deleted_at IS NOT NULL", user.ID).First(&autocode.TeamMember{}).Error; err == nil {
				return global.GVA_DB.Exec("UPDATE team_member SET deleted_at=null WHERE u_id=?;", user.ID).Error
			}
			var param = autocode.TeamMember{
				TeamId:   int(member.TeamId),
				UId:      int(user.ID),
				Identify: member.Identify,
			}
			//创建团队成员
			return global.GVA_DB.Create(&param).Error
		} else {
			return err
		}
	} else {
		return errors.New("当前用户已加入团队！无法添加！")
	}
}

func (teamMemberService *TeamMemberService) GetTeamIDByUser(uid int) (error, int) {
	teamid := 0
	err := global.GVA_DB.Model(&autocode.TeamMember{}).Where("u_id=?", uid).Select("team_id").Find(&teamid).Error
	return err, teamid
}
