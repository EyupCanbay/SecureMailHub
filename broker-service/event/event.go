package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name
		"topic",      //tpye
		true,		  //durable?
		false, 		 //auto-deleted?
		false,			// internal
		false,			//no-wait
		nil,		//arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",		//name
		false,		//durable
		false,		//deletewhen unused
		true,		//exculusive
		false,		//no-wait
		nil,		//arguements
	)
}