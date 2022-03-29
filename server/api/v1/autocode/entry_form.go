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

type EntryFormApi struct {
}

var entryFormService = service.ServiceGroupApp.AutoCodeServiceGroup.EntryFormService

// CreateEntryForm 创建EntryForm
// @Tags EntryForm
// @Summary 创建EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EntryForm true "创建EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryForm/createEntryForm [post]
func (entryFormApi *EntryFormApi) CreateEntryForm(c *gin.Context) {
	var entryForm autocode.EntryForm
	_ = c.ShouldBindJSON(&entryForm)
	if err := entryFormService.CreateEntryForm(entryForm); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEntryForm 删除EntryForm
// @Tags EntryForm
// @Summary 删除EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EntryForm true "删除EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /entryForm/deleteEntryForm [delete]
func (entryFormApi *EntryFormApi) DeleteEntryForm(c *gin.Context) {
	var entryForm autocode.EntryForm
	_ = c.ShouldBindJSON(&entryForm)
	if err := entryFormService.DeleteEntryForm(entryForm); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEntryFormByIds 批量删除EntryForm
// @Tags EntryForm
// @Summary 批量删除EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /entryForm/deleteEntryFormByIds [delete]
func (entryFormApi *EntryFormApi) DeleteEntryFormByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := entryFormService.DeleteEntryFormByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEntryForm 更新EntryForm
// @Tags EntryForm
// @Summary 更新EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.EntryForm true "更新EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /entryForm/updateEntryForm [put]
func (entryFormApi *EntryFormApi) UpdateEntryForm(c *gin.Context) {
	var entryForm autocode.EntryForm
	_ = c.ShouldBindJSON(&entryForm)
	if err := entryFormService.UpdateEntryForm(entryForm); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEntryForm 用id查询EntryForm
// @Tags EntryForm
// @Summary 用id查询EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.EntryForm true "用id查询EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /entryForm/findEntryForm [get]
func (entryFormApi *EntryFormApi) FindEntryForm(c *gin.Context) {
	var entryForm autocode.EntryForm
	_ = c.ShouldBindQuery(&entryForm)
	if err, reentryForm := entryFormService.GetEntryForm(entryForm.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reentryForm": reentryForm}, c)
	}
}

// GetEntryFormList 分页获取EntryForm列表
// @Tags EntryForm
// @Summary 分页获取EntryForm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.EntryFormSearch true "分页获取EntryForm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryForm/getEntryFormList [get]
func (entryFormApi *EntryFormApi) GetEntryFormList(c *gin.Context) {
	var pageInfo autocodeReq.EntryFormSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := entryFormService.GetEntryFormInfoList(pageInfo); err != nil {
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

func (entryFormApi *EntryFormApi) GetAllEntryFormDetailInfo(c *gin.Context) {
	var param autocodeReq.EntryFormAllSearch
	err := c.ShouldBind(&param)
	if err != nil {
		response.ValidatorError(err, c)
		return
	}
	if err, list := entryFormService.GetAllEntryFormDetailInfo(param.UId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

func (entryFormApi *EntryFormApi) GetEntryFormDetailInfo(c *gin.Context) {
	param := struct {
		Fid int `form:"id" json:"fid"`
	}{}
	err := c.ShouldBind(&param)
	if err != nil {
		response.ValidatorError(err, c)
		return
	}
	if err, list := entryFormService.GetEntryFormDetailInfo(param.Fid); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(&list, c)
	}
}

func (entryFormApi *EntryFormApi) CreateEntryFormByUser(c *gin.Context) {
	var param autocodeReq.CreateEntryForm
	if err := c.ShouldBind(&param); err != nil {
		response.ValidatorError(err, c)
		return
	}
	if err := entryFormService.CreateEntryFormByUser(param); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (entryFormApi *EntryFormApi) UpdateEntryFormByUser(c *gin.Context) {
	var param autocodeReq.UpdateEntryForm
	if err := c.ShouldBind(&param); err != nil {
		response.ValidatorError(err, c)
		return
	}
	if err := entryFormService.UpdateEntryFormByUser(param); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}
