package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CompetitionInfoService struct {
}

// CreateCompetitionInfo 创建CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService) CreateCompetitionInfo(competitionInfo autocode.CompetitionInfo) (err error) {
	err = global.GVA_DB.Create(&competitionInfo).Error
	return err
}

// DeleteCompetitionInfo 删除CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService)DeleteCompetitionInfo(competitionInfo autocode.CompetitionInfo) (err error) {
	err = global.GVA_DB.Delete(&competitionInfo).Error
	return err
}

// DeleteCompetitionInfoByIds 批量删除CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService)DeleteCompetitionInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CompetitionInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCompetitionInfo 更新CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService)UpdateCompetitionInfo(competitionInfo autocode.CompetitionInfo) (err error) {
	err = global.GVA_DB.Save(&competitionInfo).Error
	return err
}

// GetCompetitionInfo 根据id获取CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService)GetCompetitionInfo(id uint) (err error, competitionInfo autocode.CompetitionInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&competitionInfo).Error
	return
}

// GetCompetitionInfoInfoList 分页获取CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService)GetCompetitionInfoInfoList(info autoCodeReq.CompetitionInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.CompetitionInfo{})
    var competitionInfos []autocode.CompetitionInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.CName != "" {
        db = db.Where("c_name LIKE ?","%"+ info.CName+"%")
    }
    if info.CType != nil {
        db = db.Where("c_type = ?",info.CType)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&competitionInfos).Error
	return err, competitionInfos, total
}
