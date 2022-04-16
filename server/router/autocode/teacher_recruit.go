package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TeacherRecruitRouter struct {
}

// InitTeacherRecruitRouter 初始化 TeacherRecruit 路由信息
func (s *TeacherRecruitRouter) InitTeacherRecruitRouter(Router *gin.RouterGroup) {
	teacherRecruitRouter := Router.Group("teacherRecruit").Use(middleware.OperationRecord())
	teacherRecruitRouterWithoutRecord := Router.Group("teacherRecruit")
	var teacherRecruitApi = v1.ApiGroupApp.AutoCodeApiGroup.TeacherRecruitApi
	{
		teacherRecruitRouter.POST("createTeacherRecruit", teacherRecruitApi.CreateTeacherRecruit)   // 新建TeacherRecruit
		teacherRecruitRouter.DELETE("deleteTeacherRecruit", teacherRecruitApi.DeleteTeacherRecruit) // 删除TeacherRecruit
		teacherRecruitRouter.DELETE("deleteTeacherRecruitByIds", teacherRecruitApi.DeleteTeacherRecruitByIds) // 批量删除TeacherRecruit
		teacherRecruitRouter.PUT("updateTeacherRecruit", teacherRecruitApi.UpdateTeacherRecruit)    // 更新TeacherRecruit
	}
	{
		teacherRecruitRouterWithoutRecord.GET("findTeacherRecruit", teacherRecruitApi.FindTeacherRecruit)        // 根据ID获取TeacherRecruit
		teacherRecruitRouterWithoutRecord.GET("getTeacherRecruitList", teacherRecruitApi.GetTeacherRecruitList)  // 获取TeacherRecruit列表
	}
}
