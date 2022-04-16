package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AchievementRouter struct {
}

// InitAchievementRouter 初始化 Achievement 路由信息
func (s *AchievementRouter) InitAchievementRouter(Router *gin.RouterGroup) {
	achievementRouter := Router.Group("achievement").Use(middleware.OperationRecord())
	achievementRouterWithoutRecord := Router.Group("achievement")
	var achievementApi = v1.ApiGroupApp.AutoCodeApiGroup.AchievementApi
	{
		achievementRouter.POST("createAchievement", achievementApi.CreateAchievement)   // 新建Achievement
		achievementRouter.DELETE("deleteAchievement", achievementApi.DeleteAchievement) // 删除Achievement
		achievementRouter.DELETE("deleteAchievementByIds", achievementApi.DeleteAchievementByIds) // 批量删除Achievement
		achievementRouter.PUT("updateAchievement", achievementApi.UpdateAchievement)    // 更新Achievement
	}
	{
		achievementRouterWithoutRecord.GET("findAchievement", achievementApi.FindAchievement)        // 根据ID获取Achievement
		achievementRouterWithoutRecord.GET("getAchievementList", achievementApi.GetAchievementList)  // 获取Achievement列表
	}
}
