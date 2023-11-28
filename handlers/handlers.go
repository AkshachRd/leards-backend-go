package handlers

import (
	"github.com/AkshachRd/leards-backend-go/docs"
	v1 "github.com/AkshachRd/leards-backend-go/handlers/api/v1"
	"github.com/AkshachRd/leards-backend-go/middlewares"
	"github.com/AkshachRd/leards-backend-go/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Allow requests from any origin
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	docs.SwaggerInfo.Title = "Leards Backend API"

	apiv1 := r.Group("/api/v1")

	basicAuthorizedV1 := apiv1.Group("", middlewares.AuthService(middlewares.BasicAuth))
	bearerAuthorizedV1 := apiv1.Group("", middlewares.AuthService(middlewares.BearerAuth))
	{

		accounts := apiv1.Group("/accounts")
		accountsBasicAuthorized := basicAuthorizedV1.Group("/accounts")
		accountsBearerAuthorized := bearerAuthorizedV1.Group("/accounts")
		{
			accounts.POST("", v1.CreateUser)
			accountsBearerAuthorized.PUT(":user_id", v1.UpdateUser)
			accountsBasicAuthorized.GET("", v1.LoginUser)
			accounts.GET(":user_id/avatar", v1.GetAvatar)
			accounts.Static("/avatars", settings.AppSettings.EnvVars.AvatarBasePath)
			accountsBearerAuthorized.PUT(":user_id/avatar", v1.UploadAvatar)

			userSettings := accountsBearerAuthorized.Group(":user_id/settings")
			{
				userSettings.GET("", v1.GetUserSettings)
				userSettings.PUT("", v1.UpdateUserSettings)
			}
		}
		auth := bearerAuthorizedV1.Group("/auth")
		{
			auth.GET(":user_id", v1.RefreshToken)
			auth.DELETE(":user_id", v1.RevokeToken)
		}
		folders := bearerAuthorizedV1.Group("/folders")
		foldersWithId := folders.Group(":folder_id")
		{
			foldersWithId.GET("", v1.GetFolder)
			folders.POST("", v1.CreateFolder)
			foldersWithId.PUT("", v1.UpdateFolder)
			foldersWithId.DELETE("", v1.DeleteFolder)
		}
		decks := bearerAuthorizedV1.Group("/decks")
		decksWithId := decks.Group(":deck_id")
		{
			decks.POST("", v1.CreateDeck)
			decksWithId.GET("", v1.GetDeck)
			decksWithId.PUT("", v1.UpdateDeck)
			decksWithId.DELETE("", v1.DeleteDeck)
		}
		cards := bearerAuthorizedV1.Group("/cards")
		{
			cards.GET(":storage_type/:storage_id", v1.GetStorageCards)
			cards.PUT("deck/:deck_id", v1.SyncCards)
		}
		library := bearerAuthorizedV1.Group("/library")
		{
			library.GET(":user_id", v1.GetFavoriteStorages)
			library.POST(":user_id/:storage_type/:storage_id", v1.AddStorageToFavorite)
			library.DELETE(":user_id/:storage_type/:storage_id", v1.RemoveStorageFromFavorite)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
