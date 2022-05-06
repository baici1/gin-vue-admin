package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AchievementService struct {
}

// CreateAchievement 创建Achievement记录
// Author [piexlmax](https://github.com/piexlmax)
func (achievementService *AchievementService) CreateAchievement(achievement autocode.Achievement) (err error) {
	err = global.GVA_DB.Create(&achievement).Error
	return err
}

// DeleteAchievement 删除Achievement记录
// Author [piexlmax](https://github.com/piexlmax)
func (achievementService *AchievementService) DeleteAchievement(achievement autocode.Achievement) (err error) {
	err = global.GVA_DB.Delete(&achievement).Error
	return err
}

// DeleteAchievementByIds 批量删除Achievement记录
// Author [piexlmax](https://github.com/piexlmax)
func (achievementService *AchievementService) DeleteAchievementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Achievement{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAchievement 更新Achievement记录
// Author [piexlmax](https://github.com/piexlmax)
func (achievementService *AchievementService) UpdateAchievement(achievement autocode.Achievement) (err error) {
	err = global.GVA_DB.Save(&achievement).Error
	return err
}

// GetAchievement 根据id获取Achievement记录
// Author [piexlmax](https://github.com/piexlmax)
func (achievementService *AchievementService) GetAchievement(id uint) (err error, achievement autocode.Achievement) {
	err = global.GVA_DB.Where("id = ?", id).First(&achievement).Error
	return
}

// GetAchievementInfoList 分页获取Achievement记录
// Author [piexlmax](https://github.com/piexlmax)
func (achievementService *AchievementService) GetAchievementInfoList(info autoCodeReq.AchievementSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Achievement{})
	var achievements []autocode.Achievement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Match != nil {
		db = db.Where("match = ?", info.Match)
	}
	if info.Rank != nil {
		db = db.Where("rank = ?", info.Rank)
	}
	if info.FormId != nil {
		db = db.Where("form_id=?", info.FormId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&achievements).Error
	return err, achievements, total
}
