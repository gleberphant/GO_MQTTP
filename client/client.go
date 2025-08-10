/*/
- aplicação do cliente.
- conectar em servidor
- subscreve em topicos
- posta mensagens
- recebe mensagens
/*/

package main

import (
	"fmt"
)

const (
	broker_url = "tcp://localhost:1883"
	clientID   = "go-mqtt-client"
	topic      = "iot"
)

type connection struct {
	status  bool
	topicId string
	msgs    string
}

// 1 - função para conectar ao servidor mqttp
func (c *connection) connectMQTTPServer() error {
	c.status = true
	return nil
}

// 2 - função para enviar mensagem

func (c *connection) sendMsg(topicId string, msg string) error {
	c.topicId = topicId
	c.msgs = msg
	return nil
}

// 3 - função para receber mensagens do topico
func (c *connection) receiveMsg() error {
	c.msgs = "mensagens "
	return nil
}

// 4 - main
func mainClient() {
	var con connection
	con.connectMQTTPServer()

	con.sendMsg("", "mensagem")

	con.receiveMsg()

	fmt.Printf("Topico: %s\n", con.topicId)
	fmt.Printf("Msgs: %s\n", con.msgs)

}
