package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	autocodeRes "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
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
func (entryFormService *EntryFormService) DeleteEntryForm(entryForm autocode.EntryForm) (err error) {
	err = global.GVA_DB.Delete(&entryForm).Error
	return err
}

// DeleteEntryFormByIds 批量删除EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService) DeleteEntryFormByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.EntryForm{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEntryForm 更新EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService) UpdateEntryForm(entryForm autocode.EntryForm) (err error) {
	err = global.GVA_DB.Save(&entryForm).Error
	return err
}

// GetEntryForm 根据id获取EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService) GetEntryForm(id uint) (err error, entryForm autocode.EntryForm) {
	err = global.GVA_DB.Where("id = ?", id).First(&entryForm).Error
	return
}

// GetEntryFormInfoList 分页获取EntryForm记录
// Author [piexlmax](https://github.com/piexlmax)
func (entryFormService *EntryFormService) GetEntryFormInfoList(info autoCodeReq.EntryFormSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.EntryForm{})
	var entryForms []autocode.EntryForm
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.CmpId != nil {
		db = db.Where("cmp_id = ?", info.CmpId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&entryForms).Error
	return err, entryForms, total
}

func (entryFormService *EntryFormService) GetAllEntryFormDetailInfo(uid int) (error, []autocodeRes.EnterysFormInfo) {
	var data = make([]autocodeRes.EnterysFormInfo, 0)
	//根据uid获取用户参与的所有的参赛表 id
	var forms = make([]int, 0)
	err := global.GVA_DB.Model(&autocode.EntryMember{}).Select("form_id").Where("u_id=?", uid).Scan(&forms).Error
	if err != nil {
		return err, nil
	}
	//根据formid获取每个form的相关信息
	db := global.GVA_DB.Table(autocode.EntryForm{}.TableName())
	//这里可以添加一些条件，展示不添加
	err = db.Where("id in ?", forms).Preload("Competition").Preload("Competition.BaseInfo").Preload("Project").Preload("Member", "u_id=?", uid).Preload("Member.BaseInfo").Find(&data).Error
	return err, data
}
func (entryFormService *EntryFormService) GetEntryFormDetailInfo(fid int) (error, *autocodeRes.EnteryFormDetail) {
	data := new(autocodeRes.EnteryFormDetail)
	//根据formid获取每个form的相关信息
	db := global.GVA_DB.Table(autocode.EntryForm{}.TableName())
	//这里可以添加一些条件，展示不添加
	err := db.Where("id=?", fid).Preload("Competition").Preload("Competition.BaseInfo").Find(data).Error
	return err, data
}

func (entryFormService *EntryFormService) CreateEntryFormByUser(param autoCodeReq.CreateEntryForm) error {
	tx := global.GVA_DB.Begin()
	var err error
	//创建参赛表
	err = tx.Model(&autocode.EntryForm{}).Select("PId", "CmpId", "Name").Create(&param).Error
	param.Status.FormId = int(param.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	//创建队长
	param.Status.Order = 1
	err = tx.Model(&autocode.EntryMember{}).Create(&param.Status).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (entryFormService *EntryFormService) UpdateEntryFormByUser(param autoCodeReq.UpdateEntryForm) error {
	tx := global.GVA_DB.Begin()
	var err error
	//更新项目或者创建项目
	if param.PId == 0 {
		if param.Project != nil {
			err = tx.Model(&autocode.ProjectInfo{}).Create(param.Project).Error
			param.PId = int(param.Project.ID)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	} else {
		err = tx.Save(param.Project).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//更新参赛表一些信息
	err = tx.Model(&autocode.EntryForm{}).Where("id=?", param.ID).Updates(autocode.EntryForm{
		Name: param.Name,
		PId:  &param.PId,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
