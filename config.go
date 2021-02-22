package amqp

type QueueConfig struct {

    // durable
    durable     bool

    // delete when unused
    delete      bool

    exclusive   bool

    noWait      bool

    args        map[string]interface{}
}