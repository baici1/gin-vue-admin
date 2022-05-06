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

type EntryTeacherApi struct {
}

var entryTeacherService = service.ServiceGroupApp.AutoCodeServiceGroup.EntryTeacherService


// CreateEntryTeacher 创建EntryTeacher
// @Tags EntryTeacher
// @Summary 创建EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EntryTeacher true "创建EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryTeacher/createEntryTeacher [post]
func (entryTeacherApi *EntryTeacherApi) CreateEntryTeacher(c *gin.Context) {
	var entryTeacher autocode.EntryTeacher
	_ = c.ShouldBindJSON(&entryTeacher)
	if err := entryTeacherService.CreateEntryTeacher(entryTeacher); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEntryTeacher 删除EntryTeacher
// @Tags EntryTeacher
// @Summary 删除EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EntryTeacher true "删除EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /entryTeacher/deleteEntryTeacher [delete]
func (entryTeacherApi *EntryTeacherApi) DeleteEntryTeacher(c *gin.Context) {
	var entryTeacher autocode.EntryTeacher
	_ = c.ShouldBindJSON(&entryTeacher)
	if err := entryTeacherService.DeleteEntryTeacher(entryTeacher); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEntryTeacherByIds 批量删除EntryTeacher
// @Tags EntryTeacher
// @Summary 批量删除EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /entryTeacher/deleteEntryTeacherByIds [delete]
func (entryTeacherApi *EntryTeacherApi) DeleteEntryTeacherByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := entryTeacherService.DeleteEntryTeacherByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEntryTeacher 更新EntryTeacher
// @Tags EntryTeacher
// @Summary 更新EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EntryTeacher true "更新EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /entryTeacher/updateEntryTeacher [put]
func (entryTeacherApi *EntryTeacherApi) UpdateEntryTeacher(c *gin.Context) {
	var entryTeacher autocode.EntryTeacher
	_ = c.ShouldBindJSON(&entryTeacher)
	if err := entryTeacherService.UpdateEntryTeacher(entryTeacher); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEntryTeacher 用id查询EntryTeacher
// @Tags EntryTeacher
// @Summary 用id查询EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.EntryTeacher true "用id查询EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /entryTeacher/findEntryTeacher [get]
func (entryTeacherApi *EntryTeacherApi) FindEntryTeacher(c *gin.Context) {
	var entryTeacher autocode.EntryTeacher
	_ = c.ShouldBindQuery(&entryTeacher)
	if err, reentryTeacher := entryTeacherService.GetEntryTeacher(entryTeacher.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reentryTeacher": reentryTeacher}, c)
	}
}

// GetEntryTeacherList 分页获取EntryTeacher列表
// @Tags EntryTeacher
// @Summary 分页获取EntryTeacher列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.EntryTeacherSearch true "分页获取EntryTeacher列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryTeacher/getEntryTeacherList [get]
func (entryTeacherApi *EntryTeacherApi) GetEntryTeacherList(c *gin.Context) {
	var pageInfo autocodeReq.EntryTeacherSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := entryTeacherService.GetEntryTeacherInfoList(pageInfo); err != nil {
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
