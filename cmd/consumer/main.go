package main

import (
	"fmt"

	"github.com/dgdraganov/noti-fire/pkg/config"
	"github.com/dgdraganov/noti-fire/pkg/log"
	"go.uber.org/zap/zapcore"
)

func main() {

	conf, err := config.NewConsumerConfig()
	if err != nil {
		panic(fmt.Sprintf("new config: %s", err))
	}

	logger := log.NewZapLogger(conf.ConsumerName, zapcore.InfoLevel)
	logger.Info("starting")

}
