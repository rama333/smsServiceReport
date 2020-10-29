package messages

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"smsServiceReport/internal/restapi"
	"smsServiceReport/internal/restapi/messages/models"
)

type messaheHandler struct {
	messageService Service
}

func (m *messaheHandler) getMessages(c *gin.Context) {
	req := json.NewDecoder(c.Request.Body)

	var dur models.DurationDate
	err := req.Decode(&dur)

	if err != nil {
		restapi.ResponseBadRequest("Couldn't parse request body", c)
		return
	}

	if messages, err := m.messageService.GetMessages(dur.StartDuration, dur.EndDuration); err != nil {
		restapi.ResponseStatusNotFound("Status not found", c)
	} else {
		c.JSON(http.StatusOK, messages)
	}
}
