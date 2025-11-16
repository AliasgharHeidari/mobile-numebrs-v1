package handler

import (
	"strconv"

	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/service"
	"github.com/gofiber/fiber/v2"
)


func GetUserList(c *fiber.Ctx) error {
	userslist, err := service.GetUserList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : "failed to get user list",
		})
	}
	return c.Status(fiber.StatusOK).JSON(userslist)
}

func GetUserByID(c *fiber.Ctx) error{
	userID := c.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "invalid user ID",
		})
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "user not found",
		})
	}
	return  c.Status(fiber.StatusOK).JSON(user)
	}