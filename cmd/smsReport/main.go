package main

import (
	"go.uber.org/zap"
	"smsServiceReport/internal/diagnostics"
	"smsServiceReport/internal/resources"
	"smsServiceReport/internal/restapi/api"
	"smsServiceReport/internal/restapi/messages/daos"
	"smsServiceReport/internal/restapi/messages/services"
	daos2 "smsServiceReport/internal/restapi/userMessages/daos"
	services2 "smsServiceReport/internal/restapi/userMessages/services"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	slogger := logger.Sugar()
	slogger.Info("Starting the application...")
	slogger.Info("Reading configuration and initializing resources...")

	rsc, err := resources.New(slogger)
	if err != nil {
		slogger.Fatalw("Can't initialize resources.", "err", err)
	}
	defer func() {
		err = rsc.Release()
		if err != nil {
			slogger.Errorw("Got an error during resources release.", "err", err)
		}
	}()

	slogger.Info("Configuring the application units...")

	diag := diagnostics.New(slogger, rsc.Config.DiagPort, rsc.Healthz)
	diag.Start(slogger)
	slogger.Info("The application is ready to serve requests.")

	dbMessages := daos.New(rsc.Conn)
	serMessages := services.NewService(dbMessages)

	dbuserMessages := daos2.New(rsc.Conn)
	serUserMessages := services2.New(dbuserMessages)

	rapi := api.New(slogger, serMessages, serUserMessages)
	rapi.Start(rsc.Config.RESTAPIPort)

}
