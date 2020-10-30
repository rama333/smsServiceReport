package daos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"smsServiceReport/internal/restapi/messages/models"
)

type MessagesDAO struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *MessagesDAO {

	return &MessagesDAO{db: db}
}

func (m *MessagesDAO) GetMessages(startDuration string, endDuration string) ([]models.Messages, error) {

	//tx := m.db.MustBegin()

	//tx.MustExec("SELECT * FROM SentMesId INNER JOIN Receive ON Receive.message_id = SentMesId.message_id INNER JOIN Send ON SentMesId.sequence = Send.sequence where Send.date BETWEEN toDateTime('$1') and toDateTime('$2');")

	var mes []models.Messages

	q := fmt.Sprintf("SELECT Send.date,submit_date,done_date,dest_addr,id,sms_text, Send.source_addr, SentMesId.message_id, stat FROM SentMesId INNER JOIN Receive ON Receive.message_id = SentMesId.message_id INNER JOIN Send ON SentMesId.sequence = Send.sequence where Send.date BETWEEN toDateTime('%s') and toDateTime('%s');", startDuration, endDuration)
	err := m.db.Select(&mes, q)

	if err != nil {
		return nil, err
	}

	return mes, nil
}
