package services

import (
	"errors"
	"fmt"
)

// AuthService предоставляет методы для аутентификации пользователей.
type AuthService struct {
}

// NewAuthService создает новый экземпляр AuthService.
//
// Возвращает указатель на AuthService.
//
// Пример использования:
//
//	authService := NewAuthService()
func NewAuthService() *AuthService {
	return &AuthService{}
}

// Auth проверяет аутентификацию пользователя.
//
// Принимает имя пользователя и пароль в качестве параметров.
//
// Возвращает ошибку, если имя пользователя или пароль неверны.
//
// Пример использования:
//
//	err := authService.Auth("root", "root")
//	if err != nil {
//		fmt.Println(err)
//	}
func (as *AuthService) Auth(username, password string) error {
	if username != "admin" && password != "admin" {
		return errors.New("неверный пользователь")
	}

	return nil
}

// Register регистрирует нового пользователя
//
// Принимает имя пользователя и пароль в качестве параметров.
//
// Возвращает ошибку, если имя пользователя уже существует.
func (as *AuthService) Register(username, password string) error {
	if username == "admin" {
		return errors.New("невозможно создать admin")
	}

	fmt.Printf("Создание пользователя %s:%s", username, password)
	return nil
}
