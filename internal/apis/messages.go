package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"smsServiceReport/internal/daos"
	"smsServiceReport/internal/models"
	"smsServiceReport/internal/restapi"
	"smsServiceReport/internal/services"
)

func GetMessages(c *gin.Context) {

	s := services.NewServiceMessages(daos.NewMessages())

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

	if messages, err := s.GetMessages(dur.StartDuration, dur.EndDuration); err != nil || len(messages) == 0 {
		restapi.ResponseStatusNotFound("Status Not Found", c)
	} else {
		c.JSON(http.StatusOK, messages)
	}

}

func GetUserMessages(c *gin.Context) {
	s := services.NewServicesUserMessages(daos.NewUserMessages())

	req := json.NewDecoder(c.Request.Body)

	var dur models.DurationDateUser
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

	if messages, err := s.GetUserMessages(dur.Dest_adr, dur.StartDuration, dur.EndDuration); err != nil || len(messages) == 0 {
		restapi.ResponseStatusNotFound("Status Not Found", c)
	} else {
		c.JSON(http.StatusOK, messages)
	}

}
