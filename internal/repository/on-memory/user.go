package onmemory

import (
	"fmt"
	"log"

	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/database"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
	dataonredis "github.com/AliasgharHeidari/mobile-numbers-v1/internal/repository/redis"
)

func InitUsers() {

	var initUsers []model.User

	for i := 1; i <= 10; i++ {
		newUser := model.User{
			ID:         i,
			Name:       fmt.Sprintf("Ali-%d", i),
			FamilyName: fmt.Sprintf("Heidari-%d", i),
			Age:        18,
			IsMarried:  false,
		}
		initUsers = append(initUsers, newUser)

	}
	DB := database.GetDB()

	for i := range initUsers {

		user := initUsers[i]

		if err := DB.Create(&initUsers[i]).Error; err != nil {
			log.Printf("failed to save user %s to database, error : %+v", user.Name, err)
		}
		err := dataonredis.SaveUserToRedis(user)
		if err != nil {
			log.Printf("failed to save user %s to redis, error: %+v", user.Name, err)
		}
	}
}
