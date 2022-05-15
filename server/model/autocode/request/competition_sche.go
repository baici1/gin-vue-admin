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
	CType      int    `json:"ctype" form:"ctype"`
	SearchInfo string `json:"searchInfo" form:"searchInfo"`
	autocode.CompetitionSche
	request.PageInfo
}
