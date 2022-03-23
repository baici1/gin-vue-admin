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

type CompetitionInfoApi struct {
}

var competitionInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.CompetitionInfoService


// CreateCompetitionInfo 创建CompetitionInfo
// @Tags CompetitionInfo
// @Summary 创建CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CompetitionInfo true "创建CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /competitionInfo/createCompetitionInfo [post]
func (competitionInfoApi *CompetitionInfoApi) CreateCompetitionInfo(c *gin.Context) {
	var competitionInfo autocode.CompetitionInfo
	_ = c.ShouldBindJSON(&competitionInfo)
	if err := competitionInfoService.CreateCompetitionInfo(competitionInfo); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCompetitionInfo 删除CompetitionInfo
// @Tags CompetitionInfo
// @Summary 删除CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CompetitionInfo true "删除CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /competitionInfo/deleteCompetitionInfo [delete]
func (competitionInfoApi *CompetitionInfoApi) DeleteCompetitionInfo(c *gin.Context) {
	var competitionInfo autocode.CompetitionInfo
	_ = c.ShouldBindJSON(&competitionInfo)
	if err := competitionInfoService.DeleteCompetitionInfo(competitionInfo); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCompetitionInfoByIds 批量删除CompetitionInfo
// @Tags CompetitionInfo
// @Summary 批量删除CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /competitionInfo/deleteCompetitionInfoByIds [delete]
func (competitionInfoApi *CompetitionInfoApi) DeleteCompetitionInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := competitionInfoService.DeleteCompetitionInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCompetitionInfo 更新CompetitionInfo
// @Tags CompetitionInfo
// @Summary 更新CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CompetitionInfo true "更新CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /competitionInfo/updateCompetitionInfo [put]
func (competitionInfoApi *CompetitionInfoApi) UpdateCompetitionInfo(c *gin.Context) {
	var competitionInfo autocode.CompetitionInfo
	_ = c.ShouldBindJSON(&competitionInfo)
	if err := competitionInfoService.UpdateCompetitionInfo(competitionInfo); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCompetitionInfo 用id查询CompetitionInfo
// @Tags CompetitionInfo
// @Summary 用id查询CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CompetitionInfo true "用id查询CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /competitionInfo/findCompetitionInfo [get]
func (competitionInfoApi *CompetitionInfoApi) FindCompetitionInfo(c *gin.Context) {
	var competitionInfo autocode.CompetitionInfo
	_ = c.ShouldBindQuery(&competitionInfo)
	if err, recompetitionInfo := competitionInfoService.GetCompetitionInfo(competitionInfo.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recompetitionInfo": recompetitionInfo}, c)
	}
}

// GetCompetitionInfoList 分页获取CompetitionInfo列表
// @Tags CompetitionInfo
// @Summary 分页获取CompetitionInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CompetitionInfoSearch true "分页获取CompetitionInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /competitionInfo/getCompetitionInfoList [get]
func (competitionInfoApi *CompetitionInfoApi) GetCompetitionInfoList(c *gin.Context) {
	var pageInfo autocodeReq.CompetitionInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := competitionInfoService.GetCompetitionInfoInfoList(pageInfo); err != nil {
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
