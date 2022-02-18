package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Cannot connect in Rabbitmq : ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed Open Channel : ", err)
	}

	q, err := ch.QueueDeclare(
		"TRY-MESSAGE",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println("Declare is error : ", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println("Failed to register a consumer : ", err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s\n", d.Body)
		}
	}()

	log.Println("[] Waiting For Messages. To Exit press CTRL+C")

	<-forever
}
