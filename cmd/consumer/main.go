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

// reader := kafka.NewKafkaReader(config.KafkaConsumerConfig{
// 	Brokers: "localhost:9092",
// 	Topic:   "notifications",
// })

// for {
// 	consumed, err := reader.ReadMessage(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(consumed.Message.Value))
// 	err = reader.CommitMessage(context.Background(), consumed)
// 	if err != nil {
// 		panic(err)
// 	}
// }
