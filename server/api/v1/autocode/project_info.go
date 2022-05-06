package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProjectInfoApi struct {
}

var projectInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.ProjectInfoService

// CreateProjectInfo 创建ProjectInfo
// @Tags ProjectInfo
// @Summary 创建ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ProjectInfo true "创建ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /projectInfo/createProjectInfo [post]
func (projectInfoApi *ProjectInfoApi) CreateProjectInfo(c *gin.Context) {
	var projectInfo autocode.ProjectInfo
	_ = c.ShouldBindJSON(&projectInfo)
	fmt.Println(projectInfo)
	if err, data := projectInfoService.CreateProjectInfo(projectInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(data, "创建成功", c)
	}
}

// DeleteProjectInfo 删除ProjectInfo
// @Tags ProjectInfo
// @Summary 删除ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ProjectInfo true "删除ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /projectInfo/deleteProjectInfo [delete]
func (projectInfoApi *ProjectInfoApi) DeleteProjectInfo(c *gin.Context) {
	var projectInfo autocode.ProjectInfo
	_ = c.ShouldBindJSON(&projectInfo)
	if err := projectInfoService.DeleteProjectInfo(projectInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProjectInfoByIds 批量删除ProjectInfo
// @Tags ProjectInfo
// @Summary 批量删除ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /projectInfo/deleteProjectInfoByIds [delete]
func (projectInfoApi *ProjectInfoApi) DeleteProjectInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := projectInfoService.DeleteProjectInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProjectInfo 更新ProjectInfo
// @Tags ProjectInfo
// @Summary 更新ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ProjectInfo true "更新ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /projectInfo/updateProjectInfo [put]
func (projectInfoApi *ProjectInfoApi) UpdateProjectInfo(c *gin.Context) {
	var projectInfo autocode.ProjectInfo
	_ = c.ShouldBindJSON(&projectInfo)
	if err := projectInfoService.UpdateProjectInfo(projectInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProjectInfo 用id查询ProjectInfo
// @Tags ProjectInfo
// @Summary 用id查询ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ProjectInfo true "用id查询ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /projectInfo/findProjectInfo [get]
func (projectInfoApi *ProjectInfoApi) FindProjectInfo(c *gin.Context) {
	var projectInfo autocode.ProjectInfo
	_ = c.ShouldBindQuery(&projectInfo)
	if err, reprojectInfo := projectInfoService.GetProjectInfo(projectInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reprojectInfo": reprojectInfo}, c)
	}
}

// GetProjectInfoList 分页获取ProjectInfo列表
// @Tags ProjectInfo
// @Summary 分页获取ProjectInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ProjectInfoSearch true "分页获取ProjectInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /projectInfo/getProjectInfoList [get]
func (projectInfoApi *ProjectInfoApi) GetProjectInfoList(c *gin.Context) {
	var pageInfo autocodeReq.ProjectInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := projectInfoService.GetProjectInfoInfoList(pageInfo); err != nil {
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
