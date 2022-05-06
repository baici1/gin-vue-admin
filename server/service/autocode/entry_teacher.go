package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EntryTeacherService struct {
}

// CreateEntryTeacher 创建EntryTeacher记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryTeacherService *EntryTeacherService) CreateEntryTeacher(entryTeacher autocode.EntryTeacher) (err error) {
	err = global.GVA_DB.Create(&entryTeacher).Error
	return err
}

// DeleteEntryTeacher 删除EntryTeacher记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryTeacherService *EntryTeacherService) DeleteEntryTeacher(entryTeacher autocode.EntryTeacher) (err error) {
	err = global.GVA_DB.Delete(&entryTeacher).Error
	return err
}

// DeleteEntryTeacherByIds 批量删除EntryTeacher记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryTeacherService *EntryTeacherService) DeleteEntryTeacherByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.EntryTeacher{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEntryTeacher 更新EntryTeacher记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryTeacherService *EntryTeacherService) UpdateEntryTeacher(entryTeacher autocode.EntryTeacher) (err error) {
	err = global.GVA_DB.Save(&entryTeacher).Error
	return err
}

// GetEntryTeacher 根据id获取EntryTeacher记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryTeacherService *EntryTeacherService) GetEntryTeacher(id uint) (err error, entryTeacher autocode.EntryTeacher) {
	err = global.GVA_DB.Where("id = ?", id).First(&entryTeacher).Error
	return
}

// GetEntryTeacherInfoList 分页获取EntryTeacher记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryTeacherService *EntryTeacherService) GetEntryTeacherInfoList(info autoCodeReq.EntryTeacherSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.EntryTeacher{})
	var entryTeachers []autocode.EntryTeacher
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("form_id=?", info.FormId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&entryTeachers).Error
	return err, entryTeachers, total
}
