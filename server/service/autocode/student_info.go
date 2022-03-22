package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
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

// UpdateStudentInfo 更新StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) UpdateStudentInfo(studentInfo autocode.StudentInfo) (err error) {
	err = global.GVA_DB.Save(&studentInfo).Error
	return err
}

// GetStudentInfo 根据id获取StudentInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (studentInfoService *StudentInfoService) GetStudentInfo(id int) (err error, studentInfo autocode.StudentInfo) {
	err = global.GVA_DB.Where("u_id = ?", id).First(&studentInfo).Error
	return
}
