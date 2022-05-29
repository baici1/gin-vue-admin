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

type StudentRecruitApi struct {
}

var studentRecruitService = service.ServiceGroupApp.AutoCodeServiceGroup.StudentRecruitService

// CreateStudentRecruit 创建StudentRecruit
// @Tags StudentRecruit
// @Summary 创建StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.StudentRecruit true "创建StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /studentRecruit/createStudentRecruit [post]
func (studentRecruitApi *StudentRecruitApi) CreateStudentRecruit(c *gin.Context) {
	var studentRecruit autocode.StudentRecruit
	_ = c.ShouldBindJSON(&studentRecruit)
	if err := studentRecruitService.CreateStudentRecruit(studentRecruit); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStudentRecruit 删除StudentRecruit
// @Tags StudentRecruit
// @Summary 删除StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.StudentRecruit true "删除StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /studentRecruit/deleteStudentRecruit [delete]
func (studentRecruitApi *StudentRecruitApi) DeleteStudentRecruit(c *gin.Context) {
	var studentRecruit autocode.StudentRecruit
	_ = c.ShouldBindJSON(&studentRecruit)
	if err := studentRecruitService.DeleteStudentRecruit(studentRecruit); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStudentRecruitByIds 批量删除StudentRecruit
// @Tags StudentRecruit
// @Summary 批量删除StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /studentRecruit/deleteStudentRecruitByIds [delete]
func (studentRecruitApi *StudentRecruitApi) DeleteStudentRecruitByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := studentRecruitService.DeleteStudentRecruitByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStudentRecruit 更新StudentRecruit
// @Tags StudentRecruit
// @Summary 更新StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.StudentRecruit true "更新StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /studentRecruit/updateStudentRecruit [put]
func (studentRecruitApi *StudentRecruitApi) UpdateStudentRecruit(c *gin.Context) {
	var studentRecruit autocode.StudentRecruit
	_ = c.ShouldBindJSON(&studentRecruit)
	if err := studentRecruitService.UpdateStudentRecruit(studentRecruit); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStudentRecruit 用id查询StudentRecruit
// @Tags StudentRecruit
// @Summary 用id查询StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.StudentRecruit true "用id查询StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /studentRecruit/findStudentRecruit [get]
func (studentRecruitApi *StudentRecruitApi) FindStudentRecruit(c *gin.Context) {
	var studentRecruit autocode.StudentRecruit
	_ = c.ShouldBindQuery(&studentRecruit)
	if err, restudentRecruit := studentRecruitService.GetStudentRecruit(studentRecruit.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restudentRecruit": restudentRecruit}, c)
	}
}

// GetStudentRecruitList 分页获取StudentRecruit列表
// @Tags StudentRecruit
// @Summary 分页获取StudentRecruit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.StudentRecruitSearch true "分页获取StudentRecruit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /studentRecruit/getStudentRecruitList [get]
func (studentRecruitApi *StudentRecruitApi) GetStudentRecruitList(c *gin.Context) {
	var pageInfo autocodeReq.StudentRecruitSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := studentRecruitService.GetStudentRecruitInfoList(pageInfo); err != nil {
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

func (studentRecruitApi *StudentRecruitApi) ProduceStudentRecruitInfomation(c *gin.Context) {
	var studentRecruit autocodeReq.StudentRecruitToRabbitmq
	_ = c.ShouldBindJSON(&studentRecruit)
	if err := studentRecruitService.ProduceStudentRecruitInfomationByHello(studentRecruit.Producer, studentRecruit.Comsumer); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithMessage("发送成功", c)
	}
}
func (studentRecruitApi *StudentRecruitApi) ProduceRecruitInfomation(c *gin.Context) {
	var Recruit autocodeReq.RecruitToRabbitmqInfo
	_ = c.ShouldBindJSON(&Recruit)
	if err := studentRecruitService.ProduceRecruitInfomationByRouting(Recruit.Producer, Recruit.Comsumer); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithMessage("发送成功", c)
	}
}
