package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type input int

// Ping
// @Summary Ping
// @Tags ping
// @Security ApiKey
// @Produce json
// @Param input path int true "User input"
// @Success 200 {object} User
// @Failure 404
// @Failure 401
// @Router /ping [get]
func Ping(context *gin.Context) {
	h := context.GetHeader("Authorization")
	fmt.Println(h)
	if h != "AndyLarkin" {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var inp input
	fmt.Println(inp)

	user := User{Name: "Anna", Age: 16}
	context.JSON(http.StatusOK, user)
}

// Name
// @Summary Name
// @Tags ping -- группа роутов по тегам
// @Produce json -- тип отдаваемой информации
// @Param input path int true "User input"
// @Success 200 {object} User
// @Router /name [get]
func Name(context *gin.Context) {
	var inp input
	fmt.Println(inp)

	user := User{Name: "Name", Age: 404}
	context.JSON(http.StatusOK, user)
}

func Auth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"auth-key": "AndyLarkin"})
}
