package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type StudentRecruitService struct {
}

// CreateStudentRecruit 创建StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService) CreateStudentRecruit(studentRecruit autocode.StudentRecruit) (err error) {
	err = global.GVA_DB.Create(&studentRecruit).Error
	return err
}

// DeleteStudentRecruit 删除StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService)DeleteStudentRecruit(studentRecruit autocode.StudentRecruit) (err error) {
	err = global.GVA_DB.Delete(&studentRecruit).Error
	return err
}

// DeleteStudentRecruitByIds 批量删除StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService)DeleteStudentRecruitByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.StudentRecruit{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStudentRecruit 更新StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService)UpdateStudentRecruit(studentRecruit autocode.StudentRecruit) (err error) {
	err = global.GVA_DB.Save(&studentRecruit).Error
	return err
}

// GetStudentRecruit 根据id获取StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService)GetStudentRecruit(id uint) (err error, studentRecruit autocode.StudentRecruit) {
	err = global.GVA_DB.Where("id = ?", id).First(&studentRecruit).Error
	return
}

// GetStudentRecruitInfoList 分页获取StudentRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentRecruitService *StudentRecruitService)GetStudentRecruitInfoList(info autoCodeReq.StudentRecruitSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.StudentRecruit{})
    var studentRecruits []autocode.StudentRecruit
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.ComId != nil {
        db = db.Where("com_id = ?",info.ComId)
    }
    if info.UId != nil {
        db = db.Where("u_id = ?",info.UId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&studentRecruits).Error
	return err, studentRecruits, total
}
