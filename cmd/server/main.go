package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dgdraganov/noti-fire/internal/http/handler/notification"
	"github.com/dgdraganov/noti-fire/internal/http/middleware"
	"github.com/dgdraganov/noti-fire/internal/http/router"
	"github.com/dgdraganov/noti-fire/internal/http/server"
	"github.com/dgdraganov/noti-fire/internal/process"
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

	logger := log.NewZapLogger(conf.ServerName, zapcore.InfoLevel)

	// middleware initialization
	i := middleware.NewRequestIdMiddleware(logger)
	a := middleware.NewAuthenticatorMiddleware(logger)
	l := middleware.NewLoggerMiddleware(logger)

	kafkaWriter := kafka.NewKafkaWriter(conf.KafkaProducerConfig)
	kafkaProducer := producer.NewMessageProducer(kafkaWriter)
	processor := process.NewProcessAction(kafkaProducer)

	var processHandler http.Handler
	processHandler = notification.NewNotificationHandler(http.MethodPost, processor, logger)
	processHandler = i.Id(l.Log(a.Auth(processHandler)))

	serviceRouter := router.NewNotificationRouter()

	// POST
	serviceRouter.Register("/notify", processHandler)

	logger.Infow(
		"server starting...",
		"service_port", conf.ServerPort,
	)

	server := server.NewHTTPServer(conf.ServerPort, serviceRouter.ServeMux(), logger)
	// starts the server asynchronously
	server.Start(conf.ServerPort)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sig

	logger.Info("shut down signal received")
	server.Shutdown()
}
