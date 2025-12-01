package onmemory

import (
	"log"
	"fmt"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
	dataonredis "github.com/AliasgharHeidari/mobile-numbers-v1/internal/repository/redis"
)

func InitUsers() {

	var initUsers []model.User

	for i := 1; i <= 100; i++ {
	newUser := model.User{
			ID: i,
			Name: fmt.Sprintf("Ali-%d", i),
			FamilyName: fmt.Sprintf("Heidari-%d", i),
			Age: 18,
			IsMarried: false,
		}
		initUsers = append(initUsers, newUser)

	}

	for _, user := range initUsers {
		err := dataonredis.SaveUserToRedis(user)
		if err != nil {
			log.Printf("failed to save user %s to redis, error: %+v", user.Name, err)
		}
	}
}
