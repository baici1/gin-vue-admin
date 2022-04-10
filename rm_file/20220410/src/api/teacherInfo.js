import service from '@/utils/request'

// @Tags TeacherInfo
// @Summary 创建TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeacherInfo true "创建TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacherInfo/createTeacherInfo [post]
export const createTeacherInfo = (data) => {
  return service({
    url: '/teacherInfo/createTeacherInfo',
    method: 'post',
    data
  })
}

// @Tags TeacherInfo
// @Summary 删除TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeacherInfo true "删除TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacherInfo/deleteTeacherInfo [delete]
export const deleteTeacherInfo = (data) => {
  return service({
    url: '/teacherInfo/deleteTeacherInfo',
    method: 'delete',
    data
  })
}

// @Tags TeacherInfo
// @Summary 删除TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teacherInfo/deleteTeacherInfo [delete]
export const deleteTeacherInfoByIds = (data) => {
  return service({
    url: '/teacherInfo/deleteTeacherInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags TeacherInfo
// @Summary 更新TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeacherInfo true "更新TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teacherInfo/updateTeacherInfo [put]
export const updateTeacherInfo = (data) => {
  return service({
    url: '/teacherInfo/updateTeacherInfo',
    method: 'put',
    data
  })
}

// @Tags TeacherInfo
// @Summary 用id查询TeacherInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeacherInfo true "用id查询TeacherInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teacherInfo/findTeacherInfo [get]
export const findTeacherInfo = (params) => {
  return service({
    url: '/teacherInfo/findTeacherInfo',
    method: 'get',
    params
  })
}

// @Tags TeacherInfo
// @Summary 分页获取TeacherInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeacherInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teacherInfo/getTeacherInfoList [get]
export const getTeacherInfoList = (params) => {
  return service({
    url: '/teacherInfo/getTeacherInfoList',
    method: 'get',
    params
  })
}
