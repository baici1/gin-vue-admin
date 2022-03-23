import service from '@/utils/request'

// @Tags EntryForm
// @Summary 创建EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EntryForm true "创建EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryForm/createEntryForm [post]
export const createEntryForm = (data) => {
  return service({
    url: '/entryForm/createEntryForm',
    method: 'post',
    data
  })
}

// @Tags EntryForm
// @Summary 删除EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EntryForm true "删除EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /entryForm/deleteEntryForm [delete]
export const deleteEntryForm = (data) => {
  return service({
    url: '/entryForm/deleteEntryForm',
    method: 'delete',
    data
  })
}

// @Tags EntryForm
// @Summary 删除EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /entryForm/deleteEntryForm [delete]
export const deleteEntryFormByIds = (data) => {
  return service({
    url: '/entryForm/deleteEntryFormByIds',
    method: 'delete',
    data
  })
}

// @Tags EntryForm
// @Summary 更新EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EntryForm true "更新EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /entryForm/updateEntryForm [put]
export const updateEntryForm = (data) => {
  return service({
    url: '/entryForm/updateEntryForm',
    method: 'put',
    data
  })
}

// @Tags EntryForm
// @Summary 用id查询EntryForm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EntryForm true "用id查询EntryForm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /entryForm/findEntryForm [get]
export const findEntryForm = (params) => {
  return service({
    url: '/entryForm/findEntryForm',
    method: 'get',
    params
  })
}

// @Tags EntryForm
// @Summary 分页获取EntryForm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EntryForm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryForm/getEntryFormList [get]
export const getEntryFormList = (params) => {
  return service({
    url: '/entryForm/getEntryFormList',
    method: 'get',
    params
  })
}
