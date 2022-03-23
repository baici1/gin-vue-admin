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

type CompetitionScheApi struct {
}

var competitionScheService = service.ServiceGroupApp.AutoCodeServiceGroup.CompetitionScheService


// CreateCompetitionSche 创建CompetitionSche
// @Tags CompetitionSche
// @Summary 创建CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CompetitionSche true "创建CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /competitionSche/createCompetitionSche [post]
func (competitionScheApi *CompetitionScheApi) CreateCompetitionSche(c *gin.Context) {
	var competitionSche autocode.CompetitionSche
	_ = c.ShouldBindJSON(&competitionSche)
	if err := competitionScheService.CreateCompetitionSche(competitionSche); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCompetitionSche 删除CompetitionSche
// @Tags CompetitionSche
// @Summary 删除CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CompetitionSche true "删除CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /competitionSche/deleteCompetitionSche [delete]
func (competitionScheApi *CompetitionScheApi) DeleteCompetitionSche(c *gin.Context) {
	var competitionSche autocode.CompetitionSche
	_ = c.ShouldBindJSON(&competitionSche)
	if err := competitionScheService.DeleteCompetitionSche(competitionSche); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCompetitionScheByIds 批量删除CompetitionSche
// @Tags CompetitionSche
// @Summary 批量删除CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /competitionSche/deleteCompetitionScheByIds [delete]
func (competitionScheApi *CompetitionScheApi) DeleteCompetitionScheByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := competitionScheService.DeleteCompetitionScheByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCompetitionSche 更新CompetitionSche
// @Tags CompetitionSche
// @Summary 更新CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CompetitionSche true "更新CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /competitionSche/updateCompetitionSche [put]
func (competitionScheApi *CompetitionScheApi) UpdateCompetitionSche(c *gin.Context) {
	var competitionSche autocode.CompetitionSche
	_ = c.ShouldBindJSON(&competitionSche)
	if err := competitionScheService.UpdateCompetitionSche(competitionSche); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCompetitionSche 用id查询CompetitionSche
// @Tags CompetitionSche
// @Summary 用id查询CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CompetitionSche true "用id查询CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /competitionSche/findCompetitionSche [get]
func (competitionScheApi *CompetitionScheApi) FindCompetitionSche(c *gin.Context) {
	var competitionSche autocode.CompetitionSche
	_ = c.ShouldBindQuery(&competitionSche)
	if err, recompetitionSche := competitionScheService.GetCompetitionSche(competitionSche.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recompetitionSche": recompetitionSche}, c)
	}
}

// GetCompetitionScheList 分页获取CompetitionSche列表
// @Tags CompetitionSche
// @Summary 分页获取CompetitionSche列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CompetitionScheSearch true "分页获取CompetitionSche列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /competitionSche/getCompetitionScheList [get]
func (competitionScheApi *CompetitionScheApi) GetCompetitionScheList(c *gin.Context) {
	var pageInfo autocodeReq.CompetitionScheSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := competitionScheService.GetCompetitionScheInfoList(pageInfo); err != nil {
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
