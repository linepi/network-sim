package application

import (
	"net"
	"log"
)

type Node struct {
	Conn net.Conn
}

func (N *Node) Send(data []byte) (n int, err error) {
	return N.Conn.Write(data)
}

func (N *Node) Recv(buffer []byte) (n int, err error) {
	return N.Conn.Read(buffer)
}

func (N *Node) Connect(address string) {
	c, err := net.Dial("tcp4", address)
	if err != nil {
		log.Fatalln(err)
	}
	N.Conn = c
}

// wait until connection
func (N *Node) Serve(port string) {
	l, err := net.Listen("tcp4", "localhost:"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	c, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	N.Conn = c
}
