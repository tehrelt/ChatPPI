package ai

import (
	"fmt"
	"strings"
)

// AiGate представляет собой структуру, которая содержит конечную точку для вызова AI.
type AiGate struct {
	endpoint string
}

// NewAiGate создает новый экземпляр AiGate с заданной конечной точкой.
func NewAiGate(endpoint string) *AiGate {
	return &AiGate{endpoint: endpoint}
}

// Invoke вызывает AI с заданным сообщением и возвращает ответ AI.
func (*AiGate) Invoke(message string) string {
	fmt.Printf("invoke %s\n", message)
	return fmt.Sprintf("ai: %s", strings.ToUpper(message))
}
