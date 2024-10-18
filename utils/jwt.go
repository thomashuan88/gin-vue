package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/spf13/viper"
)

var stSecretKey = []byte(viper.GetString("jwt.secretKey"))

type JwtCustClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id int, name string) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(viper.GetInt("jwt.tokenExpire")) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)

	return token.SignedString(stSecretKey)
}

func ParseToken(tokenString string) (JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return stSecretKey, nil
	})

	if err == nil && !token.Valid {
		err = errors.New("invalid token")
	}

	return iJwtCustClaims, err
}

func IsTokenValid(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}
