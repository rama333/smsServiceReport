package services

import (
	"smsServiceReport/internal/restapi/messages/models"
	"time"
)

type Service struct {
	fg FinderAndGetMessages
}

type FinderAndGetMessages interface {
	GetMessages(startDuration time.Time, endDuration time.Time) ([]models.Messages, error)
}

func NewService(messages FinderAndGetMessages) *Service {

	return &Service{fg: messages}
}

func (s *Service) GetMessages(startDuration time.Time, endDuration time.Time) ([]models.Messages, error) {

	return s.fg.GetMessages(startDuration, endDuration)
}
