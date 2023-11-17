package main

import (
	"github.com/AkshachRd/leards-backend-go/handlers"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/AkshachRd/leards-backend-go/settings"
)

func init() {
	models.Setup()
	settings.Setup()
}

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
	r := handlers.SetupRouters()

	r.Run(":8080")
}
