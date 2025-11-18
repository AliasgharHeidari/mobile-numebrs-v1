package middleware

 import (
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
 )

 func JwtProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if  authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error" : "Unauthorized",
			})
		}

		tokenString := authHeader[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}

			return []byte(utils.SecretKey), nil
		})
	
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "Invalid or expired token",

			})
		}

		claims := token.Claims.(jwt.MapClaims)
		
		c.Locals("userName", claims["userName"])

		return c.Next()
	}




}