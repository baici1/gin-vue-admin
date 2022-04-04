package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	autocodeRes "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	sys "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

type UserInfoApi struct {
}

var userInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.UserInfoService

// CreateUserInfo 创建UserInfo
// @Tags UserInfo
// @Summary 创建UserInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.UserInfo true "创建UserInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userInfo/createUserInfo [post]
func (userInfoApi *UserInfoApi) CreateUserInfo(c *gin.Context) {
	var userInfo autocode.UserInfo
	_ = c.ShouldBindJSON(&userInfo)
	if err := userInfoService.CreateUserInfo(userInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserInfo 删除UserInfo
// @Tags UserInfo
// @Summary 删除UserInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.UserInfo true "删除UserInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userInfo/deleteUserInfo [delete]
func (userInfoApi *UserInfoApi) DeleteUserInfo(c *gin.Context) {
	var userInfo autocode.UserInfo
	_ = c.ShouldBindJSON(&userInfo)
	if err := userInfoService.DeleteUserInfo(userInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserInfoByIds 批量删除UserInfo
// @Tags UserInfo
// @Summary 批量删除UserInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userInfo/deleteUserInfoByIds [delete]
func (userInfoApi *UserInfoApi) DeleteUserInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := userInfoService.DeleteUserInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserInfo 更新UserInfo
// @Tags UserInfo
// @Summary 更新UserInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.UserInfo true "更新UserInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userInfo/updateUserInfo [put]
func (userInfoApi *UserInfoApi) UpdateUserInfo(c *gin.Context) {
	var userInfo autocode.UserInfo
	_ = c.ShouldBindJSON(&userInfo)
	if err := userInfoService.UpdateUserInfo(userInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserInfo 用id查询UserInfo
// @Tags UserInfo
// @Summary 用id查询UserInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.UserInfo true "用id查询UserInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userInfo/findUserInfo [get]
func (userInfoApi *UserInfoApi) FindUserInfo(c *gin.Context) {
	var userInfo autocode.UserInfo
	_ = c.ShouldBindQuery(&userInfo)
	if err, reuserInfo := userInfoService.GetUserInfo(userInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserInfo": reuserInfo}, c)
	}
}

// GetUserInfoList 分页获取UserInfo列表
// @Tags UserInfo
// @Summary 分页获取UserInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.UserInfoSearch true "分页获取UserInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userInfo/getUserInfoList [get]
func (userInfoApi *UserInfoApi) GetUserInfoList(c *gin.Context) {
	var pageInfo autocodeReq.UserInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := userInfoService.GetUserInfoInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (userInfoApi *UserInfoApi) CreateUserByRegister(c *gin.Context) {
	var param autocodeReq.UserRegister
	err := c.ShouldBind(&param)
	//参数校验
	if err != nil {
		response.ValidatorError(err, c)
		return
	}
	//校验码验证：现阶段先不放
	if store.Verify(param.CaptchaId, param.Captcha, true) {
		if err = userInfoService.CreateUserByRegister(autocode.UserInfo{
			Password: param.Password,
			Phone:    param.Phone,
			Identity: param.Identity,
		}); err != nil {
			global.GVA_LOG.Error(param.Phone+"注册失败！", zap.Error(err))
			response.FailWithMessage("注册失败！", c)
		} else {
			response.OkWithMessage("注册成功", c)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}

}

func (userInfoApi *UserInfoApi) UserToLogin(c *gin.Context) {
	var param autocodeReq.UserLogin
	err := c.ShouldBind(&param)
	//参数校验
	if err != nil {
		response.ValidatorError(err, c)
		return
	}
	//校验码验证：现阶段先不放
	if store.Verify(param.CaptchaId, param.Captcha, true) {
		if err, user := userInfoService.UserToLogin(autocode.UserInfo{
			Password: param.Password,
			Phone:    param.Phone,
		}); err != nil {
			global.GVA_LOG.Error("登陆失败!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			//签发token
			userInfoApi.tokenNext(c, *user)
			//response.OkWithData(user, c)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

func (userInfoApi *UserInfoApi) tokenNext(c *gin.Context, user autocode.UserInfo) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		ID:          user.ID,
		Username:    user.Phone,
		AuthorityId: strconv.Itoa(user.Identity),
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(autocodeRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if err, jwtStr := jwtService.GetRedisJWT(user.Phone); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Phone); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(autocodeRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT sys.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Phone); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(autocodeRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

func (userInfoApi *UserInfoApi) GetInfoByPhone(c *gin.Context) {
	var param = struct {
		Phone string `json:"phone" form:"phone" `
	}{}
	if err := c.ShouldBind(&param); err != nil {
		response.ValidatorError(err, c)
		return
	} else {
		if err, list := userInfoService.GetInfoByPhone(param.Phone); err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			response.FailWithMessage("获取失败", c)
		} else {
			response.OkWithData(list, c)
		}
	}
}
