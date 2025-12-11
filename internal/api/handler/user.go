package handler

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/service"
	"github.com/gofiber/fiber/v2"
)

// @Summary 		Get all users
// @Description 	Get all users from database and return
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param  			page  query int false "page number, default value: 1"
// @Param  			limit query int false "item per page, default value: 5"
// @Success 		200 {object} []model.User
// @Failure 		500 {object} model.StatusInternalServerErrorResponse
// @Failure 		401 {object} model.StatusUnauthorizedResponse
// @Security 		BearerAuth
// @Router 			/user [get]
func GetUserList(c *fiber.Ctx) error {
	var (
		start int = 0
		end   int = 0
	)

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid value for page. valid values are: -1: disable paging, 0<page: enable paging",
		})
	}
	// TODO: handle raised error
	limit, err := strconv.Atoi(c.Query("limit", "5"))
	if err != nil || limit == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid value for limit.",
		})
	}

	if page > 0 && limit > 0 {
		start = (page - 1) * limit
		end = start + limit - 1
	}

	UserList, err := service.GetUserList(start, end)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get user list",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"page":  page,
		"limit": limit,
		"users": UserList,
	})
}

// @Summary 		Get user by ID
// @Description 	Get user by ID from database and return
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "User ID"
// @Success 		200 {object} model.User
// @Failure 		400 {object} model.StatusBadRequestResponse
// @Failure 		404 {object} model.StatusNotFoundResponse
// @Failure 		401 {object} model.StatusUnauthorizedResponse
// @Failure 		500 {object} model.StatusInternalServerErrorResponse
// @Router 			/user/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user ID",
		})
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// @Summary 		Create user
// @Description 	Create user and save it to database
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			user body model.CreateUserRequest true "User object"
// @Success 		201 {object} model.CreateUserSuccessResponse
// @Failure 		400 {object} model.StatusBadRequestResponse
// @Failure 		401 {object} model.StatusUnauthorizedResponse
// @Failure 		500 {object} model.StatusInternalServerErrorResponse
// @Router 		/user [post]
func CreateUser(c *fiber.Ctx) error {
	var NewUser model.User

	if err := c.BodyParser(&NewUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	id, err := service.CreateUser(NewUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
		"user_id": id,
	})
}

// @Summary 		Update user by ID
// @Description 	Update user by ID and save it to database
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "User ID"
// @Param 			user body model.CreateUserRequest true "User object"
// @Success 		200 {object} model.UpdateUserSuccessResponse
// @Failure 	    400 {object} model.StatusBadRequestResponse
// @Failure 		404 {object} model.StatusNotFoundResponse
// @Failure 		401 {object} model.StatusUnauthorizedResponse
// @Failure 		500 {object} model.StatusInternalServerErrorResponse
// @Router 			/user/{id} [put]
func UpdateUserByID(c *fiber.Ctx) error {
	UserID := c.Params("id")
	id, err := strconv.Atoi(UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var UpdatedUser model.User

	if err := c.BodyParser(&UpdatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := service.UpdateUserByID(id, UpdatedUser); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User has Been updated successfully",
	})
}

// @Summary 		Delete user by ID
// @Description  	Delete user from database by ID
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "User ID"
// @Success 		200 {object} model.DeleteUserSuccessResponse
// @Failure 		404 {object} model.StatusNotFoundResponse
// @Failure 		401 {object} model.StatusUnauthorizedResponse
// @Failure 		400 {object} model.StatusBadRequestResponse
// @Failure 		500 {object} model.StatusInternalServerErrorResponse
// @Router 			/user/{id} [delete]
func DeleteUserByID(c *fiber.Ctx) error {
	UserID := c.Params("id")
	id, err := strconv.Atoi(UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	if err := service.DeleteUserByID(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user deleted successfully",
	})
}

// @Summary 		Add mobile number to user by ID
// @Description  	Add mobile number to user by ID and save it to database
// @Tags 			mobile-numbers
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "User ID"
// @Param 			mobileNumber body model.MobileNumber true "Mobile number object"
// @Success 		200 {object} model.AddMobileNumberSuccessResponse
// @Failure 		400 {object} model.StatusBadRequestResponse
// @Failure 		404 {object} model.StatusNotFoundResponse
// @Failure 		401 {object} model.StatusUnauthorizedResponse
// @Failure 		500 {object} model.StatusInternalServerErrorResponse
// @Router 			/user/{id}/mobilenumber [post]
func AddMobileNumber(c *fiber.Ctx) error {
	UserID := c.Params("id")
	id, err := strconv.Atoi(UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var mobileNumber model.MobileNumber
	if err := c.BodyParser(&mobileNumber); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := service.AddMobileNumber(id, mobileNumber); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Mobile number has been added successfully",
	})
}

// @Summary 		Delete mobile number by ID
// @Description  	Delete mobile number from database by ID
// @Tags			mobile-numbers
// @Accept 			json
// @Produce 		json
// @Param 			id path int true "Mobile number ID"
// @Success 		200 {object} model.DeleteMobileNumberSuccessResponse
// @Failure 		400 {object} model.StatusBadRequestResponse
// @Failure 		404 {object} model.StatusNotFoundResponse
// @Failure 		401 {object} model.StatusUnauthorizedResponse
// @Failure 		500 {object} model.StatusInternalServerErrorResponse
// @Router			/user/{id}/mobilenumber [delete]
func DeleteMobileNumber(c *fiber.Ctx) error {
	UserID := c.Params("id")
	id, err := strconv.Atoi(UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	number := c.Params("number")

	if err := service.DeleteMobileNumber(id, number); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user Not Found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Mobile number has been deleted successfully",
	})

}

// @Summery 	 upload file
// @Description  upload file to server
// @Tags 		 Upload
// @Accept 		 multipart/formdata
// @Produce 	 json
// @Param        file formData file true "upload file"
// @Router 		 /profile/upload [post]
func UploadImage(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))

	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "only jpg & png allowed",
		})
	}
	path := "./uploads" + "/" + file.Filename

	if err := c.SaveFile(file, path); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save file",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "photo has been uploaded successfuly",
		"path":    path,
	})
}
