package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var JwtSecret = []byte("NigTusg")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 通过token可以获得用户的id和username
func GenerateToken(id uint, username string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todolist_db",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

// ParseToken 用来验证用户Token
// 解析Token的过程就是生成Token的逆向过程，即为先获得JwtSecret后，再依次对Token进行拆解
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			//tokenClaims.Valid 是一个布尔值，用于表示 JWT（JSON Web Token）是否有效，并且需要检查 JWT 的签名是否有效、是否在有效期内等
			return claims, nil
		}
	}
	return nil, err
}
