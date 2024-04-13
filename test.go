package main

import (
	"cn/application"
	"fmt"
	_ "io"
	"log"
	"math/rand"
	_ "net"
	_ "os"
	"sync"
	"time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	loopnum = 1000
	stringsize = 10000
)

func RandBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	begin := time.Now()
	checkv := make(chan string)
	recvdone := make(chan bool)

	go func() {
		var server application.Node
		server.Serve(CONN_PORT)
		for i := 0; i < loopnum; i++ {
			bytes := RandBytes(stringsize)
			checkv <- string(bytes)
			acc := 0
			for acc < stringsize {
				sendn, err := server.Send(bytes)
				if err != nil {
					log.Fatalln(err)
				}
				acc += sendn
			}
			<-recvdone // wait until last recv done
		}
		wg.Done()
	} ()

	go func() {
		time.Sleep(10 * time.Millisecond)
		var client application.Node
		client.Connect(CONN_HOST + ":" + CONN_PORT)
		buffer := make([]byte, stringsize)
		for i := 0; i < loopnum; i++ {
			checks := <-checkv // wait until randbytes created
			shouldrecvn := len(checks)
			accstring := ""
			acc := 0
			for acc < shouldrecvn {
				recvn, err := client.Recv(buffer)
				if err != nil {
					log.Fatalln(err)
				}
				accstring += string(buffer[0:recvn])
				acc += recvn
			}
			if acc != shouldrecvn {
				log.Fatalln("panic")
			}
			if checks != accstring {
				log.Fatalf("Not same %s != %s\n", checks, accstring)
			}
			recvdone <- true
		}
		wg.Done()
	} ()

	wg.Wait()
	end := time.Now()
	fmt.Printf("%fs\n", end.Sub(begin).Seconds())
}