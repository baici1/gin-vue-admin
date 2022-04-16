import service from '@/utils/request'

// @Tags TeacherRecruit
// @Summary 创建TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeacherRecruit true "创建TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacherRecruit/createTeacherRecruit [post]
export const createTeacherRecruit = (data) => {
  return service({
    url: '/teacherRecruit/createTeacherRecruit',
    method: 'post',
    data
  })
}

// @Tags TeacherRecruit
// @Summary 删除TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeacherRecruit true "删除TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacherRecruit/deleteTeacherRecruit [delete]
export const deleteTeacherRecruit = (data) => {
  return service({
    url: '/teacherRecruit/deleteTeacherRecruit',
    method: 'delete',
    data
  })
}

// @Tags TeacherRecruit
// @Summary 删除TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacherRecruit/deleteTeacherRecruit [delete]
export const deleteTeacherRecruitByIds = (data) => {
  return service({
    url: '/teacherRecruit/deleteTeacherRecruitByIds',
    method: 'delete',
    data
  })
}

// @Tags TeacherRecruit
// @Summary 更新TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeacherRecruit true "更新TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teacherRecruit/updateTeacherRecruit [put]
export const updateTeacherRecruit = (data) => {
  return service({
    url: '/teacherRecruit/updateTeacherRecruit',
    method: 'put',
    data
  })
}

// @Tags TeacherRecruit
// @Summary 用id查询TeacherRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeacherRecruit true "用id查询TeacherRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teacherRecruit/findTeacherRecruit [get]
export const findTeacherRecruit = (params) => {
  return service({
    url: '/teacherRecruit/findTeacherRecruit',
    method: 'get',
    params
  })
}

// @Tags TeacherRecruit
// @Summary 分页获取TeacherRecruit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeacherRecruit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacherRecruit/getTeacherRecruitList [get]
export const getTeacherRecruitList = (params) => {
  return service({
    url: '/teacherRecruit/getTeacherRecruitList',
    method: 'get',
    params
  })
}
