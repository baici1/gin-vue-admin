package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type StudentInfoService struct {
}

// CreateStudentInfo 创建StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) CreateStudentInfo(studentInfo autocode.StudentInfo) (err error) {
	err = global.GVA_DB.Create(&studentInfo).Error
	return err
}

// DeleteStudentInfo 删除StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) DeleteStudentInfo(studentInfo autocode.StudentInfo) (err error) {
	err = global.GVA_DB.Delete(&studentInfo).Error
	return err
}

// DeleteStudentInfoByIds 批量删除StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) DeleteStudentInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.StudentInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateStudentInfo 更新StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) UpdateStudentInfo(studentInfo autocode.StudentInfo) (err error) {
	err = global.GVA_DB.Save(&studentInfo).Error
	return err
}

// GetStudentInfo 根据id获取StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) GetStudentInfo(id uint) (err error, studentInfo autocode.StudentInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&studentInfo).Error
	return
}

// GetStudentInfoInfoList 分页获取StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) GetStudentInfoInfoList(info autoCodeReq.StudentInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.StudentInfo{})
	var studentInfos []autocode.StudentInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StudentId != "" {
		db = db.Where("student_id LIKE ?", "%"+info.StudentId+"%")
	}
	if info.RealName != "" {
		db = db.Where("real_name LIKE ?", "%"+info.RealName+"%")
	}
	if info.Gender != nil {
		db = db.Where("gender = ?", info.Gender)
	}
	if info.Grade != "" {
		db = db.Where("grade = ?", info.Grade)
	}
	if info.Department != "" {
		db = db.Where("department LIKE ?", "%"+info.Department+"%")
	}
	if info.Major != "" {
		db = db.Where("major LIKE ?", "%"+info.Major+"%")
	}
	if info.ClassNum != "" {
		db = db.Where("class_num LIKE ?", "%"+info.ClassNum+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&studentInfos).Error
	return err, studentInfos, total
}
