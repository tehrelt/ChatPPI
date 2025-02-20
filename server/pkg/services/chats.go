package services

import "fmt"

// ChatService предоставляет методы для работы с чатом.
type ChatService struct {
}

// NewChatService создает новый экземпляр ChatService.
func NewChatService() *ChatService {
	return &ChatService{}
}

// Send отправляет сообщение от пользователя.
// Параметры:
//   - from: идентификатор отправителя.
//   - msg: текст сообщения.
// Возвращает:
//   - error: ошибка, если произошла ошибка при отправке сообщения.
func (cs *ChatService) Send(from, msg string) error {
	fmt.Printf("%s: %s\n", from, msg)
	return nil
}

// CreateConversation создает новую беседу для пользователя.
// Параметры:
//   - userId: идентификатор пользователя.
// Возвращает:
//   - error: ошибка, если произошла ошибка при создании беседы.
func (cs *ChatService) CreateConversation(userId string) error {
	fmt.Printf("%s created conversation\n", userId)
	return nil
}

// GetMessages возвращает все сообщения из беседы для пользователя.
// Параметры:
//   - userId: идентификатор пользователя.
//   - conversationId: идентификатор беседы.
// Возвращает:
//   - error: ошибка, если произошла ошибка при получении сообщений.
func (cs *ChatService) GetMessages(userId, conversationId string) error {
	fmt.Printf("%s read all messages from %s\n", userId, conversationId)
	return nil
}
