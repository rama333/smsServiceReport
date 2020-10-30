package services

import "smsServiceReport/internal/restapi/userMessages/models"

type ServiceUserMessages struct {
	su IUserMessages
}

type IUserMessages interface {
	GetUserMessages(dest_add string, startDuration string, endDuration string) ([]models.Messages, error)
}

func New(iU IUserMessages) *ServiceUserMessages {

	return &ServiceUserMessages{su: iU}
}

func (ser ServiceUserMessages) GetUserMessages(dest_add string, startDuration string, endDuration string) ([]models.Messages, error) {

	return ser.su.GetUserMessages(dest_add, startDuration, endDuration)
}
