package main

import (
	"rabbitdemo/consumer"
	"rabbitdemo/publisher"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		consumer.ConsumeMessages()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		publisher.PublishMessage("Hello RabbitMQ!")
	}()

	wg.Wait()
}
