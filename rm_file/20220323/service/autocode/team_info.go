package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type TeamInfoService struct {
}

// CreateTeamInfo 创建TeamInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInfoService *TeamInfoService) CreateTeamInfo(teamInfo autocode.TeamInfo) (err error) {
	err = global.GVA_DB.Create(&teamInfo).Error
	return err
}

// DeleteTeamInfo 删除TeamInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInfoService *TeamInfoService)DeleteTeamInfo(teamInfo autocode.TeamInfo) (err error) {
	err = global.GVA_DB.Delete(&teamInfo).Error
	return err
}

// DeleteTeamInfoByIds 批量删除TeamInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInfoService *TeamInfoService)DeleteTeamInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.TeamInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTeamInfo 更新TeamInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInfoService *TeamInfoService)UpdateTeamInfo(teamInfo autocode.TeamInfo) (err error) {
	err = global.GVA_DB.Save(&teamInfo).Error
	return err
}

// GetTeamInfo 根据id获取TeamInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInfoService *TeamInfoService)GetTeamInfo(id uint) (err error, teamInfo autocode.TeamInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&teamInfo).Error
	return
}

// GetTeamInfoInfoList 分页获取TeamInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teamInfoService *TeamInfoService)GetTeamInfoInfoList(info autoCodeReq.TeamInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.TeamInfo{})
    var teamInfos []autocode.TeamInfo
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&teamInfos).Error
	return err, teamInfos, total
}
