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

type AchievementApi struct {
}

var achievementService = service.ServiceGroupApp.AutoCodeServiceGroup.AchievementService


// CreateAchievement 创建Achievement
// @Tags Achievement
// @Summary 创建Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Achievement true "创建Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /achievement/createAchievement [post]
func (achievementApi *AchievementApi) CreateAchievement(c *gin.Context) {
	var achievement autocode.Achievement
	_ = c.ShouldBindJSON(&achievement)
	if err := achievementService.CreateAchievement(achievement); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAchievement 删除Achievement
// @Tags Achievement
// @Summary 删除Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Achievement true "删除Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /achievement/deleteAchievement [delete]
func (achievementApi *AchievementApi) DeleteAchievement(c *gin.Context) {
	var achievement autocode.Achievement
	_ = c.ShouldBindJSON(&achievement)
	if err := achievementService.DeleteAchievement(achievement); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAchievementByIds 批量删除Achievement
// @Tags Achievement
// @Summary 批量删除Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /achievement/deleteAchievementByIds [delete]
func (achievementApi *AchievementApi) DeleteAchievementByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := achievementService.DeleteAchievementByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAchievement 更新Achievement
// @Tags Achievement
// @Summary 更新Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Achievement true "更新Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /achievement/updateAchievement [put]
func (achievementApi *AchievementApi) UpdateAchievement(c *gin.Context) {
	var achievement autocode.Achievement
	_ = c.ShouldBindJSON(&achievement)
	if err := achievementService.UpdateAchievement(achievement); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAchievement 用id查询Achievement
// @Tags Achievement
// @Summary 用id查询Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Achievement true "用id查询Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /achievement/findAchievement [get]
func (achievementApi *AchievementApi) FindAchievement(c *gin.Context) {
	var achievement autocode.Achievement
	_ = c.ShouldBindQuery(&achievement)
	if err, reachievement := achievementService.GetAchievement(achievement.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reachievement": reachievement}, c)
	}
}

// GetAchievementList 分页获取Achievement列表
// @Tags Achievement
// @Summary 分页获取Achievement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.AchievementSearch true "分页获取Achievement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /achievement/getAchievementList [get]
func (achievementApi *AchievementApi) GetAchievementList(c *gin.Context) {
	var pageInfo autocodeReq.AchievementSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := achievementService.GetAchievementInfoList(pageInfo); err != nil {
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
