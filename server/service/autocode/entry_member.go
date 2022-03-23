package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EntryMemberService struct {
}

// CreateEntryMember 创建EntryMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryMemberService *EntryMemberService) CreateEntryMember(entryMember autocode.EntryMember) (err error) {
	err = global.GVA_DB.Create(&entryMember).Error
	return err
}

// DeleteEntryMember 删除EntryMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryMemberService *EntryMemberService) DeleteEntryMember(entryMember autocode.EntryMember) (err error) {
	err = global.GVA_DB.Delete(&entryMember).Error
	return err
}

// DeleteEntryMemberByIds 批量删除EntryMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryMemberService *EntryMemberService) DeleteEntryMemberByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.EntryMember{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEntryMember 更新EntryMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryMemberService *EntryMemberService) UpdateEntryMember(entryMember autocode.EntryMember) (err error) {
	err = global.GVA_DB.Save(&entryMember).Error
	return err
}

// GetEntryMember 根据id获取EntryMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryMemberService *EntryMemberService) GetEntryMember(id uint) (err error, entryMember autocode.EntryMember) {
	err = global.GVA_DB.Where("id = ?", id).First(&entryMember).Error
	return
}

// GetEntryMemberInfoList 分页获取EntryMember记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryMemberService *EntryMemberService) GetEntryMemberInfoList(info autoCodeReq.EntryMemberSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.EntryMember{})
	var entryMembers []autocode.EntryMember
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Identify != nil {
		db = db.Where("identify = ?", info.Identify)
	}
	db = db.Where("form_id=?", info.FormId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&entryMembers).Error
	return err, entryMembers, total
}
