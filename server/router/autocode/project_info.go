package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectInfoRouter struct {
}

// InitProjectInfoRouter 初始化 ProjectInfo 路由信息
func (s *ProjectInfoRouter) InitProjectInfoRouter(Router *gin.RouterGroup) {
	projectInfoRouter := Router.Group("projectInfo").Use(middleware.OperationRecord())
	projectInfoRouterWithoutRecord := Router.Group("projectInfo")
	var projectInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.ProjectInfoApi
	{
		projectInfoRouter.POST("createProjectInfo", projectInfoApi.CreateProjectInfo)   // 新建ProjectInfo
		projectInfoRouter.DELETE("deleteProjectInfo", projectInfoApi.DeleteProjectInfo) // 删除ProjectInfo
		projectInfoRouter.DELETE("deleteProjectInfoByIds", projectInfoApi.DeleteProjectInfoByIds) // 批量删除ProjectInfo
		projectInfoRouter.PUT("updateProjectInfo", projectInfoApi.UpdateProjectInfo)    // 更新ProjectInfo
	}
	{
		projectInfoRouterWithoutRecord.GET("findProjectInfo", projectInfoApi.FindProjectInfo)        // 根据ID获取ProjectInfo
		projectInfoRouterWithoutRecord.GET("getProjectInfoList", projectInfoApi.GetProjectInfoList)  // 获取ProjectInfo列表
	}
}
