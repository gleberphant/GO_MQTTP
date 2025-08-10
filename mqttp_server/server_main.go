/*/
 aplicação do servidor mqttp BROKER
 recebe todas mensagens
 distribuir todas mensagens
/*/

package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
)

type server struct {
	status   bool
	i        int
	server   *mqtt.Server
	listener *listeners.TCP
}

// 1 função para iniciar o servidor
func (s *server) startServer() error {

	// create new mqtt server
	s.server = mqtt.NewServer(nil)

	// create listener tcp
	s.listener = listeners.NewTCP("t1", ":1883")

	// add listener to server
	err := s.server.AddListener(s.listener, nil)

	if err != nil {
		log.Fatal(err)
	}

	//start dbroker

	err = s.server.Serve()

	if err != nil {
		log.Fatal(err)
	}

	s.status = true

	return nil
}

// 2 escutar mensagens
func (s *server) listen() error {

	fmt.Printf("Escutando clientes ..... [ %d ] \n", s.i)
	s.i++

	return nil

}

func (s *server) stopServer() error {
	s.status = false
	return nil
}
func run() {

	fmt.Printf("Iniciando servidor....")
	var srv server
	srv.startServer()

	for srv.status {
		srv.listen()
		time.Sleep(1000 * time.Millisecond)

		if srv.i > 10 {
			srv.stopServer()
		}
	}

}

func main() {
	run()
}
