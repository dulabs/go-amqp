package main

import (
    "log"
    "encoding/json"
    "github.com/streadway/amqp"
    goamqp "github.com/kxchange/go-amqp"
)

var (
    queueName = "testamqp"
)

req := EmaiilMessage{
    To: "user@example.com",
    From: "no-reply@example.com",
    Subject: "Sending email",
    Body: "<h2>email content</h2>",
    ReplyTo: "cs@example.com",
}

func main() {
    q, err := goamqp.NewQueue("amqp://guest:guest@127.0.0.1:5672/")

    if err != nil {
        log.Fatalln(err)
    }

    defer q.Close()

    err = q.PublishWithAction(queueName, func(s *amqp.Channel, q amqp.Queue) error {

        req := EmaiilMessage{
            To: "user@example.com",
            From: "no-reply@example.com",
            Subject: "Sending email",
            Body: "<h2>email content</h2>",
            ReplyTo: "cs@example.com",
        }

        msg, err := json.Marshal(req)

        if err != nil {
            return err
        }

        if err := s.Publish(
            "",    // exchange
            q.Name,  // routing key
            false,   // mandatory
            false,   // immediate
            amqp.Publishing {
                ContentType: "text/plain",
                Body:        []byte(msg),
            },
        ); err != nil {
            return err
        }

        return nil
    })

    if err != nil {
		log.Println(err)
	}else{
		log.Printf("Sent to amqp queue %s.\n", queueName)
	}
}