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

type EntryMemberApi struct {
}

var entryMemberService = service.ServiceGroupApp.AutoCodeServiceGroup.EntryMemberService


// CreateEntryMember 创建EntryMember
// @Tags EntryMember
// @Summary 创建EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EntryMember true "创建EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryMember/createEntryMember [post]
func (entryMemberApi *EntryMemberApi) CreateEntryMember(c *gin.Context) {
	var entryMember autocode.EntryMember
	_ = c.ShouldBindJSON(&entryMember)
	if err := entryMemberService.CreateEntryMember(entryMember); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEntryMember 删除EntryMember
// @Tags EntryMember
// @Summary 删除EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EntryMember true "删除EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /entryMember/deleteEntryMember [delete]
func (entryMemberApi *EntryMemberApi) DeleteEntryMember(c *gin.Context) {
	var entryMember autocode.EntryMember
	_ = c.ShouldBindJSON(&entryMember)
	if err := entryMemberService.DeleteEntryMember(entryMember); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEntryMemberByIds 批量删除EntryMember
// @Tags EntryMember
// @Summary 批量删除EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /entryMember/deleteEntryMemberByIds [delete]
func (entryMemberApi *EntryMemberApi) DeleteEntryMemberByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := entryMemberService.DeleteEntryMemberByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEntryMember 更新EntryMember
// @Tags EntryMember
// @Summary 更新EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EntryMember true "更新EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /entryMember/updateEntryMember [put]
func (entryMemberApi *EntryMemberApi) UpdateEntryMember(c *gin.Context) {
	var entryMember autocode.EntryMember
	_ = c.ShouldBindJSON(&entryMember)
	if err := entryMemberService.UpdateEntryMember(entryMember); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEntryMember 用id查询EntryMember
// @Tags EntryMember
// @Summary 用id查询EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.EntryMember true "用id查询EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /entryMember/findEntryMember [get]
func (entryMemberApi *EntryMemberApi) FindEntryMember(c *gin.Context) {
	var entryMember autocode.EntryMember
	_ = c.ShouldBindQuery(&entryMember)
	if err, reentryMember := entryMemberService.GetEntryMember(entryMember.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reentryMember": reentryMember}, c)
	}
}

// GetEntryMemberList 分页获取EntryMember列表
// @Tags EntryMember
// @Summary 分页获取EntryMember列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.EntryMemberSearch true "分页获取EntryMember列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryMember/getEntryMemberList [get]
func (entryMemberApi *EntryMemberApi) GetEntryMemberList(c *gin.Context) {
	var pageInfo autocodeReq.EntryMemberSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := entryMemberService.GetEntryMemberInfoList(pageInfo); err != nil {
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
