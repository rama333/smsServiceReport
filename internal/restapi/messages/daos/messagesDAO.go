package daos

import (
	"github.com/jmoiron/sqlx"
	"smsServiceReport/internal/restapi/messages/models"
)

type MessagesDAO struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *MessagesDAO {

	return &MessagesDAO{db: db}
}

func (m *MessagesDAO) getMessages() ([]models.Messages, error) {

	tx := m.db.MustBegin()

	tx.MustExec("select ")

}
