package neo

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// UserClaims js jwt claims
type UserClaims struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

type JWTConfig struct {
	Id        string
	Issuer    string
	ExpiresAt int64
	Key       string
}

//time.Now().Add(60 * 24 * 360 * time.Minute).Unix()
func CreateToken(uid string, config JWTConfig) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		uid,
		jwt.StandardClaims{
			Id:        config.Id,
			Issuer:    config.Issuer,
			Audience:  "https://neobaran.com",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: config.ExpiresAt,
		},
	})
	token.Header["jti"] = config.Id

	tokenString, _ := token.SignedString([]byte(config.Key))
	return tokenString

}

func ValidatingToken(tokenString string, config JWTConfig) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Key), nil
	})

	if err != nil || token == nil || token != nil && !token.Valid {
		return "", fmt.Errorf("token is invalid " + err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), false) == false {
			return "", fmt.Errorf("token is expired")
		}
		if claims["jti"] != config.Id {
			return "", fmt.Errorf("jti is invalid")
		}
		return claims["uid"].(string), nil
	}
	return "", err
}
