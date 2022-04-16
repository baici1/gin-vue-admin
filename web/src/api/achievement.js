import service from '@/utils/request'

// @Tags Achievement
// @Summary 创建Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Achievement true "创建Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /achievement/createAchievement [post]
export const createAchievement = (data) => {
  return service({
    url: '/achievement/createAchievement',
    method: 'post',
    data
  })
}

// @Tags Achievement
// @Summary 删除Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Achievement true "删除Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /achievement/deleteAchievement [delete]
export const deleteAchievement = (data) => {
  return service({
    url: '/achievement/deleteAchievement',
    method: 'delete',
    data
  })
}

// @Tags Achievement
// @Summary 删除Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /achievement/deleteAchievement [delete]
export const deleteAchievementByIds = (data) => {
  return service({
    url: '/achievement/deleteAchievementByIds',
    method: 'delete',
    data
  })
}

// @Tags Achievement
// @Summary 更新Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Achievement true "更新Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /achievement/updateAchievement [put]
export const updateAchievement = (data) => {
  return service({
    url: '/achievement/updateAchievement',
    method: 'put',
    data
  })
}

// @Tags Achievement
// @Summary 用id查询Achievement
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Achievement true "用id查询Achievement"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /achievement/findAchievement [get]
export const findAchievement = (params) => {
  return service({
    url: '/achievement/findAchievement',
    method: 'get',
    params
  })
}

// @Tags Achievement
// @Summary 分页获取Achievement列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Achievement列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /achievement/getAchievementList [get]
export const getAchievementList = (params) => {
  return service({
    url: '/achievement/getAchievementList',
    method: 'get',
    params
  })
}
