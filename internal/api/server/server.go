package server

import (
	_ "github.com/AliasgharHeidari/mobile-numbers-v1/docs/api"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/api/handler"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/api/middleware"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)
func Start() {
	app:= fiber.New()

	//Load config
	cfg, err:= config.LoadConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}

  	//CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.Cors.AllowedOrigins,
		AllowMethods: cfg.Cors.AllowedMethods,
		AllowHeaders: cfg.Cors.AllowedHeaders,
	}))

	//Logger
	app.Use(logger.New())

	//Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	//User Login
	app.Post("/user/login", handler.Login)
	
	//User routes
	app.Get("/user", middleware.JwtProtected(), handler.GetUserList)
	app.Get("/user/:id", middleware.JwtProtected(), handler.GetUserByID)
	app.Post("/user", middleware.JwtProtected(), handler.CreateUser)
	app.Put("/user/:id", middleware.JwtProtected(), handler.UpdateUserByID)
	app.Delete("/user/:id", middleware.JwtProtected(), handler.DeleteUserByID)

	//Mobile Number routes
	app.Post("/user/:id/mobilenumber", middleware.JwtProtected(), handler.AddMobileNumber)
	app.Delete("/user/:id/mobilenumber/:number", middleware.JwtProtected(), handler.DeleteMobileNumber)

	//Listen port
	app.Listen(cfg.Server.Host + ":" + cfg.Server.Port)
}