package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
)

func main() {
	ctx, _ := zmq.NewContext()
	defer ctx.Term()
	zock_in, _ := ctx.NewSocket(zmq.PAIR)
	zock_out, _ := ctx.NewSocket(zmq.PAIR)
	defer zock_in.Close()
	defer zock_out.Close()

	if len(os.Args) != 3 {
		fmt.Println("Give 2 arguments! 'url_socket_in url_socket_out'")
		os.Exit(1)
	}

	zock_in.Connect(os.Args[1])
	zock_out.Bind(os.Args[2])

	msg := []byte{}
	for {
		msg, _ = zock_in.RecvBytes(0)
		if string(msg) == "This is the end message that has to be send to stop the benmark" {
			break
		}
	}
	msg = []byte("OK")
	zock_out.SendBytes(msg, 0)
}
