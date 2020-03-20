package packages

import "testing"

func TestRabbitMQProvider(t *testing.T) {
	for i :=0;i < 10;i ++ {
		RabbitMQProvider()
	}
}

func TestRabbitMQConsumer(t *testing.T) {
	RabbitMQConsumer()
}
