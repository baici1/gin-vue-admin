package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompetitionScheRouter struct {
}

// InitCompetitionScheRouter 初始化 CompetitionSche 路由信息
func (s *CompetitionScheRouter) InitCompetitionScheRouter(Router *gin.RouterGroup) {
	competitionScheRouter := Router.Group("competitionSche").Use(middleware.OperationRecord())
	competitionScheRouterWithoutRecord := Router.Group("competitionSche")
	var competitionScheApi = v1.ApiGroupApp.AutoCodeApiGroup.CompetitionScheApi
	{
		competitionScheRouter.POST("createCompetitionSche", competitionScheApi.CreateCompetitionSche)   // 新建CompetitionSche
		competitionScheRouter.DELETE("deleteCompetitionSche", competitionScheApi.DeleteCompetitionSche) // 删除CompetitionSche
		competitionScheRouter.DELETE("deleteCompetitionScheByIds", competitionScheApi.DeleteCompetitionScheByIds) // 批量删除CompetitionSche
		competitionScheRouter.PUT("updateCompetitionSche", competitionScheApi.UpdateCompetitionSche)    // 更新CompetitionSche
	}
	{
		competitionScheRouterWithoutRecord.GET("findCompetitionSche", competitionScheApi.FindCompetitionSche)        // 根据ID获取CompetitionSche
		competitionScheRouterWithoutRecord.GET("getCompetitionScheList", competitionScheApi.GetCompetitionScheList)  // 获取CompetitionSche列表
	}
}
