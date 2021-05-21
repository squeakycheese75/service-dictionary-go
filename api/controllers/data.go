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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func requestDataRpc(uuid string) (res string, err error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

    corrId := satori_uuid.NewV4().String()

	type Message struct {
		Name string
		UUID string
	}
	
	m := Message{"Alice", uuid}
	b, err := json.Marshal(m)


	err = ch.Publish(
		"",          // exchange
		"sd_orchestrator", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:		   b,
			// Body:          []byte(uuid),
		})
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {

			res = string(d.Body)
			failOnError(err, "Failed to convert body to integer")
			break
		}
	}
	return
}

func GetData(env *env.Env) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		uuid := vars["id"]

		resval, err := requestDataRpc(uuid)
		if err != nil {
			utils.RespondWithError(res, http.StatusServiceUnavailable, "Unable to request data from orchestrator")
			return
		}
		// Then look up dataset with uuid
		utils.RespondWithJSON(res, http.StatusOK, resval)
	}
}