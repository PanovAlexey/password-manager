package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	jwt_secret = "eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9"
	tokenTTL   = 12 * time.Hour
)

type claims struct {
	jwt.StandardClaims
	UserId string `json:"username"`
}

type JWTAuthorization struct {
}

func GetJWTAuthorizationService() JWTAuthorization {
	return JWTAuthorization{}
}

func (u JWTAuthorization) GetJWTToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(jwt_secret))
}

func (u JWTAuthorization) CheckGetJWTToken(accessToken string) (*claims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(jwt_secret), nil
	})

	if err != nil {
		return nil, err
	}

	payload, ok := token.Claims.(*claims)

	if !ok {
		return nil, errors.New("token claims are not of type *claims")
	}

	return payload, nil
}
