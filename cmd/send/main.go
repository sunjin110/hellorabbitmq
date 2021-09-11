package main

import (
	"fmt"

	"hellorabbitmq/pkg/chk"

	"github.com/streadway/amqp"
)

// queueのoptionがあって、それに従って queueの特徴を定義することができる
// https://www.rabbitmq.com/queues.html

// send message
func main() {
	fmt.Println("Hello RabbitMQ")

	// connect
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	chk.SE(err, "failed dial")
	defer conn.Close()

	// make channel
	ch, err := conn.Channel()
	chk.SE(err, "failed make channel")
	defer ch.Close()

	// make queue
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	chk.SE(err, "failed create queue")

	body := "Hello World! aaaaa"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	chk.SE(err, "failed to publish message")

}
