package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EntryFormRouter struct {
}

// InitEntryFormRouter 初始化 EntryForm 路由信息
func (s *EntryFormRouter) InitEntryFormRouter(Router *gin.RouterGroup) {
	entryFormRouter := Router.Group("entryForm").Use(middleware.OperationRecord())
	entryFormRouterWithoutRecord := Router.Group("entryForm")
	var entryFormApi = v1.ApiGroupApp.AutoCodeApiGroup.EntryFormApi
	{
		entryFormRouter.POST("createEntryForm", entryFormApi.CreateEntryForm)             // 新建EntryForm
		entryFormRouter.DELETE("deleteEntryForm", entryFormApi.DeleteEntryForm)           // 删除EntryForm
		entryFormRouter.DELETE("deleteEntryFormByIds", entryFormApi.DeleteEntryFormByIds) // 批量删除EntryForm
		entryFormRouter.PUT("updateEntryForm", entryFormApi.UpdateEntryForm)              // 更新EntryForm

	}
	{
		entryFormRouterWithoutRecord.GET("findEntryForm", entryFormApi.FindEntryForm)       // 根据ID获取EntryForm
		entryFormRouterWithoutRecord.GET("getEntryFormList", entryFormApi.GetEntryFormList) // 获取EntryForm列表
	}
}
