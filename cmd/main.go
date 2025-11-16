package main

import (
	"fmt"
	"github.com/AliasgharHeidari/mobile-numbers-v1/internal/api/server"
	onmemory "github.com/AliasgharHeidari/mobile-numbers-v1/internal/repository/on-memory"
)

func main() {
	fmt.Println("Starting the server...")
	onmemory.InitUsers()
	server.Start()
}