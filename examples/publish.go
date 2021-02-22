package main

import (
    "log"
    "encoding/json"
    goamqp "github.com/kxchange/go-amqp"
)

var (
    queueName = "testamqp"
)

type EmailMessage struct {
	To   string `json:"to"`
	From string `json:"from"`
	Subject string `json:"subject"`
	Body string `json:"msg"`
	ReplyTo string `json:"reply_to"`
}

func main() {
    q, err := goamqp.NewQueue("amqp://guest:guest@127.0.0.1:5672/")

    if err != nil {
        log.Fatalln(err)
    }

    defer q.Close()

    req := EmaiilMessage{
        To: "user@example.com",
        From: "no-reply@example.com",
        Subject: "Sending email",
        Body: "<h2>email content</h2>",
        ReplyTo: "cs@example.com",
    }

    msg, err := json.Marshal(req)

    if err != nil {
        log.Fatalln(err)
    }

    err = q.Publish(queueName, msg)

    if err != nil {
		log.Println(err)
	}else{
		log.Printf("Sent to amqp queue %s.\n", queueName)
	}
}