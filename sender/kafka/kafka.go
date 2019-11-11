// Copyright 2019 The OpenZipkin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package kafka implements a Kafka reporter to send spans to a Kafka server/cluster.
*/
package kafka

import (
	"github.com/Shopify/sarama"
)

// defaultKafkaTopic sets the standard Kafka topic our Reporter will publish
// on. The default topic for zipkin-receiver-kafka is "zipkin", see:
// https://github.com/openzipkin/zipkin/tree/master/zipkin-receiver-kafka
const defaultKafkaTopic = "zipkin"

type kafkaSender struct {
	producer sarama.AsyncProducer
	topic    string
}

// SenderOption sets a parameter for the kafkaSender
type SenderOption func(s *kafkaSender)

// Producer sets the producer used to send messages to Kafka.
func Producer(p sarama.AsyncProducer) SenderOption {
	return func(s *kafkaSender) {
		s.producer = p
	}
}

// Topic sets the Kafka topic name to send messages to.
func Topic(t string) SenderOption {
	return func(s *kafkaSender) {
		s.topic = t
	}
}

func (r *kafkaSender) Send(b []byte) {
	r.producer.Input() <- &sarama.ProducerMessage{
		Topic: r.topic,
		Key:   nil,
		Value: sarama.ByteEncoder(b),
	}
}
