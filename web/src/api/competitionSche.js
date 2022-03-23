import service from '@/utils/request'

// @Tags CompetitionSche
// @Summary 创建CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompetitionSche true "创建CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /competitionSche/createCompetitionSche [post]
export const createCompetitionSche = (data) => {
  return service({
    url: '/competitionSche/createCompetitionSche',
    method: 'post',
    data
  })
}

// @Tags CompetitionSche
// @Summary 删除CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompetitionSche true "删除CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /competitionSche/deleteCompetitionSche [delete]
export const deleteCompetitionSche = (data) => {
  return service({
    url: '/competitionSche/deleteCompetitionSche',
    method: 'delete',
    data
  })
}

// @Tags CompetitionSche
// @Summary 删除CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /competitionSche/deleteCompetitionSche [delete]
export const deleteCompetitionScheByIds = (data) => {
  return service({
    url: '/competitionSche/deleteCompetitionScheByIds',
    method: 'delete',
    data
  })
}

// @Tags CompetitionSche
// @Summary 更新CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompetitionSche true "更新CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /competitionSche/updateCompetitionSche [put]
export const updateCompetitionSche = (data) => {
  return service({
    url: '/competitionSche/updateCompetitionSche',
    method: 'put',
    data
  })
}

// @Tags CompetitionSche
// @Summary 用id查询CompetitionSche
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CompetitionSche true "用id查询CompetitionSche"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /competitionSche/findCompetitionSche [get]
export const findCompetitionSche = (params) => {
  return service({
    url: '/competitionSche/findCompetitionSche',
    method: 'get',
    params
  })
}

// @Tags CompetitionSche
// @Summary 分页获取CompetitionSche列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CompetitionSche列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /competitionSche/getCompetitionScheList [get]
export const getCompetitionScheList = (params) => {
  return service({
    url: '/competitionSche/getCompetitionScheList',
    method: 'get',
    params
  })
}
