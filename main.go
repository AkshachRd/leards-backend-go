package main

import (
	"github.com/AkshachRd/leards-backend-go/docs"
	"github.com/AkshachRd/leards-backend-go/handlers"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
)

// @title           Leards Backend API
// @version         1.0
// @description     This is a leards language learning app api.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := SetupRouter()

	r.Run(":8080")
}

func DbInit() *gorm.DB {
	db, err := models.Setup()
	if err != nil {
		log.Println("Problem setting up database")
	}
	return db
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := DbInit()

	server := handlers.NewServer(db)
	docs.SwaggerInfo.Title = "Leards Backend API"

	v1 := r.Group("/api/v1")
	authorizedV1 := v1.Group("/", server.AuthService())
	{
		accounts := v1.Group("/accounts")
		{
			accounts.POST("", server.CreateUser)
		}
		auth := authorizedV1.Group("auth")
		{
			auth.GET(":id", server.RefreshToken)
			auth.DELETE(":id", server.RevokeToken)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
