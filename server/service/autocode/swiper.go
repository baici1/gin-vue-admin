package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type SwiperService struct {
}

// CreateSwiper 创建Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService) CreateSwiper(swiper autocode.Swiper) (err error) {
	err = global.GVA_DB.Create(&swiper).Error
	return err
}

// DeleteSwiper 删除Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService)DeleteSwiper(swiper autocode.Swiper) (err error) {
	err = global.GVA_DB.Delete(&swiper).Error
	return err
}

// DeleteSwiperByIds 批量删除Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService)DeleteSwiperByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Swiper{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSwiper 更新Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService)UpdateSwiper(swiper autocode.Swiper) (err error) {
	err = global.GVA_DB.Save(&swiper).Error
	return err
}

// GetSwiper 根据id获取Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService)GetSwiper(id uint) (err error, swiper autocode.Swiper) {
	err = global.GVA_DB.Where("id = ?", id).First(&swiper).Error
	return
}

// GetSwiperInfoList 分页获取Swiper记录
// Author [piexlmax](https://github.com/piexlmax)
func (swiperService *SwiperService)GetSwiperInfoList(info autoCodeReq.SwiperSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Swiper{})
    var swipers []autocode.Swiper
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.SwiperName != "" {
        db = db.Where("swiper_name LIKE ?","%"+ info.SwiperName+"%")
    }
    if info.IsShow != nil {
        db = db.Where("is_show = ?",info.IsShow)
    }
    if info.Group != nil {
        db = db.Where("group = ?",info.Group)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&swipers).Error
	return err, swipers, total
}
