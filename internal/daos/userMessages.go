package daos

import (
	"fmt"
	"smsServiceReport/internal/config"
	"smsServiceReport/internal/restapi/userMessages/models"
)

type UserMessagesDAO struct {
}

func NewUserMessages() *UserMessagesDAO {

	return &UserMessagesDAO{}
}

func (userMes UserMessagesDAO) GetUserMessages(dest_add string, startDuration string, endDuration string) ([]models.Messages, error) {

	var mes []models.Messages

	q := fmt.Sprintf("SELECT Send.date,submit_date,done_date,Receive.destination_addr,id,sms_text, Send.source_addr, SentMesId.message_id, stat FROM SentMesId INNER JOIN Receive ON Receive.message_id = SentMesId.message_id INNER JOIN Send ON SentMesId.sequence = Send.sequence and Receive.destination_addr = Send.dest_addr  where Send.date BETWEEN toDateTime('%s') and toDateTime('%s') and destination_addr = %s;", startDuration, endDuration, dest_add)
	err := config.Config.DB.Select(&mes, q)

	if err != nil {
		return nil, err
	}

	return mes, nil
}
