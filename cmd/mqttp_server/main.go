/*/
 aplicação do servidor mqttp BROKER
 recebe todas mensagens
 distribuir todas mensagens
/*/

package main

import (
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"

	mqtt_server "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/events"
	"github.com/mochi-co/mqtt/server/listeners"
)

type server struct {
	status   bool
	server   *mqtt_server.Server
	listener *listeners.TCP
}

// 1 função para iniciar o servidor
func (s *server) startServer() error {

	// create new mqtt server
	s.server = mqtt_server.NewServer(nil)

	// add listener to server
	fmt.Println(aurora.BgBlue("Escutando TCP1 : 1883"))
	if err := s.server.AddListener(listeners.NewTCP("tcp1", ":1883"), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println(aurora.BgBlue("Escutando WebSocker: 1884 ...."))
	if err := s.server.AddListener(listeners.NewWebsocket("ws1", ":1884"), nil); err != nil {
		log.Fatal(err)
	}
	//start broker

	if err := s.server.Serve(); err != nil {
		log.Fatal(err)
	}

	s.server.Events.OnConnect = func(client events.Client, pkg events.Packet) {
		fmt.Printf("\n [%s] Conectou ao servidor", client.ID)
	}

	s.server.Events.OnDisconnect = func(client events.Client, err error) {
		fmt.Printf("\n [%s] Desconectou do servidor", client.ID)
	}

	s.server.Events.OnSubscribe = func(filter string, client events.Client, qos byte) {
		fmt.Printf("\n [%s] fez inscrição no topico [%s]", client.ID, filter)
	}

	s.server.Events.OnMessage = func(client events.Client, msg events.Packet) (pkx events.Packet, err error) {

		payloadFormatado := fmt.Sprintf("%s disse : %s ", string(client.ID), string(msg.Payload))

		fmt.Printf("\n [%s] - [%s]\n", msg.TopicName, payloadFormatado)

		msg.Payload = []byte(payloadFormatado)

		return msg, nil
	}

	s.status = true

	return nil
}

func (s *server) stopServer() error {
	s.status = false
	fmt.Printf("Encerrando servidor.... \n")
	return nil
}
func run() {

	var srv server
	fmt.Println(aurora.BgBlue("Iniciando Servidor"))
	srv.startServer()
	defer srv.stopServer()

	fmt.Println(aurora.BgBlue("Escutando clientes"))
	select {}
}

func main() {
	run()
}
