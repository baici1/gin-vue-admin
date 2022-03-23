package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EntryMemberRouter struct {
}

// InitEntryMemberRouter 初始化 EntryMember 路由信息
func (s *EntryMemberRouter) InitEntryMemberRouter(Router *gin.RouterGroup) {
	entryMemberRouter := Router.Group("entryMember").Use(middleware.OperationRecord())
	entryMemberRouterWithoutRecord := Router.Group("entryMember")
	var entryMemberApi = v1.ApiGroupApp.AutoCodeApiGroup.EntryMemberApi
	{
		entryMemberRouter.POST("createEntryMember", entryMemberApi.CreateEntryMember)   // 新建EntryMember
		entryMemberRouter.DELETE("deleteEntryMember", entryMemberApi.DeleteEntryMember) // 删除EntryMember
		entryMemberRouter.DELETE("deleteEntryMemberByIds", entryMemberApi.DeleteEntryMemberByIds) // 批量删除EntryMember
		entryMemberRouter.PUT("updateEntryMember", entryMemberApi.UpdateEntryMember)    // 更新EntryMember
	}
	{
		entryMemberRouterWithoutRecord.GET("findEntryMember", entryMemberApi.FindEntryMember)        // 根据ID获取EntryMember
		entryMemberRouterWithoutRecord.GET("getEntryMemberList", entryMemberApi.GetEntryMemberList)  // 获取EntryMember列表
	}
}
