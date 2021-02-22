package amqp

import(
    "github.com/streadway/amqp"
)

type ConsumeMessage struct {
    Delivery amqp.Delivery
    Error error
}

type ConsumeConfig struct { 
    QueueConfig QueueConfig
    consumer string
    exclusive bool
    autoAck  bool
    noLocal bool
    noWait bool
    args map[string]interface{}
}

func (s *Queue) Consume(queueName string, opts ...ConsumeOption) <-chan *ConsumeMessage {

    message := make(chan *ConsumeMessage, 1)

    c := &ConsumeConfig{
        QueueConfig: QueueConfig {
            durable: false,
            // delete when unused
            delete: false,
            exclusive: false,
            noWait: false,
            args: nil,
        },
        consumer: "",
        exclusive: false,
        autoAck: true,
        noLocal: false,
        noWait: false,
        args: nil,
    }

    for _, opt := range opts {
        if err := opt(c); err != nil {
            message <- &ConsumeMessage{
                Delivery: amqp.Delivery{},
                Error: err,
            }
        }
    }

    q, err := s.declare(queueName, c.QueueConfig)

    if err != nil {
        message <- &ConsumeMessage{
            Delivery: amqp.Delivery{},
            Error: err,
        }
        return message
    }


    msgs, err := s.channel.Consume(
        q.Name,       // queue
        c.consumer,   // consumer
        c.autoAck,    // auto-ack
        c.exclusive,  // exclusive
        c.noLocal,    // no-local
        c.noWait,     // no-wait
        c.args,       // args
    )

    if err != nil {
        message <- &ConsumeMessage{
            Delivery: amqp.Delivery{},
            Error: err,
        }
        return message
    }
      
    go func() {

        for msg := range msgs {
            message <- &ConsumeMessage{ 
                Delivery: msg,
                Error: nil,
            }
        }
    }()
    
    return message
}