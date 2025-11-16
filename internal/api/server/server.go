package server

import (
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/api/handler"
	"github.com/gofiber/fiber/v2"
)
func Start() {
	app:= fiber.New()

	//User Login
	/*
	app.Post("/user/login", handler.Login)
*/
	//User routes
	app.Get("/user", handler.GetUserList)
	app.Get("/user/:id", handler.GetUserByID)
	/*
	app.Post("/user/:id/mobilenumber", handler.AddMobileNumber)
	app.Put("/user/:id", handler.UpdateUserByID)
	app.Delete("/user/:id/mobilenumber/:number", handler.DeleteMobileNumber)

	//Mobile Number routes
	app.Post("/user/:id", handler.AddMobileNumber)
	app.Delete("/user/:id", handler.DeleteMobileNumber)
*/
	//Listen port
	app.Listen(":9898")

}