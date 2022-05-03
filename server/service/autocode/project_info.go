package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProjectInfoService struct {
}

// CreateProjectInfo 创建ProjectInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectInfoService *ProjectInfoService) CreateProjectInfo(projectInfo autocode.ProjectInfo) (err error, ID int) {
	err = global.GVA_DB.Create(&projectInfo).Error
	return err, int(projectInfo.ID)
}

// DeleteProjectInfo 删除ProjectInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectInfoService *ProjectInfoService) DeleteProjectInfo(projectInfo autocode.ProjectInfo) (err error) {
	err = global.GVA_DB.Delete(&projectInfo).Error
	return err
}

// DeleteProjectInfoByIds 批量删除ProjectInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectInfoService *ProjectInfoService) DeleteProjectInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ProjectInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProjectInfo 更新ProjectInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectInfoService *ProjectInfoService) UpdateProjectInfo(projectInfo autocode.ProjectInfo) (err error) {
	err = global.GVA_DB.Save(&projectInfo).Error
	return err
}

// GetProjectInfo 根据id获取ProjectInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectInfoService *ProjectInfoService) GetProjectInfo(id uint) (err error, projectInfo autocode.ProjectInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&projectInfo).Error
	return
}

// GetProjectInfoInfoList 分页获取ProjectInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectInfoService *ProjectInfoService) GetProjectInfoInfoList(info autoCodeReq.ProjectInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.ProjectInfo{})
	var projectInfos []autocode.ProjectInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ProjectName != "" {
		db = db.Where("project_name LIKE ?", "%"+info.ProjectName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&projectInfos).Error
	return err, projectInfos, total
}
