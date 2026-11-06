package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func getid(c *gin.Context) (int, error) {
	claims := &Claims{}
	tokentext, err := c.Cookie("login")
	if err != nil {
		return 0, fmt.Errorf("无cookie")
	}
	token, err := jwt.ParseWithClaims(
		tokentext,
		claims,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("error in signing method ")
			}
			return []byte(cfg.JWTSigningSecret), nil
		})
	if err != nil || token.Valid == false {
		return 0, fmt.Errorf("error in cookie")
	}
	return claims.Userid, nil
}
