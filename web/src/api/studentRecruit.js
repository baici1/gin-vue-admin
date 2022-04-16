import service from '@/utils/request'

// @Tags StudentRecruit
// @Summary 创建StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentRecruit true "创建StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /studentRecruit/createStudentRecruit [post]
export const createStudentRecruit = (data) => {
  return service({
    url: '/studentRecruit/createStudentRecruit',
    method: 'post',
    data
  })
}

// @Tags StudentRecruit
// @Summary 删除StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentRecruit true "删除StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /studentRecruit/deleteStudentRecruit [delete]
export const deleteStudentRecruit = (data) => {
  return service({
    url: '/studentRecruit/deleteStudentRecruit',
    method: 'delete',
    data
  })
}

// @Tags StudentRecruit
// @Summary 删除StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /studentRecruit/deleteStudentRecruit [delete]
export const deleteStudentRecruitByIds = (data) => {
  return service({
    url: '/studentRecruit/deleteStudentRecruitByIds',
    method: 'delete',
    data
  })
}

// @Tags StudentRecruit
// @Summary 更新StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StudentRecruit true "更新StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /studentRecruit/updateStudentRecruit [put]
export const updateStudentRecruit = (data) => {
  return service({
    url: '/studentRecruit/updateStudentRecruit',
    method: 'put',
    data
  })
}

// @Tags StudentRecruit
// @Summary 用id查询StudentRecruit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.StudentRecruit true "用id查询StudentRecruit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /studentRecruit/findStudentRecruit [get]
export const findStudentRecruit = (params) => {
  return service({
    url: '/studentRecruit/findStudentRecruit',
    method: 'get',
    params
  })
}

// @Tags StudentRecruit
// @Summary 分页获取StudentRecruit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取StudentRecruit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /studentRecruit/getStudentRecruitList [get]
export const getStudentRecruitList = (params) => {
  return service({
    url: '/studentRecruit/getStudentRecruitList',
    method: 'get',
    params
  })
}
