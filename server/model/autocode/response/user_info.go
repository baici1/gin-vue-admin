package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
)

type LoginResponse struct {
	User      autocode.UserInfo `json:"user"`
	Token     string            `json:"token"`
	ExpiresAt int64             `json:"expiresAt"`
}
