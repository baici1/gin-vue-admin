package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeamInfoRouter struct {
}

// InitTeamInfoRouter 初始化 TeamInfo 路由信息
func (s *TeamInfoRouter) InitTeamInfoRouter(Router *gin.RouterGroup) {
	teamInfoRouter := Router.Group("teamInfo").Use(middleware.OperationRecord())
	teamInfoRouterWithoutRecord := Router.Group("teamInfo")
	var teamInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.TeamInfoApi
	{
		teamInfoRouter.POST("createTeamInfo", teamInfoApi.CreateTeamInfo)   // 新建TeamInfo
		teamInfoRouter.DELETE("deleteTeamInfo", teamInfoApi.DeleteTeamInfo) // 删除TeamInfo
		teamInfoRouter.DELETE("deleteTeamInfoByIds", teamInfoApi.DeleteTeamInfoByIds) // 批量删除TeamInfo
		teamInfoRouter.PUT("updateTeamInfo", teamInfoApi.UpdateTeamInfo)    // 更新TeamInfo
	}
	{
		teamInfoRouterWithoutRecord.GET("findTeamInfo", teamInfoApi.FindTeamInfo)        // 根据ID获取TeamInfo
		teamInfoRouterWithoutRecord.GET("getTeamInfoList", teamInfoApi.GetTeamInfoList)  // 获取TeamInfo列表
	}
}
