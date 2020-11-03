package services

import (
	models2 "smsServiceReport/internal/models"
)

type Service struct {
	fg FinderAndGetMessages
}

type FinderAndGetMessages interface {
	GetMessages(startDuration string, endDuration string) ([]models2.Messages, error)
}

func NewServiceMessages(messages FinderAndGetMessages) *Service {

	return &Service{fg: messages}
}

func (s *Service) GetMessages(startDuration string, endDuration string) ([]models2.Messages, error) {

	return s.fg.GetMessages(startDuration, endDuration)
}
