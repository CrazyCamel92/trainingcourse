package main

import (
//	"bufio"
//	"os"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

func main() {
	/*for i:=0;i<=100;i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("etner text to produce: ")
		text, _ := reader.ReadString('\n')
		PushNewMessage(text)
		println("new message was pushed to queue");
	}*/
	queueData:= QueueHandler("printing queue")
	for i:=0;i<=100;i++{
		NewMessage(strconv.Itoa(i),queueData)
	}
	queueData.DisposeQueue()
}

type QueueData struct {
	Connection *amqp.Connection
	Channel *amqp.Channel
	Q amqp.Queue
}

func (data QueueData) DisposeQueue ()  {
	data.Channel.Close()
	data.Connection.Close()
}

func NewMessage(msg string,queueData QueueData )  {
	body := msg
	err := queueData.Channel.Publish(
		"",     // exchange
		queueData.Q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	onError(err, "Failed to publish a message")
}

func QueueHandler(queueName string) QueueData {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	onError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	//creating rabbit mq channel
	ch, err := conn.Channel()
	onError(err, "Failed to open a channel")
	//defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	onError(err, "Failed to declare a queue")

	return QueueData{Connection: conn,Channel:ch,Q:q};
}
func onError(err error,msg string){
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
