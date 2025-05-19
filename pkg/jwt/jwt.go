package jwt

import (
	"fmt"
	"strings"

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

// GenerateJwt 签名方法为 RS256
func GenerateJwt(claims UserClaims, secret interface{}) (string, error) {
	j := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return j.SignedString(secret)
}
