package jwt

import (
	"m-server-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SessionUserInfo struct {
	Id       int64  `json:"id"`       // ID
	TenantId int64  `json:"tenantId"` // 租户ID
	Platform string `json:"platform"` // 平台
}

type Claims struct {
	SessionUserInfo
	jwt.RegisteredClaims
}

// 生成JWT令牌
func GenerateToken(sessionUserInfo SessionUserInfo, expireDuration time.Duration) (string, error) {
	claims := Claims{
		sessionUserInfo,
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Get().Jwt.Secret))
	return token, err
}

// 解析JWT令牌
func ParseToken(tokenString string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Get().Jwt.Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
