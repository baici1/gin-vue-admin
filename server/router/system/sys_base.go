package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	userRouter := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	var userInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.UserInfoApi
	{
		userRouter.POST("register", userInfoApi.CreateUserByRegister)
		userRouter.POST("login", userInfoApi.UserToLogin)
	}
	return baseRouter
}
