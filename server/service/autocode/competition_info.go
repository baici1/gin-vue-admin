package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	autocodeRes "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
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
func (competitionInfoService *CompetitionInfoService) DeleteCompetitionInfo(competitionInfo autocode.CompetitionInfo) (err error) {
	err = global.GVA_DB.Delete(&competitionInfo).Error
	return err
}

// DeleteCompetitionInfoByIds 批量删除CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService) DeleteCompetitionInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CompetitionInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCompetitionInfo 更新CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService) UpdateCompetitionInfo(competitionInfo autocode.CompetitionInfo) (err error) {
	err = global.GVA_DB.Save(&competitionInfo).Error
	return err
}

// GetCompetitionInfo 根据id获取CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService) GetCompetitionInfo(id uint) (err error, competitionInfo autocode.CompetitionInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&competitionInfo).Error
	return
}

// GetCompetitionInfoInfoList 分页获取CompetitionInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionInfoService *CompetitionInfoService) GetCompetitionInfoInfoList(info autoCodeReq.CompetitionInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CompetitionInfo{})
	var competitionInfos []autocode.CompetitionInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CName != "" {
		db = db.Where("c_name LIKE ?", "%"+info.CName+"%")
	}
	if info.CType != nil {
		db = db.Where("c_type = ?", info.CType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&competitionInfos).Error
	return err, competitionInfos, total
}

func (competitionInfoService *CompetitionInfoService) GetComSelectList() (error, []autocodeRes.ComSelectList) {
	data := make([]autocodeRes.ComSelectList, 0)
	err := global.GVA_DB.Table(autocode.CompetitionInfo{}.TableName()).Select("competition_info.id,competition_info.c_name,competition_sche.c_id,competition_sche.id as s_id,competition_sche.level,competition_sche.version,competition_sche.version as f_version").Joins("left join competition_sche on competition_info.id=competition_sche.c_id").Order("competition_sche.start_time desc").Scan(&data).Error
	return err, data
}
