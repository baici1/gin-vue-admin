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

type StudentInfoApi struct {
}

var studentInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.StudentInfoService

// CreateStudentInfo 创建StudentInfo
// @Tags StudentInfo
// @Summary 创建StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.StudentInfo true "创建StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /studentInfo/createStudentInfo [post]
func (studentInfoApi *StudentInfoApi) CreateStudentInfo(c *gin.Context) {
	var studentInfo autocode.StudentInfo
	_ = c.ShouldBindJSON(&studentInfo)
	if err := studentInfoService.CreateStudentInfo(studentInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStudentInfo 删除StudentInfo
// @Tags StudentInfo
// @Summary 删除StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.StudentInfo true "删除StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /studentInfo/deleteStudentInfo [delete]
func (studentInfoApi *StudentInfoApi) DeleteStudentInfo(c *gin.Context) {
	var studentInfo autocode.StudentInfo
	_ = c.ShouldBindJSON(&studentInfo)
	if err := studentInfoService.DeleteStudentInfo(studentInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStudentInfoByIds 批量删除StudentInfo
// @Tags StudentInfo
// @Summary 批量删除StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /studentInfo/deleteStudentInfoByIds [delete]
func (studentInfoApi *StudentInfoApi) DeleteStudentInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := studentInfoService.DeleteStudentInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStudentInfo 更新StudentInfo
// @Tags StudentInfo
// @Summary 更新StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.StudentInfo true "更新StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /studentInfo/updateStudentInfo [put]
func (studentInfoApi *StudentInfoApi) UpdateStudentInfo(c *gin.Context) {
	var studentInfo autocode.StudentInfo
	_ = c.ShouldBindJSON(&studentInfo)
	if err := studentInfoService.UpdateStudentInfo(studentInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStudentInfo 用id查询StudentInfo
// @Tags StudentInfo
// @Summary 用id查询StudentInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.StudentInfo true "用id查询StudentInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /studentInfo/findStudentInfo [get]
func (studentInfoApi *StudentInfoApi) FindStudentInfo(c *gin.Context) {
	var studentInfo autocode.StudentInfo
	_ = c.ShouldBindQuery(&studentInfo)
	if err, restudentInfo := studentInfoService.GetStudentInfo(studentInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restudentInfo": restudentInfo}, c)
	}
}

// GetStudentInfoList 分页获取StudentInfo列表
// @Tags StudentInfo
// @Summary 分页获取StudentInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.StudentInfoSearch true "分页获取StudentInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /studentInfo/getStudentInfoList [get]
func (studentInfoApi *StudentInfoApi) GetStudentInfoList(c *gin.Context) {
	var pageInfo autocodeReq.StudentInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := studentInfoService.GetStudentInfoInfoList(pageInfo); err != nil {
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
