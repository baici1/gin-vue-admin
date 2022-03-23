package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CompanyInfoService struct {
}

// CreateCompanyInfo 创建CompanyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *CompanyInfoService) CreateCompanyInfo(companyInfo autocode.CompanyInfo) (err error) {
	err = global.GVA_DB.Create(&companyInfo).Error
	return err
}

// DeleteCompanyInfo 删除CompanyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *CompanyInfoService)DeleteCompanyInfo(companyInfo autocode.CompanyInfo) (err error) {
	err = global.GVA_DB.Delete(&companyInfo).Error
	return err
}

// DeleteCompanyInfoByIds 批量删除CompanyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *CompanyInfoService)DeleteCompanyInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CompanyInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCompanyInfo 更新CompanyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *CompanyInfoService)UpdateCompanyInfo(companyInfo autocode.CompanyInfo) (err error) {
	err = global.GVA_DB.Save(&companyInfo).Error
	return err
}

// GetCompanyInfo 根据id获取CompanyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *CompanyInfoService)GetCompanyInfo(id uint) (err error, companyInfo autocode.CompanyInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&companyInfo).Error
	return
}

// GetCompanyInfoInfoList 分页获取CompanyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (companyInfoService *CompanyInfoService)GetCompanyInfoInfoList(info autoCodeReq.CompanyInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.CompanyInfo{})
    var companyInfos []autocode.CompanyInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&companyInfos).Error
	return err, companyInfos, total
}
