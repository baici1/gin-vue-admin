package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TeacherInfoService struct {
}

// CreateTeacherInfo 创建TeacherInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherInfoService *TeacherInfoService) CreateTeacherInfo(teacherInfo autocode.TeacherInfo) (err error) {
	err = global.GVA_DB.Create(&teacherInfo).Error
	return err
}

// DeleteTeacherInfo 删除TeacherInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherInfoService *TeacherInfoService) DeleteTeacherInfo(teacherInfo autocode.TeacherInfo) (err error) {
	err = global.GVA_DB.Delete(&teacherInfo).Error
	return err
}

// DeleteTeacherInfoByIds 批量删除TeacherInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherInfoService *TeacherInfoService) DeleteTeacherInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.TeacherInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTeacherInfo 更新TeacherInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherInfoService *TeacherInfoService) UpdateTeacherInfo(teacherInfo autocode.TeacherInfo) (err error) {
	err = global.GVA_DB.Save(&teacherInfo).Error
	return err
}

// GetTeacherInfo 根据id获取TeacherInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherInfoService *TeacherInfoService) GetTeacherInfo(id uint) (err error, teacherInfo autocode.TeacherInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&teacherInfo).Error
	return
}

// GetTeacherInfoInfoList 分页获取TeacherInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (teacherInfoService *TeacherInfoService) GetTeacherInfoInfoList(info autoCodeReq.TeacherInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.TeacherInfo{})
	var teacherInfos []autocode.TeacherInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.PersonnelId != "" {
		db = db.Where("personnel_id = ?", info.PersonnelId)
	}
	if info.OfficeId != "" {
		db = db.Where("office_id = ?", info.OfficeId)
	}
	if info.FinancialId != "" {
		db = db.Where("financial_id = ?", info.FinancialId)
	}
	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	if info.RealName != "" {
		db = db.Where("real_name LIKE ?", "%"+info.RealName+"%")
	}
	if info.Gender != nil {
		db = db.Where("gender = ?", info.Gender)
	}
	if info.Department != "" {
		db = db.Where("department LIKE ?", "%"+info.Department+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&teacherInfos).Error
	return err, teacherInfos, total
}
