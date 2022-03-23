import service from '@/utils/request'

// @Tags CompetitionInfo
// @Summary 创建CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompetitionInfo true "创建CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /competitionInfo/createCompetitionInfo [post]
export const createCompetitionInfo = (data) => {
  return service({
    url: '/competitionInfo/createCompetitionInfo',
    method: 'post',
    data
  })
}

// @Tags CompetitionInfo
// @Summary 删除CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompetitionInfo true "删除CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /competitionInfo/deleteCompetitionInfo [delete]
export const deleteCompetitionInfo = (data) => {
  return service({
    url: '/competitionInfo/deleteCompetitionInfo',
    method: 'delete',
    data
  })
}

// @Tags CompetitionInfo
// @Summary 删除CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /competitionInfo/deleteCompetitionInfo [delete]
export const deleteCompetitionInfoByIds = (data) => {
  return service({
    url: '/competitionInfo/deleteCompetitionInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags CompetitionInfo
// @Summary 更新CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompetitionInfo true "更新CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /competitionInfo/updateCompetitionInfo [put]
export const updateCompetitionInfo = (data) => {
  return service({
    url: '/competitionInfo/updateCompetitionInfo',
    method: 'put',
    data
  })
}

// @Tags CompetitionInfo
// @Summary 用id查询CompetitionInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CompetitionInfo true "用id查询CompetitionInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /competitionInfo/findCompetitionInfo [get]
export const findCompetitionInfo = (params) => {
  return service({
    url: '/competitionInfo/findCompetitionInfo',
    method: 'get',
    params
  })
}

// @Tags CompetitionInfo
// @Summary 分页获取CompetitionInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CompetitionInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /competitionInfo/getCompetitionInfoList [get]
export const getCompetitionInfoList = (params) => {
  return service({
    url: '/competitionInfo/getCompetitionInfoList',
    method: 'get',
    params
  })
}
