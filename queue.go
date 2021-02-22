package amqp

import (
    "errors"
    "github.com/streadway/amqp"
)

var (
    ErrConnectAMQP = errors.New("Can't connect to amqp.")
    ErrOpenChannel = errors.New("Can't open channel.")
)

func NewQueue(url string) (*Queue, error) {
    q := &Queue{dial: url}
    
    err := q.connect()

    if err != nil {
        return q, err
    }

    return q, nil
}

type Queue struct {
    dial string
    conn *amqp.Connection
    channel *amqp.Channel
}

func (s *Queue) connect() error {
    conn, err := amqp.Dial(s.dial)
    
    if err != nil {
        return ErrConnectAMQP
    }

    ch, err := conn.Channel()

    if err != nil {
        return ErrOpenChannel
    }

    s.channel = ch
    s.conn = conn
    return nil
}

func (s *Queue) Close() {
    s.channel.Close()
    s.conn.Close()
}

func (s *Queue) declare(name string, c QueueConfig) (amqp.Queue, error) {

    q, err := s.channel.QueueDeclare(
        name,
        c.durable,
        c.delete,   // delete when unused
        c.exclusive,
        c.noWait,
        c.args,
      )

    if err != nil {
        return q, err
    }

    return q, nil
}