package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeacherInfoRouter struct {
}

// InitTeacherInfoRouter 初始化 TeacherInfo 路由信息
func (s *TeacherInfoRouter) InitTeacherInfoRouter(Router *gin.RouterGroup) {
	teacherInfoRouter := Router.Group("teacherInfo").Use(middleware.OperationRecord())
	teacherInfoRouterWithoutRecord := Router.Group("teacherInfo")
	var teacherInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.TeacherInfoApi
	{
		teacherInfoRouter.POST("createTeacherInfo", teacherInfoApi.CreateTeacherInfo)   // 新建TeacherInfo
		teacherInfoRouter.DELETE("deleteTeacherInfo", teacherInfoApi.DeleteTeacherInfo) // 删除TeacherInfo
		teacherInfoRouter.DELETE("deleteTeacherInfoByIds", teacherInfoApi.DeleteTeacherInfoByIds) // 批量删除TeacherInfo
		teacherInfoRouter.PUT("updateTeacherInfo", teacherInfoApi.UpdateTeacherInfo)    // 更新TeacherInfo
	}
	{
		teacherInfoRouterWithoutRecord.GET("findTeacherInfo", teacherInfoApi.FindTeacherInfo)        // 根据ID获取TeacherInfo
		teacherInfoRouterWithoutRecord.GET("getTeacherInfoList", teacherInfoApi.GetTeacherInfoList)  // 获取TeacherInfo列表
	}
}
