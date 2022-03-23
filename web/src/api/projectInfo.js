import service from '@/utils/request'

// @Tags ProjectInfo
// @Summary 创建ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectInfo true "创建ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /projectInfo/createProjectInfo [post]
export const createProjectInfo = (data) => {
  return service({
    url: '/projectInfo/createProjectInfo',
    method: 'post',
    data
  })
}

// @Tags ProjectInfo
// @Summary 删除ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectInfo true "删除ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /projectInfo/deleteProjectInfo [delete]
export const deleteProjectInfo = (data) => {
  return service({
    url: '/projectInfo/deleteProjectInfo',
    method: 'delete',
    data
  })
}

// @Tags ProjectInfo
// @Summary 删除ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /projectInfo/deleteProjectInfo [delete]
export const deleteProjectInfoByIds = (data) => {
  return service({
    url: '/projectInfo/deleteProjectInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags ProjectInfo
// @Summary 更新ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectInfo true "更新ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /projectInfo/updateProjectInfo [put]
export const updateProjectInfo = (data) => {
  return service({
    url: '/projectInfo/updateProjectInfo',
    method: 'put',
    data
  })
}

// @Tags ProjectInfo
// @Summary 用id查询ProjectInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProjectInfo true "用id查询ProjectInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /projectInfo/findProjectInfo [get]
export const findProjectInfo = (params) => {
  return service({
    url: '/projectInfo/findProjectInfo',
    method: 'get',
    params
  })
}

// @Tags ProjectInfo
// @Summary 分页获取ProjectInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ProjectInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /projectInfo/getProjectInfoList [get]
export const getProjectInfoList = (params) => {
  return service({
    url: '/projectInfo/getProjectInfoList',
    method: 'get',
    params
  })
}
