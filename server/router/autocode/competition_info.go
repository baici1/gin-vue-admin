package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompetitionInfoRouter struct {
}

// InitCompetitionInfoRouter 初始化 CompetitionInfo 路由信息
func (s *CompetitionInfoRouter) InitCompetitionInfoRouter(Router *gin.RouterGroup) {
	competitionInfoRouter := Router.Group("competitionInfo").Use(middleware.OperationRecord())
	competitionInfoRouterWithoutRecord := Router.Group("competitionInfo")
	var competitionInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.CompetitionInfoApi
	{
		competitionInfoRouter.POST("createCompetitionInfo", competitionInfoApi.CreateCompetitionInfo)   // 新建CompetitionInfo
		competitionInfoRouter.DELETE("deleteCompetitionInfo", competitionInfoApi.DeleteCompetitionInfo) // 删除CompetitionInfo
		competitionInfoRouter.DELETE("deleteCompetitionInfoByIds", competitionInfoApi.DeleteCompetitionInfoByIds) // 批量删除CompetitionInfo
		competitionInfoRouter.PUT("updateCompetitionInfo", competitionInfoApi.UpdateCompetitionInfo)    // 更新CompetitionInfo
	}
	{
		competitionInfoRouterWithoutRecord.GET("findCompetitionInfo", competitionInfoApi.FindCompetitionInfo)        // 根据ID获取CompetitionInfo
		competitionInfoRouterWithoutRecord.GET("getCompetitionInfoList", competitionInfoApi.GetCompetitionInfoList)  // 获取CompetitionInfo列表
	}
}
