package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CompetitionScheSearch struct {
	autocode.CompetitionSche
	request.PageInfo
}

type CompetitionDetailSearch struct {
	Status     int    `json:"status" form:"status"`
	SearchInfo string `json:"searchInfo" form:"searchInfo"`
	autocode.CompetitionSche
	request.PageInfo
}
