package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"smsServiceReport/internal/restapi"
	"smsServiceReport/internal/restapi/messages/models"
	"smsServiceReport/internal/restapi/messages/services"
)

type MessaheHandler struct {
	MessageService *services.Service
}

func (m *MessaheHandler) GetMessages(c *gin.Context) {
	req := json.NewDecoder(c.Request.Body)

	var dur models.DurationDate
	err := req.Decode(&dur)

	if err != nil {
		restapi.ResponseBadRequest("Couldn't parse request body", c)
		return
	}

	//layout := "2006-01-02 15:04:05"
	//
	//startDur, err := time.Parse(layout, dur.StartDuration)
	//if err != nil {
	//	restapi.ResponseBadRequest(err.Error(), c)
	//}
	//endDur, err := time.Parse(layout, dur.StartDuration)
	//if err != nil {
	//	restapi.ResponseBadRequest("Couldn't parse request body", c)
	//}

	if messages, err := m.MessageService.GetMessages(dur.StartDuration, dur.EndDuration); err != nil || len(messages) == 0 {
		restapi.ResponseStatusNotFound("Status Not Found", c)
	} else {
		c.JSON(http.StatusOK, messages)
	}

}
