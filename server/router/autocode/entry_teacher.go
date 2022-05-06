package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EntryTeacherRouter struct {
}

// InitEntryTeacherRouter 初始化 EntryTeacher 路由信息
func (s *EntryTeacherRouter) InitEntryTeacherRouter(Router *gin.RouterGroup) {
	entryTeacherRouter := Router.Group("entryTeacher").Use(middleware.OperationRecord())
	entryTeacherRouterWithoutRecord := Router.Group("entryTeacher")
	var entryTeacherApi = v1.ApiGroupApp.AutoCodeApiGroup.EntryTeacherApi
	{
		entryTeacherRouter.POST("createEntryTeacher", entryTeacherApi.CreateEntryTeacher)   // 新建EntryTeacher
		entryTeacherRouter.DELETE("deleteEntryTeacher", entryTeacherApi.DeleteEntryTeacher) // 删除EntryTeacher
		entryTeacherRouter.DELETE("deleteEntryTeacherByIds", entryTeacherApi.DeleteEntryTeacherByIds) // 批量删除EntryTeacher
		entryTeacherRouter.PUT("updateEntryTeacher", entryTeacherApi.UpdateEntryTeacher)    // 更新EntryTeacher
	}
	{
		entryTeacherRouterWithoutRecord.GET("findEntryTeacher", entryTeacherApi.FindEntryTeacher)        // 根据ID获取EntryTeacher
		entryTeacherRouterWithoutRecord.GET("getEntryTeacherList", entryTeacherApi.GetEntryTeacherList)  // 获取EntryTeacher列表
	}
}
