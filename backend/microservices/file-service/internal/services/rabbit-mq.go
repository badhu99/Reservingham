package services

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type DataRabbitMQ struct {
	Connection *amqp.Connection
}

func InitRabbitMQ() DataRabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}

	return DataRabbitMQ{
		Connection: conn,
	}
}

func (data *DataRabbitMQ) UploadToQueue(channelName, body string) error {
	// defer data.Connection.Close()
	ch, err := data.Connection.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(channelName, false, false, false, false, nil) // "MyFirstTest"
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	// body := "Hello for my first test"
	err = ch.PublishWithContext(ctx,
		"", q.Name, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(body)})

	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s\n", body)
	return nil
}

func (data *DataRabbitMQ) GetDataFromQueue(queueName string) (string, error) {
	ch, err := data.Connection.Channel()
	if err != nil {
		return "", err
	}
	defer ch.Close()

	// Declare a queue
	_, err = ch.QueueDeclare(
		queueName, // queue name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return "", err
	}

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatal(err)
	}

	msg := <-msgs
	return string(msg.Body), nil
}
