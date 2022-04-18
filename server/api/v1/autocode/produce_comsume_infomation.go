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

type ProduceComsumeInfomationApi struct {
}

var produceComsumeInfomationService = service.ServiceGroupApp.AutoCodeServiceGroup.ProduceComsumeInfomationService


// CreateProduceComsumeInfomation 创建ProduceComsumeInfomation
// @Tags ProduceComsumeInfomation
// @Summary 创建ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ProduceComsumeInfomation true "创建ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /produceComsumeInfomation/createProduceComsumeInfomation [post]
func (produceComsumeInfomationApi *ProduceComsumeInfomationApi) CreateProduceComsumeInfomation(c *gin.Context) {
	var produceComsumeInfomation autocode.ProduceComsumeInfomation
	_ = c.ShouldBindJSON(&produceComsumeInfomation)
	if err := produceComsumeInfomationService.CreateProduceComsumeInfomation(produceComsumeInfomation); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProduceComsumeInfomation 删除ProduceComsumeInfomation
// @Tags ProduceComsumeInfomation
// @Summary 删除ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ProduceComsumeInfomation true "删除ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /produceComsumeInfomation/deleteProduceComsumeInfomation [delete]
func (produceComsumeInfomationApi *ProduceComsumeInfomationApi) DeleteProduceComsumeInfomation(c *gin.Context) {
	var produceComsumeInfomation autocode.ProduceComsumeInfomation
	_ = c.ShouldBindJSON(&produceComsumeInfomation)
	if err := produceComsumeInfomationService.DeleteProduceComsumeInfomation(produceComsumeInfomation); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProduceComsumeInfomationByIds 批量删除ProduceComsumeInfomation
// @Tags ProduceComsumeInfomation
// @Summary 批量删除ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /produceComsumeInfomation/deleteProduceComsumeInfomationByIds [delete]
func (produceComsumeInfomationApi *ProduceComsumeInfomationApi) DeleteProduceComsumeInfomationByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := produceComsumeInfomationService.DeleteProduceComsumeInfomationByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProduceComsumeInfomation 更新ProduceComsumeInfomation
// @Tags ProduceComsumeInfomation
// @Summary 更新ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ProduceComsumeInfomation true "更新ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /produceComsumeInfomation/updateProduceComsumeInfomation [put]
func (produceComsumeInfomationApi *ProduceComsumeInfomationApi) UpdateProduceComsumeInfomation(c *gin.Context) {
	var produceComsumeInfomation autocode.ProduceComsumeInfomation
	_ = c.ShouldBindJSON(&produceComsumeInfomation)
	if err := produceComsumeInfomationService.UpdateProduceComsumeInfomation(produceComsumeInfomation); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProduceComsumeInfomation 用id查询ProduceComsumeInfomation
// @Tags ProduceComsumeInfomation
// @Summary 用id查询ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ProduceComsumeInfomation true "用id查询ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /produceComsumeInfomation/findProduceComsumeInfomation [get]
func (produceComsumeInfomationApi *ProduceComsumeInfomationApi) FindProduceComsumeInfomation(c *gin.Context) {
	var produceComsumeInfomation autocode.ProduceComsumeInfomation
	_ = c.ShouldBindQuery(&produceComsumeInfomation)
	if err, reproduceComsumeInfomation := produceComsumeInfomationService.GetProduceComsumeInfomation(produceComsumeInfomation.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproduceComsumeInfomation": reproduceComsumeInfomation}, c)
	}
}

// GetProduceComsumeInfomationList 分页获取ProduceComsumeInfomation列表
// @Tags ProduceComsumeInfomation
// @Summary 分页获取ProduceComsumeInfomation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ProduceComsumeInfomationSearch true "分页获取ProduceComsumeInfomation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /produceComsumeInfomation/getProduceComsumeInfomationList [get]
func (produceComsumeInfomationApi *ProduceComsumeInfomationApi) GetProduceComsumeInfomationList(c *gin.Context) {
	var pageInfo autocodeReq.ProduceComsumeInfomationSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := produceComsumeInfomationService.GetProduceComsumeInfomationInfoList(pageInfo); err != nil {
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
