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
	"bufio"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type connection struct {
	status             bool
	broker_url         string
	client             mqtt.Client
	publishHandler     mqtt.MessageHandler
	connectionHandler  mqtt.OnConnectHandler
	connectLostHandler mqtt.ConnectionLostHandler
}

// PUBLISHED handler

// 1 - função para conectar ao servidor MQTT

func (c *connection) connectBroker() error {

	// config client
	c.broker_url = "tcp://localhost:1883"

	c.publishHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("\n Received: Topico %s - Mensagem: %s \n", msg.Topic(), msg.Payload())
	}

	c.connectionHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}

	c.connectLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connection LOST: %v", err)
	}

	// create and connect the client
	c.client = mqtt.NewClient(mqtt.NewClientOptions().AddBroker(c.broker_url).SetDefaultPublishHandler(c.publishHandler))

	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect: %v", token.Error())
	}

	fmt.Println("Connected to BROKER  Server")
	c.status = true
	return nil
}

// 2 - publicar mensagem
func (c *connection) publishMsg(topic string, payload string) error {

	if token := c.client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		log.Printf("Failed to publish : %v", token.Error())
	}

	return nil
}

// 3 - receber mensagens de um topico
func (c *connection) subscribeTopic(topic string) error {
	// subscribe operations

	if token := c.client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {

		log.Fatalf("Failed to subscribe: %v", token.Error())
	}

	fmt.Printf(" *** Subscribed to topic : %s *** ", topic)

	return nil
}

// disconnect
func (c *connection) disconnect() {
	// disconnect the client
	c.client.Disconnect(50)
	fmt.Println("Disconnected from broker")

}

// 4 - main
func run() {
	var con connection
	var payload string
	scanner := bufio.NewScanner(os.Stdin)

	con.connectBroker()
	defer con.disconnect()

	con.subscribeTopic("testes")

	for {
		fmt.Print("\n Digita a mensagem :")

		scanner.Scan()

		if err := scanner.Err(); err != nil {
			log.Fatal("erro de leitura do teclado")
		}

		payload = scanner.Text()

		if payload == "sair" {
			break
		}

		con.publishMsg("testes", payload)

	}

}

func main() {

	run()
}
