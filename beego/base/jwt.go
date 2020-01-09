package base

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserClaims js jwt claims
type UserClaims struct {
	UID int `json:"uid"`
	jwt.StandardClaims
}

// MemoCreateToken ...
func MemoCreateToken(uid int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		uid,
		jwt.StandardClaims{
			Id:        "1f2a23a63ad",
			Issuer:    "https://www.smemo.info",
			Audience:  "https://www.smemo.info",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(60 * 24 * 360 * time.Minute).Unix(),
		},
	})
	token.Header["jti"] = "1f2a23a63ad"

	tokenString, _ := token.SignedString([]byte("smemo_info_user_token"))
	return tokenString

}

// MemoValidatingToken ...
func MemoValidatingToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("smemo_info_user_token"), nil
	})

	if err != nil || token == nil || token != nil && !token.Valid {
		return -1, fmt.Errorf("token is invalid " + err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), false) == false {
			return -1, fmt.Errorf("token is expired")
		}
		if claims["jti"] != "1f2a23a63ad" {
			return -1, fmt.Errorf("jti is invalid")
		}
		return int(claims["uid"].(float64)), nil
	}
	return -1, err
}
