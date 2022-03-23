import service from '@/utils/request'

// @Tags EntryMember
// @Summary 创建EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EntryMember true "创建EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryMember/createEntryMember [post]
export const createEntryMember = (data) => {
  return service({
    url: '/entryMember/createEntryMember',
    method: 'post',
    data
  })
}

// @Tags EntryMember
// @Summary 删除EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EntryMember true "删除EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /entryMember/deleteEntryMember [delete]
export const deleteEntryMember = (data) => {
  return service({
    url: '/entryMember/deleteEntryMember',
    method: 'delete',
    data
  })
}

// @Tags EntryMember
// @Summary 删除EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /entryMember/deleteEntryMember [delete]
export const deleteEntryMemberByIds = (data) => {
  return service({
    url: '/entryMember/deleteEntryMemberByIds',
    method: 'delete',
    data
  })
}

// @Tags EntryMember
// @Summary 更新EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EntryMember true "更新EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /entryMember/updateEntryMember [put]
export const updateEntryMember = (data) => {
  return service({
    url: '/entryMember/updateEntryMember',
    method: 'put',
    data
  })
}

// @Tags EntryMember
// @Summary 用id查询EntryMember
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EntryMember true "用id查询EntryMember"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /entryMember/findEntryMember [get]
export const findEntryMember = (params) => {
  return service({
    url: '/entryMember/findEntryMember',
    method: 'get',
    params
  })
}

// @Tags EntryMember
// @Summary 分页获取EntryMember列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EntryMember列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryMember/getEntryMemberList [get]
export const getEntryMemberList = (params) => {
  return service({
    url: '/entryMember/getEntryMemberList',
    method: 'get',
    params
  })
}
