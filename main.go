package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"jwt-security/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const KEY  = "1325342D59DF41D20335A1898D07D163"

func main() {
	e := echo.New()

	e.GET("/token", func(c echo.Context) error {
		user := model.User{
			UserId:   "123456",
			FullName: "Code4Func",
		}

		token, _ := GenToken(user)
		fmt.Println(token)

		tokenDecrypt, err := encrypt([]byte(KEY), token)
		if err != nil {
			fmt.Println(err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": tokenDecrypt,
		})
	})

	c4f := e.Group("code4func")
	c4f.Use(JWTMiddleware())

	c4f.GET("/profile", func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)

		claims := token.Claims.(*model.JwtCustomClaims)

		return c.JSON(http.StatusOK, echo.Map{
			"userId": claims.UserId,
			"fullName": claims.FullName,
		})
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(jwtKey),
		BeforeFunc: func(c echo.Context) {
			// Authorization: Bearer xxxx
			authorization := c.Request().Header.Get("Authorization")
			parts := strings.Split(authorization, " ")
			tokenDecrypt, _ := decrypt([]byte(KEY), parts[1])

			fmt.Println("**************************")
			fmt.Println(tokenDecrypt)
			fmt.Println("**************************")

			c.Request().Header.Set("Authorization", "Bearer " + tokenDecrypt)
		},
	}

	return middleware.JWTWithConfig(config)
}
