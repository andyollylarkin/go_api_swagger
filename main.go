package main

import (
	_ "github.com/andyollylarkin/go_api_swagger/docs" // импортирует template из docs.go
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// https://github.com/swaggo/swag документация

// @title Swagger Example API
// @version 1.0
// @description Example of API with swaggo
// @contact.email follenlast2@gmail.com
// @host 127.0.0.1:8080
// @scheme http
// @BasePath /api/v1
// @securitydefinitions.apikey ApiKey -- ApiKey используется в аннотации @Secutiry в защищенных роутах
// @in header -- говорит о том, что инфа об авторизации должна передаваться в заголовке
// @name Authorization -- название заголовка
// @authorizationurl http://127.0.0.1;8080/api/v1/auth
func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", Ping)
		v1.GET("/name", Name)
		v1.GET("/auth", Auth)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("127.0.0.1:8080")
}
