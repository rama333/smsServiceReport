package daos

import (
	"github.com/jmoiron/sqlx"
	"smsServiceReport/internal/restapi/messages/models"
	"time"
)

type MessagesDAO struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *MessagesDAO {

	return &MessagesDAO{db: db}
}

func (m *MessagesDAO) GetMessages(startDuration time.Time, endDuration time.Time) ([]models.Messages, error) {

	tx := m.db.MustBegin()

	tx.MustExec("select 1")

	return nil, nil
}
