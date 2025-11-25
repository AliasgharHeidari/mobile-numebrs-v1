package service

import (
	"fmt"
	"math/rand"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
	dataonredis "github.com/AliasgharHeidari/mobile-numbers-v1/internal/repository/redis"
)

func GetUserList() ([]model.User, error) {
	return dataonredis.GetAllUsersFromRedis()
}

func GetUserByID(id int) (model.User, error) {
	user, err := dataonredis.LoadUserFromRedis(id)
	if err != nil {
		return model.User{}, err
	}
	if user == nil {
		return model.User{}, fmt.Errorf("user not found")
	}
	return  *user, nil
}

func CreateUser(user model.User) (int, error) {
	// generate randon int as id
	user.ID = rand.Intn(10000000)
	err := dataonredis.SaveUserToRedis(user)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func UpdateUserByID(id int, UpdatedUser model.User) error {
	UpdatedUser.ID = id

	err := dataonredis.SaveUserToRedis(UpdatedUser)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserByID(id int) error{
	err := dataonredis.DeleteUserFromRedis(id)
	if err != nil {
		return err
	}

	return nil
}

func AddMobileNumber(id int, mobileNumber model.MobileNumber) error {
	loadedUser, err := dataonredis.LoadUserFromRedis(id)
	if err != nil {
		return err
	}

	if loadedUser != nil {
		loadedUser.MobileNumbers = append(loadedUser.MobileNumbers, mobileNumber)
		err = dataonredis.SaveUserToRedis(*loadedUser)
		if err != nil {
			return err
		}
		return nil
	
	}
	
	return fmt.Errorf("user not found")
}

func DeleteMobileNumber(id int, Number string) error {
	loadedUser, err := dataonredis.LoadUserFromRedis(id)
	if err != nil {
		return err
	}

	if loadedUser != nil {
		var updatedMobileNumbers []model.MobileNumber
		for _, mobileNumber := range loadedUser.MobileNumbers {
			if mobileNumber.Number != Number {
				updatedMobileNumbers = append(updatedMobileNumbers, mobileNumber)
			}
		}
		loadedUser.MobileNumbers = updatedMobileNumbers
		err = dataonredis.SaveUserToRedis(*loadedUser)
		if err != nil {
			return err
		}
	}

	return fmt.Errorf("user not found")
}