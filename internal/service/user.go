package service

import (
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
	return model.User{}, nil
}