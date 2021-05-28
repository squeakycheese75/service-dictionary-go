package controllers

import (
	"encoding/json"
	"net/http"

	"log"

	"github.com/gorilla/mux"
	satori_uuid "github.com/satori/go.uuid"
	"github.com/squeakycheese75/service-dictionary-go/api/env"
	"github.com/squeakycheese75/service-dictionary-go/api/utils"
	"github.com/streadway/amqp"
)

const RABBITMQ_QUEUE_NAME = "sd_orchestrator"
const RABBITMQ_CONNECTION = "amqp://guest:guest@localhost:5672/"

type Message struct {
	Name   string
	UUID   string
	Params []string
}

// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Printf("%v: %v", msg, err)
// 		return
// 	}
// }

func requestDataRpc(msg Message) (res string, err error) {
	conn, err := amqp.Dial(RABBITMQ_CONNECTION)
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Printf("Failed to register a consumer: %v", err)
		return
	}

	// Create correlationId
	corrId := satori_uuid.NewV4().String()
	json_msg, err := json.Marshal(msg)

	err = ch.Publish(
		"",                  // exchange
		RABBITMQ_QUEUE_NAME, // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          json_msg,
			// Body:          []byte(uuid),
		})
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
		return
	}

	for d := range msgs {
		if corrId == d.CorrelationId {

			res = string(d.Body)
			break
		}
	}
	return
}

func GetData(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		uuid := vars["id"]

		msg := Message{"Bob", uuid, nil}

		resval, err := requestDataRpc(msg)
		if err != nil {
			utils.RespondWithError(res, http.StatusServiceUnavailable, "Unable to request data from orchestrator")
			return
		}
		// Then look up dataset with uuid
		utils.RespondWithJSON(res, http.StatusOK, resval)
	}
}
