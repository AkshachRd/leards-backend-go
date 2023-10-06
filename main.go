package main

import (
	"github.com/AkshachRd/leards-backend-go/dbSetup"
	"github.com/AkshachRd/leards-backend-go/docs"
	"github.com/AkshachRd/leards-backend-go/handlers"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-contrib/cors"
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
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description ATTENTION! HOW TO USE: Type "Bearer" followed by a space and a token. Example: "Bearer \<token\>".

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := SetupRouter()

	r.Run(":8080")
}

func MockData(db *gorm.DB) error {
	_, err := models.NewUser(db, "Owner", "owner", "12345Q")
	if err != nil {
		return err
	}

	accessType := models.AccessType{Type: string(models.Public)}
	db.First(&accessType)

	models.NewFolder(db, "folder1", accessType.ID)

	return nil
}

func DbInit() *gorm.DB {
	db, err := dbSetup.Setup()
	if err != nil {
		log.Println("Problem setting up database", err)
	}

	acc := models.AccessType{Type: string(models.Public)}
	db.Create(&acc)
	acc = models.AccessType{Type: string(models.Private)}
	db.Create(&acc)

	err = MockData(db)
	if err != nil {
		log.Println("Problem setting up database", err)
	}

	return db
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Allow requests from any origin
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	db := DbInit()

	server := handlers.NewServer(db)
	docs.SwaggerInfo.Title = "Leards Backend API"

	v1 := r.Group("/api/v1")
	basicAuthorizedV1 := v1.Group("", server.AuthService(handlers.BasicAuth))
	bearerAuthorizedV1 := v1.Group("", server.AuthService(handlers.BearerAuth))
	{
		accounts := v1.Group("/accounts")
		accountsAuthorized := basicAuthorizedV1.Group("/accounts")
		{
			accounts.POST("", server.CreateUser)
			accountsAuthorized.GET("", server.LoginUser)
		}
		auth := bearerAuthorizedV1.Group("/auth")
		{
			auth.GET(":id", server.RefreshToken)
			auth.DELETE(":id", server.RevokeToken)
		}
		folders := bearerAuthorizedV1.Group("/folders")
		{
			folders.GET(":id", server.GetSingleFolder)
		}
		decks := bearerAuthorizedV1.Group("/decks")
		{
			decks.GET(":id", server.GetSingleDeck)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
