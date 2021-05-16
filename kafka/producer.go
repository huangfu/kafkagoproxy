package kafka

import (
	"log"
	"strconv"
	"strings"

	"github.com/Shopify/sarama"
)

var config = sarama.NewConfig()

func initSeting() {
	config.Producer.MaxMessageBytes = 1024 * 1024 * 1024
}

//config.Producer.Timeout =time
func SendMessage(k Kafkaconn) (partition int32, offset int64, err error) {

	initSeting()
	config.Producer.Return.Successes = true //成功返回
	var inputparti int
	if k.Partion != "" {
		if inputparti, err = strconv.Atoi(k.Partion); err != nil {
			return
		} else {
			config.Producer.Partitioner = sarama.NewManualPartitioner
		}
	} else {
		inputparti = -1
	}

	producer, errp := sarama.NewSyncProducer(strings.Split(k.Brokers, ","), config)
	if errp != nil {
		err = errp
		log.Println("Failed to create producer:", err)
		return
	}
	defer func() {
		if errdef := producer.Close(); errdef != nil {

			err = errdef
		}
	}()
	var msg *sarama.ProducerMessage
	if inputparti != -1 {
		msg = &sarama.ProducerMessage{Topic: k.Topics, Partition: int32(inputparti), Value: sarama.StringEncoder(string(k.Data))}
	} else {

		msg = &sarama.ProducerMessage{Topic: k.Topics, Value: sarama.StringEncoder(string(k.Data))}
	}

	partition, offset, err = producer.SendMessage(msg)

	return

}

func AsycSendMessage(k Kafkaconn) error {
	initSeting()
	producer, err := sarama.NewAsyncProducer(strings.Split(k.Brokers, ","), config)
	if err != nil {
		return err
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println("Closed Failed to produce", err)
		}
	}()
	select {
	case producer.Input() <- &sarama.ProducerMessage{Topic: k.Topics, Key: nil, Value: sarama.StringEncoder(k.Data)}:
		return nil
	case err := <-producer.Errors():
		log.Println("Failed to produce message", err)
		return err
	}

}
