package autocode

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type UserInfoService struct {
}

// CreateUserInfo 创建UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) CreateUserInfo(userInfo autocode.UserInfo) (err error) {
	var user autocode.UserInfo
	if errors.Is(global.GVA_DB.Where("phone=?", userInfo.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户已注册")
	}
	userInfo.Password = utils.MD5V([]byte(userInfo.Password))
	err = global.GVA_DB.Create(&userInfo).Error
	return err
}

// DeleteUserInfo 删除UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) DeleteUserInfo(userInfo autocode.UserInfo) (err error) {
	err = global.GVA_DB.Delete(&userInfo).Error
	return err
}

// DeleteUserInfoByIds 批量删除UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) DeleteUserInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.UserInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserInfo 更新UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) UpdateUserInfo(userInfo autocode.UserInfo) (err error) {
	userInfo.Password = utils.MD5V([]byte(userInfo.Password))
	err = global.GVA_DB.Save(&userInfo).Error
	return err
}

// GetUserInfo 根据id获取UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) GetUserInfo(id uint) (err error, userInfo autocode.UserInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&userInfo).Error
	return
}

// GetUserInfoInfoList 分页获取UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) GetUserInfoInfoList(info autoCodeReq.UserInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.UserInfo{})
	var userInfos []autocode.UserInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userInfos).Error
	return err, userInfos, total
}
