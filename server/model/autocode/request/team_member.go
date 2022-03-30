package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type TeamMemberSearch struct {
	autocode.TeamMember
	request.PageInfo
}

type OwnTeamMember struct {
	Phone string `json:"phone" form:"phone"`
	autocode.TeamMember
}
