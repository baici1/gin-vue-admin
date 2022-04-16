package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type TeacherRecruitService struct {
}

// CreateTeacherRecruit 创建TeacherRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherRecruitService *TeacherRecruitService) CreateTeacherRecruit(teacherRecruit autocode.TeacherRecruit) (err error) {
	err = global.GVA_DB.Create(&teacherRecruit).Error
	return err
}

// DeleteTeacherRecruit 删除TeacherRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherRecruitService *TeacherRecruitService)DeleteTeacherRecruit(teacherRecruit autocode.TeacherRecruit) (err error) {
	err = global.GVA_DB.Delete(&teacherRecruit).Error
	return err
}

// DeleteTeacherRecruitByIds 批量删除TeacherRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherRecruitService *TeacherRecruitService)DeleteTeacherRecruitByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.TeacherRecruit{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTeacherRecruit 更新TeacherRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherRecruitService *TeacherRecruitService)UpdateTeacherRecruit(teacherRecruit autocode.TeacherRecruit) (err error) {
	err = global.GVA_DB.Save(&teacherRecruit).Error
	return err
}

// GetTeacherRecruit 根据id获取TeacherRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherRecruitService *TeacherRecruitService)GetTeacherRecruit(id uint) (err error, teacherRecruit autocode.TeacherRecruit) {
	err = global.GVA_DB.Where("id = ?", id).First(&teacherRecruit).Error
	return
}

// GetTeacherRecruitInfoList 分页获取TeacherRecruit记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherRecruitService *TeacherRecruitService)GetTeacherRecruitInfoList(info autoCodeReq.TeacherRecruitSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.TeacherRecruit{})
    var teacherRecruits []autocode.TeacherRecruit
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.TId != nil {
        db = db.Where("t_id = ?",info.TId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&teacherRecruits).Error
	return err, teacherRecruits, total
}
