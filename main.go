package main

import (
	"log"

	"github.com/dimasardnt6/dimasardnt6-ulbi/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/dimasardnt6/dimasardnt6-ulbi/url"

	"github.com/gofiber/fiber/v2"

	_ "github.com/dimasardnt6/dimasardnt6-ulbi/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host  https://dimasardnt6-ulbi.herokuapp.com
// @BasePath /

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
