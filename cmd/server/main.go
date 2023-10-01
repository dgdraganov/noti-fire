package main

import (
	"fmt"

	"github.com/dgdraganov/noti-fire/internal/http/handler/notification"
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
	notifyHdlr := notification.NewNotificationHandler(notifyer)

	serviceRouter := router.NewNotificationRouter()
	serviceRouter.Register("/notify", notifyHdlr)

	serveMux := serviceRouter.ServeMux()

	logger.Infow(
		"server starting...",
		"service_port", conf.ServerPort,
	)

	// todo: add graceful shut down

	server := server.NewHTTPServer(serveMux, logger)
	if err = server.Start(conf.ServerPort); err != nil {
		logger.Fatalf("server stopped unexpectedly: %s", err)
	}
}
