package amqp

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func ConnectAMQP() {
	fmt.Println("Connecting to rabbitmq")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatalf("Could not connect to AMQP: %s", err)
	}

	channel, err := conn.Channel()

	if err != nil {
		log.Fatalf("%s, cannot create amqp channel", err)
	}

	channel.ExchangeDeclare("email", "direct", true, false, false, false, amqp.Table{})

	_, err = channel.QueueDeclare("email-queue", true, false, false, false, amqp.Table{})

	if err != nil {
		log.Fatalf("could not create queue: %s", err)
	}

	err = channel.Qos(1, 0, false)

	if err != nil {
		log.Fatalf("could not declare qos: %s", err)
	}

	channel.QueueBind("email-queue", "email-queue", "email", false, amqp.Table{})

	channel.Publish("email", "email-queue", false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte("This is a new message"),
		DeliveryMode: amqp.Persistent,
	})

	defer conn.Close()
	fmt.Println("Successfully connected to RabbitMQ instance.")
}
