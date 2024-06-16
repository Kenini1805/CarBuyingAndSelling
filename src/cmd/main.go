package cmd

import (
	"CarBuyingAndSelling/src/api"
	"CarBuyingAndSelling/src/config"
	"CarBuyingAndSelling/src/infra/cache"
	database "CarBuyingAndSelling/src/infra/persistence/database"
	"CarBuyingAndSelling/src/infra/persistence/migration"
	"CarBuyingAndSelling/src/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up1()

	api.InitServer(cfg)
}
