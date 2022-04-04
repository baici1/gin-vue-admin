package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TeamInfoApi struct {
}

var teamInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.TeamInfoService

// CreateTeamInfo 创建TeamInfo
// @Tags TeamInfo
// @Summary 创建TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeamInfo true "创建TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamInfo/createTeamInfo [post]
func (teamInfoApi *TeamInfoApi) CreateTeamInfo(c *gin.Context) {
	var teamInfo autocode.TeamInfo
	_ = c.ShouldBindJSON(&teamInfo)
	if err := teamInfoService.CreateTeamInfo(&teamInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(teamInfo.ID, c)
	}
}

// DeleteTeamInfo 删除TeamInfo
// @Tags TeamInfo
// @Summary 删除TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeamInfo true "删除TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamInfo/deleteTeamInfo [delete]
func (teamInfoApi *TeamInfoApi) DeleteTeamInfo(c *gin.Context) {
	var teamInfo autocode.TeamInfo
	_ = c.ShouldBindJSON(&teamInfo)
	if err := teamInfoService.DeleteTeamInfo(teamInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeamInfoByIds 批量删除TeamInfo
// @Tags TeamInfo
// @Summary 批量删除TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teamInfo/deleteTeamInfoByIds [delete]
func (teamInfoApi *TeamInfoApi) DeleteTeamInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := teamInfoService.DeleteTeamInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeamInfo 更新TeamInfo
// @Tags TeamInfo
// @Summary 更新TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeamInfo true "更新TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamInfo/updateTeamInfo [put]
func (teamInfoApi *TeamInfoApi) UpdateTeamInfo(c *gin.Context) {
	var teamInfo autocode.TeamInfo
	_ = c.ShouldBindJSON(&teamInfo)
	if err := teamInfoService.UpdateTeamInfo(teamInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeamInfo 用id查询TeamInfo
// @Tags TeamInfo
// @Summary 用id查询TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TeamInfo true "用id查询TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamInfo/findTeamInfo [get]
func (teamInfoApi *TeamInfoApi) FindTeamInfo(c *gin.Context) {
	var teamInfo autocode.TeamInfo
	_ = c.ShouldBindQuery(&teamInfo)
	if err, reteamInfo := teamInfoService.GetTeamInfo(teamInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteamInfo": reteamInfo}, c)
	}
}

// GetTeamInfoList 分页获取TeamInfo列表
// @Tags TeamInfo
// @Summary 分页获取TeamInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TeamInfoSearch true "分页获取TeamInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamInfo/getTeamInfoList [get]
func (teamInfoApi *TeamInfoApi) GetTeamInfoList(c *gin.Context) {
	var pageInfo autocodeReq.TeamInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := teamInfoService.GetTeamInfoInfoList(pageInfo); err != nil {
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
