package main

import (
	"fmt"

	"github.com/alfredoprograma/filez-server/internal/app"
	"github.com/alfredoprograma/filez-server/internal/config"
	"github.com/alfredoprograma/filez-server/internal/database"
	"github.com/alfredoprograma/filez-server/internal/routes"
)

func main() {
	config := config.NewConfig()
	db := database.Connect(config)

	app := app.NewApplication(config, db)
	routes.LoadRoutes(&app)

	if err := app.Server.Listen(fmt.Sprintf(":%d", config.API.Port)); err != nil {
		panic(err)
	}
}
