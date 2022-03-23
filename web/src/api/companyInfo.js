import service from '@/utils/request'

// @Tags CompanyInfo
// @Summary 创建CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompanyInfo true "创建CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyInfo/createCompanyInfo [post]
export const createCompanyInfo = (data) => {
  return service({
    url: '/companyInfo/createCompanyInfo',
    method: 'post',
    data
  })
}

// @Tags CompanyInfo
// @Summary 删除CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompanyInfo true "删除CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /companyInfo/deleteCompanyInfo [delete]
export const deleteCompanyInfo = (data) => {
  return service({
    url: '/companyInfo/deleteCompanyInfo',
    method: 'delete',
    data
  })
}

// @Tags CompanyInfo
// @Summary 删除CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /companyInfo/deleteCompanyInfo [delete]
export const deleteCompanyInfoByIds = (data) => {
  return service({
    url: '/companyInfo/deleteCompanyInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags CompanyInfo
// @Summary 更新CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompanyInfo true "更新CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /companyInfo/updateCompanyInfo [put]
export const updateCompanyInfo = (data) => {
  return service({
    url: '/companyInfo/updateCompanyInfo',
    method: 'put',
    data
  })
}

// @Tags CompanyInfo
// @Summary 用id查询CompanyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CompanyInfo true "用id查询CompanyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /companyInfo/findCompanyInfo [get]
export const findCompanyInfo = (params) => {
  return service({
    url: '/companyInfo/findCompanyInfo',
    method: 'get',
    params
  })
}

// @Tags CompanyInfo
// @Summary 分页获取CompanyInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CompanyInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /companyInfo/getCompanyInfoList [get]
export const getCompanyInfoList = (params) => {
  return service({
    url: '/companyInfo/getCompanyInfoList',
    method: 'get',
    params
  })
}
