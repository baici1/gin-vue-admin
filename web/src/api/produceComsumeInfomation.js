import service from '@/utils/request'

// @Tags ProduceComsumeInfomation
// @Summary 创建ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProduceComsumeInfomation true "创建ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /produceComsumeInfomation/createProduceComsumeInfomation [post]
export const createProduceComsumeInfomation = (data) => {
  return service({
    url: '/produceComsumeInfomation/createProduceComsumeInfomation',
    method: 'post',
    data
  })
}

// @Tags ProduceComsumeInfomation
// @Summary 删除ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProduceComsumeInfomation true "删除ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /produceComsumeInfomation/deleteProduceComsumeInfomation [delete]
export const deleteProduceComsumeInfomation = (data) => {
  return service({
    url: '/produceComsumeInfomation/deleteProduceComsumeInfomation',
    method: 'delete',
    data
  })
}

// @Tags ProduceComsumeInfomation
// @Summary 删除ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /produceComsumeInfomation/deleteProduceComsumeInfomation [delete]
export const deleteProduceComsumeInfomationByIds = (data) => {
  return service({
    url: '/produceComsumeInfomation/deleteProduceComsumeInfomationByIds',
    method: 'delete',
    data
  })
}

// @Tags ProduceComsumeInfomation
// @Summary 更新ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProduceComsumeInfomation true "更新ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /produceComsumeInfomation/updateProduceComsumeInfomation [put]
export const updateProduceComsumeInfomation = (data) => {
  return service({
    url: '/produceComsumeInfomation/updateProduceComsumeInfomation',
    method: 'put',
    data
  })
}

// @Tags ProduceComsumeInfomation
// @Summary 用id查询ProduceComsumeInfomation
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProduceComsumeInfomation true "用id查询ProduceComsumeInfomation"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /produceComsumeInfomation/findProduceComsumeInfomation [get]
export const findProduceComsumeInfomation = (params) => {
  return service({
    url: '/produceComsumeInfomation/findProduceComsumeInfomation',
    method: 'get',
    params
  })
}

// @Tags ProduceComsumeInfomation
// @Summary 分页获取ProduceComsumeInfomation列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ProduceComsumeInfomation列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /produceComsumeInfomation/getProduceComsumeInfomationList [get]
export const getProduceComsumeInfomationList = (params) => {
  return service({
    url: '/produceComsumeInfomation/getProduceComsumeInfomationList',
    method: 'get',
    params
  })
}
