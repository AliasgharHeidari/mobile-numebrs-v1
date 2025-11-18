package server

import (
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/api/middleware"
)
func Start() {
	app:= fiber.New()

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
	app.Delete("/user/:id/mobilenumber", middleware.JwtProtected(), handler.DeleteMobileNumber)

	//Listen port
	app.Listen(":9898")

}