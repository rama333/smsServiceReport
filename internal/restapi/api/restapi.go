package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	_ "smsServiceReport/cmd/smsReport/docs"
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return &RESTAPI{server: r, logger: logger}
}

func (rapi *RESTAPI) Start(port int) {

	rapi.server.Run(fmt.Sprintf(":%v", port))

	//go func() {
	//	rapi.server.Run(fmt.Sprintf(":%v", port))
	//}()
}
