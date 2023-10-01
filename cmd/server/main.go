package main

import (
	"fmt"

	"github.com/dgdraganov/noti-fire/internal/http/handler/notification"
	"github.com/dgdraganov/noti-fire/internal/http/middleware"
	"github.com/dgdraganov/noti-fire/internal/http/router"
	"github.com/dgdraganov/noti-fire/internal/http/server"
	"github.com/dgdraganov/noti-fire/internal/notifyer"
	"github.com/dgdraganov/noti-fire/pkg/config"
	"github.com/dgdraganov/noti-fire/pkg/kafka"
	"github.com/dgdraganov/noti-fire/pkg/log"
	"go.uber.org/zap/zapcore"
)

func main() {

	conf, err := config.NewServerConfig()
	if err != nil {
		panic(fmt.Sprintf("new config: %s", err))
	}

	logger := log.NewZapLogger(conf.ServiceName, zapcore.InfoLevel)

	kafkaProducer := kafka.NewKafkaProducer(conf.KafkaProducerConfig)
	notifyer := notifyer.NewNotifyerAction(kafkaProducer)

	// middleware initialization
	i := middleware.NewRequestIdMiddleware(logger)
	a := middleware.NewAuthenticatorMiddleware(logger)
	l := middleware.NewLoggerMiddleware(logger)

	notificationHandler := i.Id(l.Log(a.Auth(notification.NewNotificationHandler("POST", notifyer, logger))))

	serviceRouter := router.NewNotificationRouter()

	// POST method
	serviceRouter.Register("/notify", notificationHandler)

	logger.Infow(
		"server starting...",
		"service_port", conf.ServerPort,
	)

	// todo: add graceful shut down

	server := server.NewHTTPServer(serviceRouter.ServeMux(), logger)
	if err = server.Start(conf.ServerPort); err != nil {
		logger.Fatalf("server stopped unexpectedly: %s", err)
	}
}
