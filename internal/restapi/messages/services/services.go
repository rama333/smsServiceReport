package services

import (
	"smsServiceReport/internal/restapi/messages/models"
)

type Service struct {
	fg FinderAndGetMessages
}

type FinderAndGetMessages interface {
	GetMessages(startDuration string, endDuration string) ([]models.Messages, error)
}

func NewService(messages FinderAndGetMessages) *Service {

	return &Service{fg: messages}
}

func (s *Service) GetMessages(startDuration string, endDuration string) ([]models.Messages, error) {

	return s.fg.GetMessages(startDuration, endDuration)
}
