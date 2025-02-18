package main

import (
	"chat-ppi-server/pkg/services"
	"fmt"
)

// main Точка входа
func main() {
	authsvc := services.NewAuthService()

	err := authsvc.Auth("admin", "admin")
	if err != nil {
		panic(err)
	}
	fmt.Println("Авторизация успешна")

	authsvc.Register("pi-21a", "donntu")
}
