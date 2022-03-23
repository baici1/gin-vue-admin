package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type EntryFormService struct {
}

// CreateEntryForm 创建EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService) CreateEntryForm(entryForm autocode.EntryForm) (err error) {
	err = global.GVA_DB.Create(&entryForm).Error
	return err
}

// DeleteEntryForm 删除EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService)DeleteEntryForm(entryForm autocode.EntryForm) (err error) {
	err = global.GVA_DB.Delete(&entryForm).Error
	return err
}

// DeleteEntryFormByIds 批量删除EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService)DeleteEntryFormByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.EntryForm{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEntryForm 更新EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService)UpdateEntryForm(entryForm autocode.EntryForm) (err error) {
	err = global.GVA_DB.Save(&entryForm).Error
	return err
}

// GetEntryForm 根据id获取EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService)GetEntryForm(id uint) (err error, entryForm autocode.EntryForm) {
	err = global.GVA_DB.Where("id = ?", id).First(&entryForm).Error
	return
}

// GetEntryFormInfoList 分页获取EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService)GetEntryFormInfoList(info autoCodeReq.EntryFormSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.EntryForm{})
    var entryForms []autocode.EntryForm
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.CmpId != nil {
        db = db.Where("cmp_id = ?",info.CmpId)
    }
    if info.Rank != nil {
        db = db.Where("rank = ?",info.Rank)
    }
    if info.AchName != "" {
        db = db.Where("ach_name LIKE ?","%"+ info.AchName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&entryForms).Error
	return err, entryForms, total
}
