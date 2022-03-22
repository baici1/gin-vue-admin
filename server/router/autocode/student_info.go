package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StudentInfoRouter struct {
}

// InitStudentInfoRouter 初始化 StudentInfo 路由信息
func (s *StudentInfoRouter) InitStudentInfoRouter(Router *gin.RouterGroup) {
	studentInfoRouter := Router.Group("studentInfo").Use(middleware.OperationRecord())
	studentInfoRouterWithoutRecord := Router.Group("studentInfo")
	var studentInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.StudentInfoApi
	{
		studentInfoRouter.POST("createStudentInfo", studentInfoApi.CreateStudentInfo)   // 新建StudentInfo
		studentInfoRouter.DELETE("deleteStudentInfo", studentInfoApi.DeleteStudentInfo) // 删除StudentInfo
		studentInfoRouter.PUT("updateStudentInfo", studentInfoApi.UpdateStudentInfo)    // 更新StudentInfo
	}
	{
		studentInfoRouterWithoutRecord.GET("findStudentInfo", studentInfoApi.FindStudentInfo) // 根据UID获取StudentInfo
	}
}
