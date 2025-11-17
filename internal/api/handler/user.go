package handler

import (
	"strconv"

	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
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


	func CreateUser(c *fiber.Ctx) error {
		var NewUser model.User

		if err := c.BodyParser(&NewUser) ; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "Invalid request body",
			})
		}

		id, err := service.CreateUser(NewUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error" : "Failed to create user",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message" : "user created successfully",
			"user_id" : id,
		})
	}

	func UpdateUserByID(c *fiber.Ctx) error {
		UserID := c.Params("id")
		id, err := strconv.Atoi(UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "Invalid user ID",
			})
		}

	var UpdatedUser model.User

	if err := c.BodyParser(&UpdatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : "Invalid request body",
		})
	}

	if err := service.UpdateUserByID(id, UpdatedUser); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "User has Been updated successfully",
	})
	}


	func DeleteUserByID(c *fiber.Ctx) error {
		UserID:= c.Params("id")
		id, err := strconv.Atoi(UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "Invalid user ID",
			})
		}
			if err := service.DeleteUserByID(id); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error" : "user not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message" : "user deleted successfully",
		})
	}

	func AddMobileNumber(c *fiber.Ctx) error {
		UserID := c.Params("id")
		id, err := strconv.Atoi(UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "Invalid user ID",
			})
		}

		var mobileNumber model.MobileNumber
		if err := c.BodyParser(&mobileNumber); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "Invalid request body",
			})
		}
 
		if err := service.AddMobileNumber(id, mobileNumber); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error" : "user not found",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message" : "Mobile number has been added successfully",
		})
	}

	func DeleteMobileNumber(c *fiber.Ctx) error {
		UserID := c.Params("id")
		id, err := strconv.Atoi(UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error" : "Invalid user ID",
			})
		}

		if err := service.DeleteMobileNumber(id); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error" : "user Not Found",
 			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"Message" : "Mobile number has been deleted successfully",
		})

	}