package amqp

import (
	"github.com/streadway/amqp"
)

func WithConsumeQueue(durable bool, delete bool, exclusive bool, noWait bool, args map[string]interface{}) ConsumeOption {
	return func(cfg *ConsumeConfig) error {
		c := cfg.QueueConfig
		c.durable = durable
		c.delete = delete
		c.exclusive = exclusive
		c.noWait = noWait
		c.args = args

		cfg.QueueConfig = c
		return nil
	}
}

// Consumer Option configurates client with defined option.
type ConsumeOption func(*ConsumeConfig) error

func WithConsume(consumer string, autoAck bool, noLocal bool, exclusive bool, noWait bool, args map[string]interface{}) ConsumeOption {
	return func(c *ConsumeConfig) error {
		c.consumer = consumer
		c.autoAck = autoAck
		c.noLocal = noLocal
		c.exclusive = exclusive
		c.noWait = noWait
		c.args = args
		return nil
	}
}



type PublishOption func(*PublishConfig) error

func WithPublishQueue(durable bool, delete bool, exclusive bool, noWait bool, args map[string]interface{}) PublishOption {
	return func(cfg *PublishConfig) error {
		c := cfg.QueueConfig
		c.durable = durable
		c.delete = delete
		c.exclusive = exclusive
		c.noWait = noWait
		c.args = args

		cfg.QueueConfig = c
		return nil
	}
}

func WithPublish(exchange string, mandatory bool, immediate bool) PublishOption {
	return func(c *PublishConfig) error {
		c.exchange = exchange
		c.mandatory = mandatory
		c.immediate = immediate
		return nil
	}
}

func WithPublishing(publishing amqp.Publishing) PublishOption {
	return func(c *PublishConfig) error {
		c.publishing = publishing
		return nil
	}
}