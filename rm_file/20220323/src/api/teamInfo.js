import service from '@/utils/request'

// @Tags TeamInfo
// @Summary 创建TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamInfo true "创建TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamInfo/createTeamInfo [post]
export const createTeamInfo = (data) => {
  return service({
    url: '/teamInfo/createTeamInfo',
    method: 'post',
    data
  })
}

// @Tags TeamInfo
// @Summary 删除TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamInfo true "删除TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamInfo/deleteTeamInfo [delete]
export const deleteTeamInfo = (data) => {
  return service({
    url: '/teamInfo/deleteTeamInfo',
    method: 'delete',
    data
  })
}

// @Tags TeamInfo
// @Summary 删除TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /teamInfo/deleteTeamInfo [delete]
export const deleteTeamInfoByIds = (data) => {
  return service({
    url: '/teamInfo/deleteTeamInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags TeamInfo
// @Summary 更新TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TeamInfo true "更新TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /teamInfo/updateTeamInfo [put]
export const updateTeamInfo = (data) => {
  return service({
    url: '/teamInfo/updateTeamInfo',
    method: 'put',
    data
  })
}

// @Tags TeamInfo
// @Summary 用id查询TeamInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TeamInfo true "用id查询TeamInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /teamInfo/findTeamInfo [get]
export const findTeamInfo = (params) => {
  return service({
    url: '/teamInfo/findTeamInfo',
    method: 'get',
    params
  })
}

// @Tags TeamInfo
// @Summary 分页获取TeamInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TeamInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /teamInfo/getTeamInfoList [get]
export const getTeamInfoList = (params) => {
  return service({
    url: '/teamInfo/getTeamInfoList',
    method: 'get',
    params
  })
}
