package transport

import (
	"cn/internet"
	"crypto/sha1"
	"strconv"
)

const N = 10

// transport GBN package
type pack struct {
	seqnum int
	data []byte
	checksum [sha1.Size]byte
}

type receiver struct {
	expectedseqnum int	
}

type sender struct {
	base int
	nextseqnum int
	sndpkt []pack
}

type GBN struct {
	route internet.Route
	sender sender
	receiver receiver
}

func (gbn *GBN) Write(data []byte) (n int, err error) {
	sd := &gbn.sender
	if sd.nextseqnum < sd.base + N {
		return 0, nil
	}
	// compute the checksum of pack
	checksum := sha1.Sum(append([]byte(strconv.Itoa(sd.nextseqnum)), data...))
	sd.sndpkt[sd.nextseqnum - sd.base] = pack{sd.nextseqnum, data, checksum}
	return 0, nil
}

func (gbn *GBN) Read(buffer []byte) (n int, err error) {
	return 0, nil
}

func (gbn *GBN) init() {
	// sender init
	gbn.sender.base = 1
	gbn.sender.nextseqnum = 1
	gbn.sender.sndpkt = make([]pack, N)
	// receiver init
	gbn.receiver.expectedseqnum = 1
}

func (gbn *GBN) Connect(address string) {
	gbn.init() // init when connect to server
	gbn.route.Connect(address)
}

func (gbn *GBN) Serve(port string) {
	gbn.init() // init when open a port
	gbn.route.Serve(port)
}
