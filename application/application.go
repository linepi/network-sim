package application

import (
	"cn/transport"
)

type Node struct {
	T transport.TransportProtocol
}

func (N *Node) Send(data []byte) (n int, err error) {
	return N.T.Write(data)
}

func (N *Node) Recv(buffer []byte) (n int, err error) {
	return N.T.Read(buffer)
}

func (N *Node) Connect(address string) {
	N.T.Connect(address)
}

// wait until connection
func (N *Node) Serve(port string) {
	N.T.Serve(port)
}

func NewTCPNode() *Node {
	node := Node{}
	node.T = &transport.TCP{}
	return &node
}

func NewGBNNode() *Node {
	node := Node{}
	node.T = &transport.TCP{}
	return &node
}