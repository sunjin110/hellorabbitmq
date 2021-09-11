package main

import (
	"fmt"
	"hellorabbitmq/pkg/chk"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Receive Queue")

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
	chk.SE(err, "failed make queue")

	// receive message
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	chk.SE(err, "failede consume")

	go func() {
		for d := range msgs {
			log.Printf("receive message is %s\n", d.Body)
		}
	}()

	// 永遠に待つ
	forever := make(chan bool)
	<-forever

}
