package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"smsServiceReport/internal/apis"
)

type RESTAPI struct {
	server *gin.Engine
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *RESTAPI {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/")
	{
		v1.POST("StatMessage", apis.GetMessages)
		v1.POST("UserMessages", apis.GetUserMessages)
		//v1.POST("messages", apis.GetSumService)
	}

	return &RESTAPI{server: r, logger: logger}
}

func (rapi *RESTAPI) Start(port int) {

	rapi.server.Run(fmt.Sprintf(":%v", port))

	//go func() {
	//	rapi.server.Run(fmt.Sprintf(":%v", port))
	//}()
}
