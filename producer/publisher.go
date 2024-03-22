package main

import (
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

	producer(conn)
}

func producer(conn *amqp.Connection) {

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	err = ch.Publish(
		constants.EXCHANGE_NAME,
		constants.ROUTING_KEY,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
