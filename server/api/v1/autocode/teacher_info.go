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

type TeacherInfoApi struct {
}

var teacherInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.TeacherInfoService

// CreateTeacherInfo 创建TeacherInfo
// @Tags TeacherInfo
// @Summary 创建TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeacherInfo true "创建TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacherInfo/createTeacherInfo [post]
func (teacherInfoApi *TeacherInfoApi) CreateTeacherInfo(c *gin.Context) {
	var teacherInfo autocode.TeacherInfo
	_ = c.ShouldBindJSON(&teacherInfo)
	if err := teacherInfoService.CreateTeacherInfo(teacherInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeacherInfo 删除TeacherInfo
// @Tags TeacherInfo
// @Summary 删除TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeacherInfo true "删除TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacherInfo/deleteTeacherInfo [delete]
func (teacherInfoApi *TeacherInfoApi) DeleteTeacherInfo(c *gin.Context) {
	var teacherInfo autocode.TeacherInfo
	_ = c.ShouldBindJSON(&teacherInfo)
	if err := teacherInfoService.DeleteTeacherInfo(teacherInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeacherInfoByIds 批量删除TeacherInfo
// @Tags TeacherInfo
// @Summary 批量删除TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teacherInfo/deleteTeacherInfoByIds [delete]
func (teacherInfoApi *TeacherInfoApi) DeleteTeacherInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := teacherInfoService.DeleteTeacherInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeacherInfo 更新TeacherInfo
// @Tags TeacherInfo
// @Summary 更新TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeacherInfo true "更新TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teacherInfo/updateTeacherInfo [put]
func (teacherInfoApi *TeacherInfoApi) UpdateTeacherInfo(c *gin.Context) {
	var teacherInfo autocode.TeacherInfo
	_ = c.ShouldBindJSON(&teacherInfo)
	if err := teacherInfoService.UpdateTeacherInfo(teacherInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeacherInfo 用id查询TeacherInfo
// @Tags TeacherInfo
// @Summary 用id查询TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TeacherInfo true "用id查询TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teacherInfo/findTeacherInfo [get]
func (teacherInfoApi *TeacherInfoApi) FindTeacherInfo(c *gin.Context) {
	var teacherInfo autocode.TeacherInfo
	_ = c.ShouldBindQuery(&teacherInfo)
	if err, reteacherInfo := teacherInfoService.GetTeacherInfo(teacherInfo.UId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteacherInfo": reteacherInfo}, c)
	}
}

// GetTeacherInfoList 分页获取TeacherInfo列表
// @Tags TeacherInfo
// @Summary 分页获取TeacherInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TeacherInfoSearch true "分页获取TeacherInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacherInfo/getTeacherInfoList [get]
func (teacherInfoApi *TeacherInfoApi) GetTeacherInfoList(c *gin.Context) {
	var pageInfo autocodeReq.TeacherInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := teacherInfoService.GetTeacherInfoInfoList(pageInfo); err != nil {
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
