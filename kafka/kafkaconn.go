package kafka

import "encoding/json"

type Kafkaconn struct {
	Topics  string          `json:"topics"`
	Brokers string          `json:"brokers"`
	Data    json.RawMessage `json:"data"`
	Partion string          `json:"partion"`
}

type KafkaResult struct {
	Topics  string `json:"topics"`
	Partion int32  `json:"partion"`
	Offset  int64  `json:"offset"`
}
