package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProduceComsumeInfomationRouter struct {
}

// InitProduceComsumeInfomationRouter 初始化 ProduceComsumeInfomation 路由信息
func (s *ProduceComsumeInfomationRouter) InitProduceComsumeInfomationRouter(Router *gin.RouterGroup) {
	produceComsumeInfomationRouter := Router.Group("produceComsumeInfomation").Use(middleware.OperationRecord())
	produceComsumeInfomationRouterWithoutRecord := Router.Group("produceComsumeInfomation")
	var produceComsumeInfomationApi = v1.ApiGroupApp.AutoCodeApiGroup.ProduceComsumeInfomationApi
	{
		produceComsumeInfomationRouter.POST("createProduceComsumeInfomation", produceComsumeInfomationApi.CreateProduceComsumeInfomation)   // 新建ProduceComsumeInfomation
		produceComsumeInfomationRouter.DELETE("deleteProduceComsumeInfomation", produceComsumeInfomationApi.DeleteProduceComsumeInfomation) // 删除ProduceComsumeInfomation
		produceComsumeInfomationRouter.DELETE("deleteProduceComsumeInfomationByIds", produceComsumeInfomationApi.DeleteProduceComsumeInfomationByIds) // 批量删除ProduceComsumeInfomation
		produceComsumeInfomationRouter.PUT("updateProduceComsumeInfomation", produceComsumeInfomationApi.UpdateProduceComsumeInfomation)    // 更新ProduceComsumeInfomation
	}
	{
		produceComsumeInfomationRouterWithoutRecord.GET("findProduceComsumeInfomation", produceComsumeInfomationApi.FindProduceComsumeInfomation)        // 根据ID获取ProduceComsumeInfomation
		produceComsumeInfomationRouterWithoutRecord.GET("getProduceComsumeInfomationList", produceComsumeInfomationApi.GetProduceComsumeInfomationList)  // 获取ProduceComsumeInfomation列表
	}
}
