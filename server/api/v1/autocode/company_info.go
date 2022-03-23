package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CompanyInfoApi struct {
}

var companyInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.CompanyInfoService


// CreateCompanyInfo 创建CompanyInfo
// @Tags CompanyInfo
// @Summary 创建CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CompanyInfo true "创建CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyInfo/createCompanyInfo [post]
func (companyInfoApi *CompanyInfoApi) CreateCompanyInfo(c *gin.Context) {
	var companyInfo autocode.CompanyInfo
	_ = c.ShouldBindJSON(&companyInfo)
	if err := companyInfoService.CreateCompanyInfo(companyInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCompanyInfo 删除CompanyInfo
// @Tags CompanyInfo
// @Summary 删除CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CompanyInfo true "删除CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /companyInfo/deleteCompanyInfo [delete]
func (companyInfoApi *CompanyInfoApi) DeleteCompanyInfo(c *gin.Context) {
	var companyInfo autocode.CompanyInfo
	_ = c.ShouldBindJSON(&companyInfo)
	if err := companyInfoService.DeleteCompanyInfo(companyInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCompanyInfoByIds 批量删除CompanyInfo
// @Tags CompanyInfo
// @Summary 批量删除CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /companyInfo/deleteCompanyInfoByIds [delete]
func (companyInfoApi *CompanyInfoApi) DeleteCompanyInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := companyInfoService.DeleteCompanyInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCompanyInfo 更新CompanyInfo
// @Tags CompanyInfo
// @Summary 更新CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CompanyInfo true "更新CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /companyInfo/updateCompanyInfo [put]
func (companyInfoApi *CompanyInfoApi) UpdateCompanyInfo(c *gin.Context) {
	var companyInfo autocode.CompanyInfo
	_ = c.ShouldBindJSON(&companyInfo)
	if err := companyInfoService.UpdateCompanyInfo(companyInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCompanyInfo 用id查询CompanyInfo
// @Tags CompanyInfo
// @Summary 用id查询CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CompanyInfo true "用id查询CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /companyInfo/findCompanyInfo [get]
func (companyInfoApi *CompanyInfoApi) FindCompanyInfo(c *gin.Context) {
	var companyInfo autocode.CompanyInfo
	_ = c.ShouldBindQuery(&companyInfo)
	if err, recompanyInfo := companyInfoService.GetCompanyInfo(companyInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recompanyInfo": recompanyInfo}, c)
	}
}

// GetCompanyInfoList 分页获取CompanyInfo列表
// @Tags CompanyInfo
// @Summary 分页获取CompanyInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CompanyInfoSearch true "分页获取CompanyInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyInfo/getCompanyInfoList [get]
func (companyInfoApi *CompanyInfoApi) GetCompanyInfoList(c *gin.Context) {
	var pageInfo autocodeReq.CompanyInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := companyInfoService.GetCompanyInfoInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
