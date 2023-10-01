package main

import (
	"fmt"
	"net/http"

	"github.com/dgdraganov/noti-fire/internal/http/handler/notification"
	"github.com/dgdraganov/noti-fire/internal/http/middleware"
	"github.com/dgdraganov/noti-fire/internal/http/router"
	"github.com/dgdraganov/noti-fire/internal/http/server"
	"github.com/dgdraganov/noti-fire/internal/notifyer"
	"github.com/dgdraganov/noti-fire/pkg/config"
	"github.com/dgdraganov/noti-fire/pkg/kafka"
	"github.com/dgdraganov/noti-fire/pkg/log"
	"github.com/dgdraganov/noti-fire/pkg/producer"
	"go.uber.org/zap/zapcore"
)

func main() {

	conf, err := config.NewServerConfig()
	if err != nil {
		panic(fmt.Sprintf("new config: %s", err))
	}

	logger := log.NewZapLogger(conf.ServiceName, zapcore.InfoLevel)

	// middleware initialization
	i := middleware.NewRequestIdMiddleware(logger)
	a := middleware.NewAuthenticatorMiddleware(logger)
	l := middleware.NewLoggerMiddleware(logger)

	kafkaWriter := kafka.NewKafkaWriter(conf.KafkaProducerConfig)
	kafkaProducer := producer.NewMessageProducer(kafkaWriter)
	notifyer := notifyer.NewNotifyerAction(kafkaProducer)

	var notificationHandler http.Handler
	notificationHandler = notification.NewNotificationHandler("POST", notifyer, logger)
	notificationHandler = i.Id(l.Log(a.Auth(notificationHandler)))

	serviceRouter := router.NewNotificationRouter()

	// POST
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
