package beego

import (
	"github.com/dgrijalva/jwt-go"
)

// UserClaims js jwt claims
type UserClaims struct {
	UID int `json:"uid"`
	jwt.StandardClaims
}