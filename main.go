package main

import (
	"log"

	"github.com/Fatwaff/ws-fatwa/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/Fatwaff/ws-fatwa/url"

	// swagger handler
	_ "github.com/Fatwaff/ws-fatwa/docs"
	"github.com/gofiber/fiber/v2"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server

// @contact.name API Support
// @contact.url https://github.com/Fatwaff
// @contact.email 1214038@std.ulbi.ac.id

// @host ws-fatwa.herokuapp.com
// @BasePath /
// @schemes http https

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
