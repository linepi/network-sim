package transport
import (
	"net"
	"log"
)

type TCP struct {
	conn net.Conn
}

func (tcp *TCP) Write(data []byte) (n int, err error) {
	return tcp.conn.Write(data)
}

func (tcp *TCP) Read(buffer []byte) (n int, err error) {
	return tcp.conn.Read(buffer)
}

func (tcp *TCP) Connect(address string) {
	c, err := net.Dial("tcp4", address)
	if err != nil {
		log.Fatalln(err)
	}
	tcp.conn = c
}

func (tcp *TCP) Serve(port string) {
	l, err := net.Listen("tcp4", "localhost:"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	c, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	tcp.conn = c
}