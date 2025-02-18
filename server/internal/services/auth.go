package services

import "errors"

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

	return errors.New("неверный пользователь")
}
