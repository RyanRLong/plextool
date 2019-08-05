// Package plextool .
//
// Copyright 2015 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package plextool

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// IncomingData is used to store json the response
// when receiving data from RabbitMQ.
type IncomingData struct {
	Account  map[string]interface{}
	Metadata map[string]interface{}
	Event    string
}

// GetFilmData Extracts the film data from request
// TODO not sure if this belongs here
func (data IncomingData) GetFilmData() Film {
	if data.Metadata["title"] == nil {
		data.Metadata["title"] = ""
	}
	if data.Metadata["viewOffset"] == nil {
		data.Metadata["viewOffset"] = 0.0
	}
	if data.Metadata["viewCount"] == nil {
		data.Metadata["viewCount"] = 0.0
	}
	Film := Film{
		Title:      data.Metadata["title"].(string),
		ViewOffset: data.Metadata["viewOffset"].(float64) / 60000, // why this number
		ViewCount:  data.Metadata["viewCount"].(float64),
	}
	return Film
}

// Receiver contains the Connection and Channel for a RabbitMQ
// receiver
type Receiver struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

// NewReceiver creates and returns a receiver instance pointing
// to the RabbitMQ exchange
func NewReceiver() (receiver *Receiver, err error) {
	exchangeActivityURL := os.Getenv("EXCHANGE_ACTIVITY_URL")
	if exchangeActivityURL == "" {
		err = &receiverError{"Environment variable \"EXCHANGE_ACTIVITY_URL\" not found."}
		return
	}
	conn, err := amqp.Dial(exchangeActivityURL)
	if err != nil {
		return
	}

	ch, err := conn.Channel()
	if err != nil {
		return
	}

	receiver = &Receiver{
		Connection: conn,
		Channel:    ch,
	}
	return
}

// Listen listens for exchange messages from the RabbitMQ
// exchange.
func (Receiver Receiver) Listen() (msgs <-chan amqp.Delivery) {
	err := Receiver.Channel.ExchangeDeclare(
		"activity", // name
		"fanout",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Panic(err)
	}

	q, err := Receiver.Channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Panic(err)
	}

	err = Receiver.Channel.QueueBind(
		q.Name,     // queue name
		"",         // routing key
		"activity", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}

	msgs, err = Receiver.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Panic(err)
	}

	return
}

type receiverError struct {
	message string
}

func (e *receiverError) Error() string {
	return fmt.Sprintf("%s", e.message)
}
