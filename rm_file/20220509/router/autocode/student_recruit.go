package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StudentRecruitRouter struct {
}

// InitStudentRecruitRouter 初始化 StudentRecruit 路由信息
func (s *StudentRecruitRouter) InitStudentRecruitRouter(Router *gin.RouterGroup) {
	studentRecruitRouter := Router.Group("studentRecruit").Use(middleware.OperationRecord())
	studentRecruitRouterWithoutRecord := Router.Group("studentRecruit")
	var studentRecruitApi = v1.ApiGroupApp.AutoCodeApiGroup.StudentRecruitApi
	{
		studentRecruitRouter.POST("createStudentRecruit", studentRecruitApi.CreateStudentRecruit)   // 新建StudentRecruit
		studentRecruitRouter.DELETE("deleteStudentRecruit", studentRecruitApi.DeleteStudentRecruit) // 删除StudentRecruit
		studentRecruitRouter.DELETE("deleteStudentRecruitByIds", studentRecruitApi.DeleteStudentRecruitByIds) // 批量删除StudentRecruit
		studentRecruitRouter.PUT("updateStudentRecruit", studentRecruitApi.UpdateStudentRecruit)    // 更新StudentRecruit
	}
	{
		studentRecruitRouterWithoutRecord.GET("findStudentRecruit", studentRecruitApi.FindStudentRecruit)        // 根据ID获取StudentRecruit
		studentRecruitRouterWithoutRecord.GET("getStudentRecruitList", studentRecruitApi.GetStudentRecruitList)  // 获取StudentRecruit列表
	}
}
