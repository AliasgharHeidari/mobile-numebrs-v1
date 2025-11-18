package service

import (
	"fmt"

	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/model"
	onmemory "github.com/AliasgharHeidari/mobile-numbers-v1/internal/repository/on-memory"
)

func GetUserList() ([]model.User, error) {
	return onmemory.Users, nil
}

func GetUserByID(id int) (model.User, error) {
	for i := range onmemory.Users {
		if onmemory.Users[i].ID == id {
			return onmemory.Users[i], nil
		}
	}
	return model.User{}, fmt.Errorf("user not found")
}

func CreateUser(user model.User) (int, error) {
	user.ID = len(onmemory.Users)+ 1

	onmemory.Users = append(onmemory.Users, user)
	return user.ID, nil
}

func UpdateUserByID(id int, UpdatedUser model.User) error {
	for i := range onmemory.Users {
		if onmemory.Users[i].ID == id {
		onmemory.Users[i] = UpdatedUser		
		onmemory.Users[i].ID = id
		return nil
		}
	}
	return fmt.Errorf("user not found")
}

func DeleteUserByID(id int) error{
	for i := range onmemory.Users {
		if onmemory.Users[i].ID == id {
			onmemory.Users = append(onmemory.Users[:i], onmemory.Users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func AddMobileNumber(id int, mobileNumber model.MobileNumber) error {
	for i := range onmemory.Users{
		if onmemory.Users[i].ID == id{
			onmemory.Users[i].MobileNumbers = append(onmemory.Users[i].MobileNumbers, mobileNumber)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func DeleteMobileNumber(id int) error {
	for i := range onmemory.Users{
		if onmemory.Users[i].ID == id {
			onmemory.Users[i].MobileNumbers = nil
			return nil
		}
	}
	return fmt.Errorf("user not found")
}