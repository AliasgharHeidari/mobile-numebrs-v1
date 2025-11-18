package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/api/utils"
)

	type LoginRequest struct {
		InputUserName string `json:"userName"`
		InputPassword string `json:"password"`
	}

	type Creadentials struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}

func Login(c *fiber.Ctx) error {
	
	var UserLogin LoginRequest
	var Creadentials Creadentials

	Creadentials.UserName = "Aliasghar"
	Creadentials.Password = "1234"

	if err := c.BodyParser(&UserLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Invalid request body",
		})
	}

	if UserLogin.InputUserName == "" || UserLogin.InputPassword == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "Username & password required",
		})
	} 


	if UserLogin.InputUserName != Creadentials.UserName || UserLogin.InputPassword != Creadentials.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "Invalid username or password",
		})
	}

	var token string
	var err error
	
	if token, err = utils.GenerateToken(Creadentials.UserName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"token" : token,
	})
}