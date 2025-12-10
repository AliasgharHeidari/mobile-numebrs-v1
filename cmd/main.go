package main

//@title Mobile Numbers API
//@version 1.0
//@description This is a sample server for Mobile Numbers API
//@termsOfService http://swagger.io/terms/
//@host 127.0.0.1:9898
//BasePath /

//@contact.name Aliasghar Heidari
//@contact.url https://github.com/AliasgharHeidari

//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html

//@securityDefinitions.apikey BearerAuth
//@in header
//@name Authorization
//@description Type "Bearer" followed by a space and JWT token

//Apply security globally
//@Security BearerAuth

import (
	"fmt"

	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/api/server"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/database"
	onmemory "github.com/AliasgharHeidari/mobile-numbers-v1/internal/repository/on-memory"
	dataonredis "github.com/AliasgharHeidari/mobile-numbers-v1/internal/repository/redis"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/seeder"
)

func main() {
	fmt.Println("Starting the server...")
	database.ConnectDB()
	database.AutoMigrate()
	seeder.SeedUser()
	dataonredis.InitRedisClient()
	onmemory.InitUsers()
	server.Start()
}