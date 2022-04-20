package main

import (
	"os"

	"github.com/bancodobrasil/featws-api/config"
	"github.com/bancodobrasil/featws-api/database"
	_ "github.com/bancodobrasil/featws-api/docs"
	"github.com/bancodobrasil/featws-api/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func setupLog() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)
}

// @title FeatWS API
// @version 1.0
// @description API Project to provide operations to manage FeatWS knowledge repositories rules
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9007
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

// Run start the resolver server with resolverFunc
func main() {

	setupLog()

	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Não foi possível carregar as configurações: %s\n", err)
		return
	}

	cfg := config.GetConfig()
	if cfg == nil {
		log.Fatalf("Não foi carregada configuracão!\n")
		return
	}

	database.ConnectDB()

	router := gin.New()

	routes.SetupRoutes(router)

	port := cfg.Port

	router.Run(":" + port)

	log.Infof("Listen on http://0.0.0.0:%s\n", port)
}
