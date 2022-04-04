package autocode

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type UserInfoService struct {
}

// CreateUserInfo 创建UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) CreateUserInfo(userInfo autocode.UserInfo) (err error) {
	var user autocode.UserInfo
	if !errors.Is(global.GVA_DB.Where("phone=?", userInfo.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户已注册")
	}
	userInfo.Password = utils.MD5V([]byte(userInfo.Password))
	err = global.GVA_DB.Create(&userInfo).Error
	return err
}

// DeleteUserInfo 删除UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) DeleteUserInfo(userInfo autocode.UserInfo) (err error) {
	//使用事务进行连接删除
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	//删除个人基本信息
	if err = tx.Delete(&userInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除个人的详情信息
	if userInfo.Identity == 1 {
		//删除老师信息
		if err = tx.Where("u_id=?", int(userInfo.ID)).Delete(&autocode.TeacherInfo{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		//删除学生信息
		if err = tx.Where("u_id=?", int(userInfo.ID)).Delete(&autocode.StudentInfo{}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
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

//CreateUserByRegister
/* @Description: 用户注册
 * @receiver userInfoService
 * @param register
 * @return err
 */
func (userInfoService *UserInfoService) CreateUserByRegister(userInfo autocode.UserInfo) (err error) {
	tx := global.GVA_DB.Begin()
	var user autocode.UserInfo
	//重复注册
	if !errors.Is(tx.Where("phone=?", userInfo.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户已注册")
	}
	//密码加密
	userInfo.Password = utils.MD5V([]byte(userInfo.Password))
	//判断身份:如果是学生选择直接审核通过
	if userInfo.Identity == 2019 {
		userInfo.Check = 1
	}
	err = global.GVA_DB.Create(&userInfo).Error
	if err != nil {
		tx.Rollback()
		return
	}
	//需要根据身份去创建一个基本的用户信息（老师与学生）
	if userInfo.Identity == 2019 {
		stu := autocode.StudentInfo{
			UId:      int(userInfo.ID),
			Nickname: userInfo.Phone,
			Avatar:   `http://project.yangdiy.cn/16483710248249a5f4881211ebb6edd017c2d2eca2.jpg`,
			RealName: userInfo.Phone,
		}
		err = global.GVA_DB.Create(&stu).Error
	} else {
		tea := autocode.TeacherInfo{
			UId:      int(userInfo.ID),
			Nickname: userInfo.Phone,
			Avatar:   `http://project.yangdiy.cn/16483710248249a5f4881211ebb6edd017c2d2eca2.jpg`,
			RealName: userInfo.Phone,
		}
		err = global.GVA_DB.Create(&tea).Error
	}
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (userInfoService *UserInfoService) UserToLogin(u autocode.UserInfo) (err error, userInter *autocode.UserInfo) {
	var user autocode.UserInfo
	//密码加密进行比较
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("phone=? and password=?", u.Phone, u.Password).Preload("Authority").First(&user).Error
	if user.Check != 1 {
		return errors.New("当前用户需要审核！请等待审核！"), nil
	}
	//如果可以找到
	if err == nil {
		var am system.SysMenu
		ferr := global.GVA_DB.First(&am, "name = ? AND authority_id = ?", user.Authority.DefaultRouter, user.Identity).Error
		if errors.Is(ferr, gorm.ErrRecordNotFound) {
			user.Authority.DefaultRouter = "404"
		}
	}
	return err, &user
}

func (userInfoService *UserInfoService) GetInfoByPhone(pre string) (err error, data autocode.UserInfo) {
	fmt.Println(pre)
	err = global.GVA_DB.Where("phone = ? ", pre).First(&data).Error
	return err, data
}
