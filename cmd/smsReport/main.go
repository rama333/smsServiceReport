package main

import (
	"fmt"
	"go.uber.org/zap"
	"smsServiceReport/internal/config"
	"smsServiceReport/internal/diagnostics"
	"smsServiceReport/internal/rabbitMQ"
	"smsServiceReport/internal/resources"
	"smsServiceReport/internal/restapi/api"
)

// @title sms service report Swagger API
// @version 1.0
// @description Swagger API for Golang Project sms service report Swagger API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support

//@BasePath /api/v1
func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	slogger := logger.Sugar()
	slogger.Info("Starting the application...")
	slogger.Info("Reading configuration and initializing resources...")

	if err := config.LoadConfig("/Users/ramilramilev/go/src/smsServiceReport/config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	rsc, err := resources.New(slogger)
	if err != nil {
		slogger.Fatalw("Can't initialize resources.", "err", rsc)
	}

	defer func() {
		err := rsc.Release()
		if err != nil {
			slogger.Errorw("Got an error during resources release.", "err", err)
		}
	}()

	slogger.Info("Configuring the application units...")

	diag := diagnostics.New(slogger, config.Config.DIAGPORT, rsc.Healthz)
	diag.Start(slogger)
	slogger.Info("The application is ready to serve requests.")

	rabbitChan, err := rabbitMQ.StartRabbitMQ()
	if err != nil {
		slogger.Info(err)
	}

	rapi := api.New(slogger)
	rapi.Start(config.Config.RESTAPIPort)

	rabbitChan.Close()

}
