package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	hmacSampleSecret = "rahasiadong"
	expiredPeriod    = 90
)

func JwtEncoder(userName, customKey string) (string, error) {
	expiredDate := time.Now().Add(time.Second * expiredPeriod)
	claims := jwt.MapClaims{
		"name":      userName,
		"customKey": customKey,
		"expiredAt": expiredDate.Format("2006-01-02 15:04:22"),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}

func JwtDecoder(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hmacSampleSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println(claims["name"], claims["customKey"])
		expiredAt := claims["expiredAt"].(string)
		timeExpired, err := time.Parse("2006-01-02 15:04:05", expiredAt)
		if err != nil {
			return nil, err
		}

		thisTimeString := time.Now().Format("2006-01-02 15:04:05")
		thisTime, err := time.Parse("2006-01-02 15:04:05", thisTimeString)

		diffTime := timeExpired.Sub(thisTime).Seconds()
		if diffTime > 0 {
			return claims, nil
		}
		return nil, errors.New("Expired Token")
	} else {
		return nil, err
	}
}
