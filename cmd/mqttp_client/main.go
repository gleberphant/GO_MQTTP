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
	status bool
	client mqtt.Client
}

// 1 - função para conectar ao servidor MQTT

func (c *connection) connectBroker() error {

	// config client
	clientOptions := mqtt.NewClientOptions()
	clientOptions.AddBroker("ws://localhost:1884")
	clientOptions.SetUsername("Cliente01")
	clientOptions.DefaultPublishHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("\n [ %s] - %s \n", msg.Topic(), string(msg.Payload()))
	}

	// create and connect the client
	c.client = mqtt.NewClient(clientOptions)

	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect: %v", token.Error())
	}

	fmt.Println("Connected to BROKER  Server")
	c.status = true
	return nil
}

// - subscrever um tópico
func (c *connection) subscribeTopic(topic string) error {

	if token := c.client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to subscribe: %v", token.Error())
	}

	fmt.Printf(" *** Subscribed to topic : %s *** ", topic)

	return nil
}

// 2 - publicar mensagem em um topico
func (c *connection) publishMsg(topic string, payload string) error {

	if token := c.client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
		log.Printf("Failed to publish : %v", token.Error())
	}

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

	// iniciar buf de scan do teclado
	scanner := bufio.NewScanner(os.Stdin)

	//conectar ao servidor broker
	con.connectBroker()
	defer con.disconnect()

	//subescrever em um topico
	con.subscribeTopic("testes")

	//loop de envio de mensagens
	for {
		fmt.Print("\n Digita a mensagem :")

		scanner.Scan()

		if err := scanner.Err(); err != nil {
			log.Fatal("erro de leitura do teclado")
		}

		payload = scanner.Text()

		if payload == "sair" || payload == "" {
			break
		}

		con.publishMsg("testes", payload)

	}

}

func main() {

	run()
}
