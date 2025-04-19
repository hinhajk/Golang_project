package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTsecret = []byte("ABAB")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	jwt.StandardClaims
}

// GenerateToken token签发
func GenerateToken(id uint, userName string, passWord string) (tokenString string, err error) {
	notTime := time.Now()
	expireTime := notTime.AddDate(1, 0, 0)
	claims := Claims{
		Id:       id,
		UserName: userName,
		PassWord: passWord,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blogSystem",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
