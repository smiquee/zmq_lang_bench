package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"os"
	"time"
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

	msg := []byte("This is a message that has to be transmitted across all the components")
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		zock_out.SendBytes(msg, 0)
	}
	msg = []byte("This is the end message that has to be send to stop the benmark")
	zock_out.SendBytes(msg, 0)

	msg, _ = zock_in.RecvBytes(0)

	stop := time.Now()
	var duration time.Duration = stop.Sub(start)

	fmt.Printf("%.6fs\n", duration.Seconds())
}
