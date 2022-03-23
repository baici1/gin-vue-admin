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

type TeamMemberApi struct {
}

var teamMemberService = service.ServiceGroupApp.AutoCodeServiceGroup.TeamMemberService


// CreateTeamMember 创建TeamMember
// @Tags TeamMember
// @Summary 创建TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeamMember true "创建TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamMember/createTeamMember [post]
func (teamMemberApi *TeamMemberApi) CreateTeamMember(c *gin.Context) {
	var teamMember autocode.TeamMember
	_ = c.ShouldBindJSON(&teamMember)
	if err := teamMemberService.CreateTeamMember(teamMember); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeamMember 删除TeamMember
// @Tags TeamMember
// @Summary 删除TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeamMember true "删除TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamMember/deleteTeamMember [delete]
func (teamMemberApi *TeamMemberApi) DeleteTeamMember(c *gin.Context) {
	var teamMember autocode.TeamMember
	_ = c.ShouldBindJSON(&teamMember)
	if err := teamMemberService.DeleteTeamMember(teamMember); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeamMemberByIds 批量删除TeamMember
// @Tags TeamMember
// @Summary 批量删除TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teamMember/deleteTeamMemberByIds [delete]
func (teamMemberApi *TeamMemberApi) DeleteTeamMemberByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := teamMemberService.DeleteTeamMemberByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeamMember 更新TeamMember
// @Tags TeamMember
// @Summary 更新TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeamMember true "更新TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamMember/updateTeamMember [put]
func (teamMemberApi *TeamMemberApi) UpdateTeamMember(c *gin.Context) {
	var teamMember autocode.TeamMember
	_ = c.ShouldBindJSON(&teamMember)
	if err := teamMemberService.UpdateTeamMember(teamMember); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeamMember 用id查询TeamMember
// @Tags TeamMember
// @Summary 用id查询TeamMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TeamMember true "用id查询TeamMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamMember/findTeamMember [get]
func (teamMemberApi *TeamMemberApi) FindTeamMember(c *gin.Context) {
	var teamMember autocode.TeamMember
	_ = c.ShouldBindQuery(&teamMember)
	if err, reteamMember := teamMemberService.GetTeamMember(teamMember.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteamMember": reteamMember}, c)
	}
}

// GetTeamMemberList 分页获取TeamMember列表
// @Tags TeamMember
// @Summary 分页获取TeamMember列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TeamMemberSearch true "分页获取TeamMember列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamMember/getTeamMemberList [get]
func (teamMemberApi *TeamMemberApi) GetTeamMemberList(c *gin.Context) {
	var pageInfo autocodeReq.TeamMemberSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := teamMemberService.GetTeamMemberInfoList(pageInfo); err != nil {
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
