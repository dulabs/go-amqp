package main

import (
    "log"
    "github.com/kxchange/go-amqp"
)

var (
    queueName = "testamqp"
)

func main() {

    q, err := amqp.NewQueue("amqp://guest:guest@127.0.0.1:5672/")

    if err != nil {
        log.Fatalln(err)
    }

    defer q.Close()

    log.Println("Connected to amqp.")

    for msg := range q.Consume(queueName) {

        if msg.Error != nil {
            log.Println(err)
        }else {
            log.Printf("Consume results: %v\n", string(msg.Delivery.Body))
        }
    }

}