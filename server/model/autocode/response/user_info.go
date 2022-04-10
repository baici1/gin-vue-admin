package response

type LoginResponse struct {
	User      interface{} `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}
