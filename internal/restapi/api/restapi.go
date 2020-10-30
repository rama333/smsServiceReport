package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"smsServiceReport/internal/restapi/messages/apis"
	"smsServiceReport/internal/restapi/messages/services"
	_ "smsServiceReport/internal/restapi/userMessages/apis"
	apis2 "smsServiceReport/internal/restapi/userMessages/apis"
	_ "smsServiceReport/internal/restapi/userMessages/services"
	services2 "smsServiceReport/internal/restapi/userMessages/services"
)

type RESTAPI struct {
	server *gin.Engine
	error  chan error
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, mes *services.Service, ser *services2.ServiceUserMessages) *RESTAPI {

	handler := apis.MessaheHandler{
		MessageService: mes,
	}

	handlerUserMessages := apis2.UserMessagesHandler{
		Hand: ser,
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/")
	{
		v1.POST("StatMessage", handler.GetMessages)
		v1.POST("UserMessages", handlerUserMessages.GetMessages)
		//v1.POST("messages", apis.GetSumService)
	}

	return &RESTAPI{
		server: r,
		error:  make(chan error, 1),
		logger: logger,
	}
}

func (rapi *RESTAPI) Start(port int) {

	rapi.server.Run(fmt.Sprintf(":%v", port))

	//go func() {
	//	rapi.server.Run(fmt.Sprintf(":%v", port))
	//}()
}

func Stop(rapi *RESTAPI) {

}
