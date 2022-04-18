package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProduceComsumeInfomationSearch struct{
    autocode.ProduceComsumeInfomation
    request.PageInfo
}