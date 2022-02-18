package main

import (
	"github.com/streadway/amqp"
	"log"
	"time"
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

	for {

		log.Println("sending...")
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte("HELLO MESSAGING RABBITMQ"),
			})
		if err != nil {
			log.Println("Publish is error : ", err)
		}
		time.Sleep(2 * time.Minute)
	}

}

//docker run --name rabbitmq -p 5672:5672
