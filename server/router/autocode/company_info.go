package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyInfoRouter struct {
}

// InitCompanyInfoRouter 初始化 CompanyInfo 路由信息
func (s *CompanyInfoRouter) InitCompanyInfoRouter(Router *gin.RouterGroup) {
	companyInfoRouter := Router.Group("companyInfo").Use(middleware.OperationRecord())
	companyInfoRouterWithoutRecord := Router.Group("companyInfo")
	var companyInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.CompanyInfoApi
	{
		companyInfoRouter.POST("createCompanyInfo", companyInfoApi.CreateCompanyInfo)   // 新建CompanyInfo
		companyInfoRouter.DELETE("deleteCompanyInfo", companyInfoApi.DeleteCompanyInfo) // 删除CompanyInfo
		companyInfoRouter.DELETE("deleteCompanyInfoByIds", companyInfoApi.DeleteCompanyInfoByIds) // 批量删除CompanyInfo
		companyInfoRouter.PUT("updateCompanyInfo", companyInfoApi.UpdateCompanyInfo)    // 更新CompanyInfo
	}
	{
		companyInfoRouterWithoutRecord.GET("findCompanyInfo", companyInfoApi.FindCompanyInfo)        // 根据ID获取CompanyInfo
		companyInfoRouterWithoutRecord.GET("getCompanyInfoList", companyInfoApi.GetCompanyInfoList)  // 获取CompanyInfo列表
	}
}
