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

type TeacherRecruitApi struct {
}

var teacherRecruitService = service.ServiceGroupApp.AutoCodeServiceGroup.TeacherRecruitService


// CreateTeacherRecruit 创建TeacherRecruit
// @Tags TeacherRecruit
// @Summary 创建TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeacherRecruit true "创建TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacherRecruit/createTeacherRecruit [post]
func (teacherRecruitApi *TeacherRecruitApi) CreateTeacherRecruit(c *gin.Context) {
	var teacherRecruit autocode.TeacherRecruit
	_ = c.ShouldBindJSON(&teacherRecruit)
	if err := teacherRecruitService.CreateTeacherRecruit(teacherRecruit); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTeacherRecruit 删除TeacherRecruit
// @Tags TeacherRecruit
// @Summary 删除TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeacherRecruit true "删除TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacherRecruit/deleteTeacherRecruit [delete]
func (teacherRecruitApi *TeacherRecruitApi) DeleteTeacherRecruit(c *gin.Context) {
	var teacherRecruit autocode.TeacherRecruit
	_ = c.ShouldBindJSON(&teacherRecruit)
	if err := teacherRecruitService.DeleteTeacherRecruit(teacherRecruit); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTeacherRecruitByIds 批量删除TeacherRecruit
// @Tags TeacherRecruit
// @Summary 批量删除TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /teacherRecruit/deleteTeacherRecruitByIds [delete]
func (teacherRecruitApi *TeacherRecruitApi) DeleteTeacherRecruitByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := teacherRecruitService.DeleteTeacherRecruitByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTeacherRecruit 更新TeacherRecruit
// @Tags TeacherRecruit
// @Summary 更新TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.TeacherRecruit true "更新TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teacherRecruit/updateTeacherRecruit [put]
func (teacherRecruitApi *TeacherRecruitApi) UpdateTeacherRecruit(c *gin.Context) {
	var teacherRecruit autocode.TeacherRecruit
	_ = c.ShouldBindJSON(&teacherRecruit)
	if err := teacherRecruitService.UpdateTeacherRecruit(teacherRecruit); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTeacherRecruit 用id查询TeacherRecruit
// @Tags TeacherRecruit
// @Summary 用id查询TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.TeacherRecruit true "用id查询TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teacherRecruit/findTeacherRecruit [get]
func (teacherRecruitApi *TeacherRecruitApi) FindTeacherRecruit(c *gin.Context) {
	var teacherRecruit autocode.TeacherRecruit
	_ = c.ShouldBindQuery(&teacherRecruit)
	if err, reteacherRecruit := teacherRecruitService.GetTeacherRecruit(teacherRecruit.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reteacherRecruit": reteacherRecruit}, c)
	}
}

// GetTeacherRecruitList 分页获取TeacherRecruit列表
// @Tags TeacherRecruit
// @Summary 分页获取TeacherRecruit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.TeacherRecruitSearch true "分页获取TeacherRecruit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacherRecruit/getTeacherRecruitList [get]
func (teacherRecruitApi *TeacherRecruitApi) GetTeacherRecruitList(c *gin.Context) {
	var pageInfo autocodeReq.TeacherRecruitSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := teacherRecruitService.GetTeacherRecruitInfoList(pageInfo); err != nil {
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
