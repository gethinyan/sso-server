package util

import (
	"time"

	"github.com/gethinyan/sso/pkg/setting"
	"github.com/dgrijalva/jwt-go"
)

// UID token 解析出的用户 ID
var UID uint

var jwtSecret = []byte(setting.App.JwtSecret)

// Claims ...
type Claims struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	ValidBefore int64  `json:"valid_before"`
	jwt.StandardClaims
}

// GetTokenExpireTime 获取 token 过期时间
func GetTokenExpireTime() int64 {
	return time.Now().Add(30 * time.Second).Unix()
}

// GetRefreshTokenExpireTime 获取 refresh token 过期时间
func GetRefreshTokenExpireTime() int64 {
	return time.Now().Add(30 * 24 * time.Hour).Unix()
}

// GenerateToken 生成 token
func GenerateToken(id uint, email string) (string, error) {
	claims := Claims{
		id,
		email,
		GetTokenExpireTime(),
		jwt.StandardClaims{
			ExpiresAt: GetRefreshTokenExpireTime(),
			Issuer:    "sso",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 解析 token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			UID = claims.ID
			return claims, nil
		}
	}

	return nil, err
}
