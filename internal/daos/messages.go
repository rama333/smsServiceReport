package daos

import (
	"fmt"
	"log"
	"smsServiceReport/internal/config"
	models2 "smsServiceReport/internal/models"
)

type MessagesDAO struct {
}

func NewMessages() *MessagesDAO {

	return &MessagesDAO{}
}

func (m *MessagesDAO) GetMessages(startDuration string, endDuration string) ([]models2.Messages, error) {

	//tx := m.db.MustBegin()

	//tx.MustExec("SELECT * FROM SentMesId INNER JOIN Receive ON Receive.message_id = SentMesId.message_id INNER JOIN Send ON SentMesId.sequence = Send.sequence where Send.date BETWEEN toDateTime('$1') and toDateTime('$2');")

	var mes []models2.Messages

	q := fmt.Sprintf("SELECT Send.date,submit_date,done_date,dest_addr,id,sms_text, Send.source_addr, SentMesId.message_id, stat FROM SentMesId INNER JOIN Receive ON Receive.message_id = SentMesId.message_id INNER JOIN Send ON SentMesId.sequence = Send.sequence and Receive.destination_addr = Send.dest_addr where Send.date BETWEEN toDateTime('%s') and toDateTime('%s');", startDuration, endDuration)
	err := config.Config.DB.Select(&mes, q)

	log.Println(len(mes))

	if err != nil {
		return nil, err
	}

	return mes, nil
}
