package main

import (
	"context"
	"log"

	"censor.services/bootstrap"
)

// @title censor.services
// @version 2.0
// @description censor.services API接口文档

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8085
// @BasePath /api/v1/censor
// @schemes http
func main() {
	app := bootstrap.App()
	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err)
	}
}
