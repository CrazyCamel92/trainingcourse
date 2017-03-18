package main

import (
	"log"
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

func main() {
	//this connection string is default rabbit mq on the local machine
	connectionSting:= "amqp://guest:guest@localhost:5672/";
	queueData := ConsumerInit(connectionSting,"printing queue")
	Consume(queueData)
	defer queueData.DisposeQueue()
}

type QueueData struct {
	Connection *amqp.Connection
	Channel *amqp.Channel
	Q amqp.Queue
	ResponseChannel <-chan(amqp.Delivery)
}

func (data QueueData) DisposeQueue ()  {
	data.Channel.Close()
	data.Connection.Close()
}

func ConsumerInit(conStr string,queueName string ) QueueData  {
	conn, err := amqp.Dial(conStr)
	onError(err, "Failed to Server")
	//defer conn.Close()

	//this is a rabbit mq channel.
	ch, err := conn.Channel()
	onError(err, "Failed to open a channel")
	//defer ch.Close()

	// each worker will get only 1 message at a time
	ch.Qos(1,0,false)

	q, err := ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	onError(err, "Failed to declare a queue")
	msgChannel, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	onError(err, "Failed to register a consumer")

	return QueueData{Connection:conn,Channel:ch,Q:q,ResponseChannel:msgChannel}
}

func Consume(queueData QueueData)  {
	//we create this channel to set up an infinite waiting loop
	infiniteChannel := make(chan bool)

	go func() {
		for d := range queueData.ResponseChannel {
			// the consumer will sleep for 2 seconds in order to see multiple workers working together
			time.Sleep(2 * time.Second)
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("waiting for new messages ")
	<-infiniteChannel
}

func onError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}