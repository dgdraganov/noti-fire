package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dgdraganov/noti-fire/drivers/email"
	"github.com/dgdraganov/noti-fire/drivers/slack"
	"github.com/dgdraganov/noti-fire/drivers/sms"
	"github.com/dgdraganov/noti-fire/internal/dispatch"
	"github.com/dgdraganov/noti-fire/internal/notify"
	"github.com/dgdraganov/noti-fire/pkg/config"
	"github.com/dgdraganov/noti-fire/pkg/consume"
	"github.com/dgdraganov/noti-fire/pkg/kafka"
	"github.com/dgdraganov/noti-fire/pkg/log"
	"go.uber.org/zap/zapcore"
)

func main() {

	conf, err := config.NewConsumerConfig()
	if err != nil {
		panic(fmt.Sprintf("new config: %s", err))
	}

	logger := log.NewZapLogger(conf.ConsumerName, zapcore.InfoLevel)

	reader := kafka.NewKafkaReader(conf.KafkaConsumerConfig)
	consumer := consume.NewMessageConsumer(reader)

	dispatcher := dispatch.NewNotificationDispatcher(logger)
	dispatcher.RegisterDriver("sms", sms.NewSMSDriver())
	dispatcher.RegisterDriver("slack", slack.NewSlackDriver())
	dispatcher.RegisterDriver("email", email.NewEmailDriver())

	notifyer := notify.NewNotifyer(consumer, dispatcher, logger)

	ctx, cancel := context.WithCancel(context.Background())

	notifyer.Process(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sig

	logger.Info("shut down signal received")
	cancel()
	<-time.After(time.Second * 10)
}
