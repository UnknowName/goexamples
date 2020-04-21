package packages

import (
	"log"

	"github.com/streadway/amqp"
)

const rabbitUrl = "amqp://guest:guest@128.0.100.170:5672/test"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
	return
}

func RabbitMQProvider() {
	conn, err := amqp.Dial(rabbitUrl)
	failOnError(err, "Connect RabbitMQ Failed")
	defer conn.Close()
	channel, err := conn.Channel()
	failOnError(err, "Get RabbitMQ Channel Failed")
	defer channel.Close()
	// fmt.Println(channel)
	queue, err := channel.QueueDeclare(
		"go-queue",
		true,
		false,
		false,
		false,
		nil,
		)
	msg := "hello,world"
	failOnError(err, "Declare queue failed")
	err = channel.Publish(
		"",
		 queue.Name,
		false,
		false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte(msg)},
		)
	failOnError(err, "Send to msg to RabbitMQ failed")
}


func RabbitMQConsumer(){
	conn, err := amqp.Dial(rabbitUrl)
	failOnError(err, "Connect RabbitMQ Failed")
	defer conn.Close()
	channel, err := conn.Channel()
	failOnError(err, "Get RabbitMQ Channel Failed")
	defer channel.Close()
	queue, err := channel.QueueDeclare(
		"go-queue",
		true,
		false,
		false,
		false,
		nil, )
	failOnError(err, "Queue declare failed")
	// 返回一个只读的chan
	msgChan, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Get RabbitMQ message failed")
	for d := range msgChan {
		log.Printf("Recived Message from RabbitMQ: %s", d.Body)
	}
}