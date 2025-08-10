/*
/
- aplicação do cliente.
- conectar em servidor
- subscreve em topicos
- posta mensagens
- recebe mensagens
/
*/
package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type connection struct {
	status             bool
	broker_url         string
	client             mqtt.Client
	topic              string
	publishHandler     mqtt.MessageHandler
	connectionHandler  mqtt.OnConnectHandler
	connectLostHandler mqtt.ConnectionLostHandler
}

// PUBLISHED handler

// 1 - função para conectar ao servidor MQTT

func (c *connection) connectBroker() error {

	// config client
	c.broker_url = "tcp://localhost:1883"
	c.topic = "test/topic"

	c.publishHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s on topic: %s \n", msg.Payload(), msg.Topic())
	}

	c.connectionHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}

	c.connectLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connection LOST: %v", err)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.broker_url)
	opts.SetDefaultPublishHandler(c.publishHandler)

	// create and connect the client
	c.client = mqtt.NewClient(opts)

	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect: %v", token.Error())
	}

	fmt.Println("Connected to BROKER  Server \n")
	c.status = true
	return nil
}

// 2 - função para enviar mensagem

func (c *connection) publishMsg(topicId string, msg string) error {

	// publish operations
	payload := "hello from go"

	if token := c.client.Publish(c.topic, 0, false, payload); token.Wait() && token.Error() != nil {
		log.Printf("Failed to publish : %v", token.Error())
	}

	fmt.Printf("Published message: %s to topic: %s \n", payload, c.topic)

	return nil
}

// 3 - função para receber mensagens do topico
func (c *connection) subscribeTopic() error {
	// subscribe operations
	if token := c.client.Subscribe(c.topic, 1, nil); token.Wait() && token.Error() != nil {

		log.Fatalf("Failed to subscribe: %v", token.Error())
	}

	fmt.Printf("Subscribe to topic : %s  \n", c.topic)

	return nil
}

// disconnect
func (c *connection) disconnect() {
	// disconnect the client
	c.client.Disconnect(250)
	fmt.Println("Disconnected from Mosquitto broker")

}

// 4 - main
func run() {
	var con connection
	con.connectBroker()

	con.subscribeTopic()

	con.publishMsg("", "mensagem")

	time.Sleep(5 * time.Second)

	con.disconnect()

}

func main() {

	run()
}
