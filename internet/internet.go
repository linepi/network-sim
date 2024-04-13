package internet
import (
	"net"
	"log"
)

type Route struct {
	/* use tcp, which is stable and accurate, 
	   to simulate internet layer, which is unexpectable */
	conn net.Conn 
}

// unreliable data transfer, won't return anything as the reality is
func (r *Route) UdtSend(data []byte) {

}
func (r *Route) UdtRecv(buffer []byte) {

}

func (r *Route) Connect(address string) {
	c, err := net.Dial("tcp4", address)
	if err != nil {
		log.Fatalln(err)
	}
	r.conn = c
}

func (r *Route) Serve(port string) {
	l, err := net.Listen("tcp4", "localhost:"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()
	c, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	r.conn = c
}