package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Level int `json:"level"`
}

type User struct {
	UserId string
	Level  int
}

func ConvertSigningMethod(method string) (jwt.SigningMethod, error) {
	switch strings.ToUpper(method) {
	case "RS256":
		return jwt.SigningMethodRS256, nil
	default:
		return nil, fmt.Errorf("unsupported signing method: %s", method)
	}
}

// GenerateJwt 签名方法为 HS256
func GenerateJwt(user *User, secret interface{}, expiresAt time.Time, method string) (string, error) {
	signingMethod, err := ConvertSigningMethod(method)
	if err != nil {
		return "", err
	}

	j := jwt.NewWithClaims(signingMethod, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.UserId,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		Level: user.Level,
	})
	return j.SignedString(secret)
}

// ParseJWT 解析 token string 到结构体，如果过期或签名无效返回错误
func ParseJWT(tokenStr string, secret string) (*UserClaims, error) {
	claims := UserClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return &claims, err
}
