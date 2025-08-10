/*/
 aplicação do servidor mqttp BROKER
 recebe todas mensagens
 distribuir todas mensagens
/*/

package main

import (
	"fmt"
)

type server struct {
	status bool
}

// 1 função para iniciar o servidor
func (s *server) startServer() error {
	return nil
}

// 2 criar topico
func (s *server) createTopic() error {
	return nil
}

// 3 escutar mensagens
func (s *server) listen() error {
	return nil
}

func mainServer() {

	fmt.Printf("Iniciando servidor....")
	var srv server
	srv.startServer()
	srv.createTopic()
	srv.listen()

}
