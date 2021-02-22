package amqp

import (
    "github.com/streadway/amqp"
)

type PublishConfig struct {
    QueueConfig QueueConfig
    exchange   string // publish to an exchange
    routingKey string // routing to 0 or more queues
    mandatory  bool   // mandatory
    immediate  bool   // immediate
    publishing amqp.Publishing
}

func (s *Queue) Publish(queueName string, data []byte, opts ...PublishOption) error {

    p := &PublishConfig{
        QueueConfig: QueueConfig {
            durable: false,
            // delete when unused
            delete: false,
            exclusive: false,
            noWait: false,
            args: nil,
        },
    }

    for _, opt := range opts {
        if err := opt(p); err != nil {
            return err
        }
    }

    q, err := s.declare(queueName, p.QueueConfig)

    if err != nil {
        return err
    }

    if err := s.PublishWithQueue(q, data, opts...); err != nil {
        return err
    }

    return nil
}

func (s *Queue) PublishWithQueue(q amqp.Queue, data []byte, opts ...PublishOption) error {

    defaultPublishing := amqp.Publishing {
        ContentType: "text/plain",
        Body:        []byte(data),
    }

    p := &PublishConfig{
        exchange: "",
        routingKey: q.Name,
        mandatory: false,
        immediate: false,
        publishing: defaultPublishing,
    }

    for _, opt := range opts {
        if err := opt(p); err != nil {
            return err
        }
    }

    if err := s.channel.Publish(
      p.exchange,    // exchange
      p.routingKey,  // routing key
      p.mandatory,   // mandatory
      p.immediate,   // immediate
      p.publishing,
    ); err != nil {
        return err
    }

    return nil
}

type PublishFunc func(s *amqp.Channel, q amqp.Queue) error

func (s *Queue) PublishWithAction(queueName string, action PublishFunc, opts ...PublishOption) error {

    p := &PublishConfig{
        QueueConfig: QueueConfig {
            durable: false,
            // delete when unused
            delete: false,
            exclusive: false,
            noWait: false,
            args: nil,
        },
    }

    for _, opt := range opts {
        if err := opt(p); err != nil {
            return err
        }
    }

    q, err := s.declare(queueName, p.QueueConfig)

    if err != nil {
        return err
    }

    if err := action(s.channel, q); err != nil {
        return err
    }

    return nil
}
