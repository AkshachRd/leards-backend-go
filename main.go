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
	user, err := models.NewUser(db, "Owner", "owner", "12345Q")
	if err != nil {
		return err
	}

	folder, err := models.NewFolder(db, "folder1", models.Private)
	if err != nil {
		return err
	}

	folder.ParentFolderID = &user.RootFolderID
	err = db.Save(&folder).Error
	if err != nil {
		return err
	}

	deck, err := models.NewDeck(db, "deck1", models.Private, folder.ID)
	if err != nil {
		return err
	}

	deck.Cards = []models.Card{
		{DeckID: deck.ID, FrontSide: "Apple", BackSide: "Яблокоф"},
		{DeckID: deck.ID, FrontSide: "Banana", BackSide: "Бананоф"},
	}
	err = db.Save(&deck).Error
	if err != nil {
		return err
	}

	return nil
}

func DbInit() *gorm.DB {
	db, err := dbSetup.Setup()
	if err != nil {
		log.Println("Problem setting up database", err)
	}

	models.FillAccessTypes(db)

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
			auth.GET(":user_id", server.RefreshToken)
			auth.DELETE(":user_id", server.RevokeToken)
		}
		foldersWithId := bearerAuthorizedV1.Group("/folders/:folder_id")
		{
			foldersWithId.GET("", server.GetSingleFolder)

			decks := foldersWithId.Group("/decks")
			decksWithId := decks.Group(":deck_id")
			{
				decksWithId.GET("", server.GetDeck)
				decks.POST("", server.CreateDeck)
				decksWithId.PUT("", server.UpdateDeck)

				cards := decksWithId.Group("/cards")
				{
					cards.GET("", server.GetCards)
					cards.PUT("", server.SyncCards)
				}
			}
		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
