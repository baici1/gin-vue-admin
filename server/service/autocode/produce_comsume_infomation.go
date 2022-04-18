package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ProduceComsumeInfomationService struct {
}

// CreateProduceComsumeInfomation 创建ProduceComsumeInfomation记录
// Author [piexlmax](https://github.com/piexlmax)
func (produceComsumeInfomationService *ProduceComsumeInfomationService) CreateProduceComsumeInfomation(produceComsumeInfomation autocode.ProduceComsumeInfomation) (err error) {
	err = global.GVA_DB.Create(&produceComsumeInfomation).Error
	return err
}

// DeleteProduceComsumeInfomation 删除ProduceComsumeInfomation记录
// Author [piexlmax](https://github.com/piexlmax)
func (produceComsumeInfomationService *ProduceComsumeInfomationService)DeleteProduceComsumeInfomation(produceComsumeInfomation autocode.ProduceComsumeInfomation) (err error) {
	err = global.GVA_DB.Delete(&produceComsumeInfomation).Error
	return err
}

// DeleteProduceComsumeInfomationByIds 批量删除ProduceComsumeInfomation记录
// Author [piexlmax](https://github.com/piexlmax)
func (produceComsumeInfomationService *ProduceComsumeInfomationService)DeleteProduceComsumeInfomationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ProduceComsumeInfomation{},"id in ?",ids.Ids).Error
	return err
}

// UpdateProduceComsumeInfomation 更新ProduceComsumeInfomation记录
// Author [piexlmax](https://github.com/piexlmax)
func (produceComsumeInfomationService *ProduceComsumeInfomationService)UpdateProduceComsumeInfomation(produceComsumeInfomation autocode.ProduceComsumeInfomation) (err error) {
	err = global.GVA_DB.Save(&produceComsumeInfomation).Error
	return err
}

// GetProduceComsumeInfomation 根据id获取ProduceComsumeInfomation记录
// Author [piexlmax](https://github.com/piexlmax)
func (produceComsumeInfomationService *ProduceComsumeInfomationService)GetProduceComsumeInfomation(id uint) (err error, produceComsumeInfomation autocode.ProduceComsumeInfomation) {
	err = global.GVA_DB.Where("id = ?", id).First(&produceComsumeInfomation).Error
	return
}

// GetProduceComsumeInfomationInfoList 分页获取ProduceComsumeInfomation记录
// Author [piexlmax](https://github.com/piexlmax)
func (produceComsumeInfomationService *ProduceComsumeInfomationService)GetProduceComsumeInfomationInfoList(info autoCodeReq.ProduceComsumeInfomationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.ProduceComsumeInfomation{})
    var produceComsumeInfomations []autocode.ProduceComsumeInfomation
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&produceComsumeInfomations).Error
	return err, produceComsumeInfomations, total
}
