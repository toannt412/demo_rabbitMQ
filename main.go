package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

	// consumer.go
ch, err := conn.Channel()
if err != nil {
    log.Fatal(err)
}
defer ch.Close()

err = ch.ExchangeDeclare(
    "hello_exchange",
    "direct",
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
    "hello",
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
    "world",
    "hello_exchange",
    false,
    nil,
)
if err != nil {
    log.Fatal(err)
}


}
