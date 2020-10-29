package restapi

import (
	"awesomeProject/cmd/epsilon5000/apis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RESTAPI struct {
	server gin.Engine
	error  chan error
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, port int) *RESTAPI {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/")
	{
		v1.POST("StatMessage", apis.GetSumService)
		v1.POST("UserMessages", apis.GetSumService)
		v1.POST("messages", apis.GetSumService)
	}

}
