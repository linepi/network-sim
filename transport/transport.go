package transport

type TransportProtocol interface {
	Write(data []byte) (n int, err error)
	Read(buffer []byte) (n int, err error)
	Connect(address string)
	Serve(port string)
}