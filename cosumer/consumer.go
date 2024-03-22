package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"github.com/toannt412/demo_rabbitMQ/constants"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	Consumer(conn)
}

func Consumer(conn *amqp.Connection) {
	// consumer.go
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		constants.EXCHANGE_NAME,
		constants.KIND,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	queue, err := ch.QueueDeclare(
		constants.QUEUE_NAME,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(
		queue.Name,
		constants.ROUTING_KEY,
		constants.EXCHANGE_NAME,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"", // consumer name
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs {
		fmt.Println(string(msg.Body))
	}

}
