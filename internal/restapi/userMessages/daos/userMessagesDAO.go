package daos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"smsServiceReport/internal/restapi/userMessages/models"
)

type UserMessagesDAO struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *UserMessagesDAO {

	return &UserMessagesDAO{db: db}
}

func (userMes UserMessagesDAO) GetUserMessages(dest_add string, startDuration string, endDuration string) ([]models.Messages, error) {

	var mes []models.Messages

	q := fmt.Sprintf("SELECT Send.date,submit_date,done_date,dest_addr,id,sms_text, Send.source_addr, SentMesId.message_id, stat FROM SentMesId INNER JOIN Receive ON Receive.message_id = SentMesId.message_id INNER JOIN Send ON SentMesId.sequence = Send.sequence  where Send.date BETWEEN toDateTime('%s') and toDateTime('%s') and dest_addr = %s;", startDuration, endDuration, dest_add)
	err := userMes.db.Select(&mes, q)

	if err != nil {
		return nil, err
	}

	return mes, nil
}
