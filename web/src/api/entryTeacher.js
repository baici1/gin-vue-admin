import service from '@/utils/request'

// @Tags EntryTeacher
// @Summary 创建EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EntryTeacher true "创建EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryTeacher/createEntryTeacher [post]
export const createEntryTeacher = (data) => {
  return service({
    url: '/entryTeacher/createEntryTeacher',
    method: 'post',
    data
  })
}

// @Tags EntryTeacher
// @Summary 删除EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EntryTeacher true "删除EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /entryTeacher/deleteEntryTeacher [delete]
export const deleteEntryTeacher = (data) => {
  return service({
    url: '/entryTeacher/deleteEntryTeacher',
    method: 'delete',
    data
  })
}

// @Tags EntryTeacher
// @Summary 删除EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /entryTeacher/deleteEntryTeacher [delete]
export const deleteEntryTeacherByIds = (data) => {
  return service({
    url: '/entryTeacher/deleteEntryTeacherByIds',
    method: 'delete',
    data
  })
}

// @Tags EntryTeacher
// @Summary 更新EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EntryTeacher true "更新EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /entryTeacher/updateEntryTeacher [put]
export const updateEntryTeacher = (data) => {
  return service({
    url: '/entryTeacher/updateEntryTeacher',
    method: 'put',
    data
  })
}

// @Tags EntryTeacher
// @Summary 用id查询EntryTeacher
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EntryTeacher true "用id查询EntryTeacher"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /entryTeacher/findEntryTeacher [get]
export const findEntryTeacher = (params) => {
  return service({
    url: '/entryTeacher/findEntryTeacher',
    method: 'get',
    params
  })
}

// @Tags EntryTeacher
// @Summary 分页获取EntryTeacher列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EntryTeacher列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /entryTeacher/getEntryTeacherList [get]
export const getEntryTeacherList = (params) => {
  return service({
    url: '/entryTeacher/getEntryTeacherList',
    method: 'get',
    params
  })
}
