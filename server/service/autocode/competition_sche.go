package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CompetitionScheService struct {
}

// CreateCompetitionSche 创建CompetitionSche记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionScheService *CompetitionScheService) CreateCompetitionSche(competitionSche autocode.CompetitionSche) (err error) {
	err = global.GVA_DB.Create(&competitionSche).Error
	return err
}

// DeleteCompetitionSche 删除CompetitionSche记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionScheService *CompetitionScheService) DeleteCompetitionSche(competitionSche autocode.CompetitionSche) (err error) {
	err = global.GVA_DB.Delete(&competitionSche).Error
	return err
}

// DeleteCompetitionScheByIds 批量删除CompetitionSche记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionScheService *CompetitionScheService) DeleteCompetitionScheByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CompetitionSche{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCompetitionSche 更新CompetitionSche记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionScheService *CompetitionScheService) UpdateCompetitionSche(competitionSche autocode.CompetitionSche) (err error) {
	err = global.GVA_DB.Save(&competitionSche).Error
	return err
}

// GetCompetitionSche 根据id获取CompetitionSche记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionScheService *CompetitionScheService) GetCompetitionSche(id uint) (err error, competitionSche autocode.CompetitionSche) {
	err = global.GVA_DB.Where("id = ?", id).First(&competitionSche).Error
	return
}

// GetCompetitionScheInfoList 分页获取CompetitionSche记录
// Author [piexlmax](https://github.com/piexlmax)
func (competitionScheService *CompetitionScheService) GetCompetitionScheInfoList(info autoCodeReq.CompetitionScheSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CompetitionSche{})
	var competitionSches []autocode.CompetitionSche
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Level != nil {
		db = db.Where("level = ?", info.Level)
	}
	if info.Version != nil {
		db = db.Where("version = ?", info.Version)
	}
	if info.CId != nil {
		db = db.Where("c_id=?", info.CId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&competitionSches).Error
	return err, competitionSches, total
}

func (competitionScheService *CompetitionScheService) GetCompetitionScheDetailList(info autoCodeReq.CompetitionDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.CompetitionSche{})
	var competitionSches []autocode.CompetitionSche
	db = db.Joins("left join competition_info on competition_sche.c_id=competition_info.id")
	//可加条件
	//报名未开始
	if info.Status == 1 {
		db = db.Where("competition_sche.start_time > now() ")
	}
	//报名种
	if info.Status == 2 {
		db = db.Where("competition_sche.start_time < now() and competition_sche.end_time > now() ")
	}
	//准备种
	if info.Status == 3 {
		db = db.Where("competition_sche.end_time < now() and now() < competition_sche.r_start_time   ")
	}
	//比赛进行种
	if info.Status == 4 {
		db = db.Where("competition_sche.r_start_time < now() and now() < competition_sche.r_end_time ")
	}
	//比赛结束
	if info.Status == 5 {
		db = db.Where("competition_sche.r_end_time < now() ")
	}
	db = db.Where("competition_info.c_name like ?", "%"+info.SearchInfo+"%")
	err = db.Limit(limit).Offset(offset).Preload("BaseInfo", "c_name like ?", "%"+info.SearchInfo+"%").Find(&competitionSches).Error
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	return err, competitionSches, total
}
