package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

var signingKey = []byte(viper.GetString("jwt.signingKey"))

type JwtCustClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

// 生成token
func GenerateToken(id int, name string) (string, error) {
	jwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustClaims)
	return token.SignedString(signingKey)
}

// 解析token
func ParseToken(token string) (JwtCustClaims, error) {
	jwtCustClaims := JwtCustClaims{}
	claims, err := jwt.ParseWithClaims(token, &jwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err == nil && !claims.Valid {
		err = errors.New("Invalid Token")
	}

	return jwtCustClaims, err
}

// 判断token
func IsTokenValid(token string) bool {
	_, err := ParseToken(token)
	if err != nil {
		return false
	}
	return true
}
