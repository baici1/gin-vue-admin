package request

import "github.com/dgrijalva/jwt-go"

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	Phone       string
	AuthorityId string
	ID          uint
}
