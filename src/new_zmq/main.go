package main

// #cgo LDFLAGS: -lzmq
// #include <zmq.h>
import "C"

import "unsafe"

func main() {

	ctx := C.zmq_ctx_new()
	defer C.zmq_ctx_destroy(ctx)
	defer C.zmq_ctx_shutdown(ctx)
	var rc C.int

	var sin *C.char = C.CString("ipc:///tmp/send_out")
	zock_in := C.zmq_socket(ctx, C.ZMQ_PAIR)
	rc = C.zmq_connect(zock_in, sin)
	defer C.zmq_close(zock_in)

	var sout *C.char = C.CString("ipc:///tmp/recv_in")
	zock_out := C.zmq_socket(ctx, C.ZMQ_PAIR)
	rc = C.zmq_bind(zock_out, sout)

	defer C.zmq_close(zock_out)

	var len C.size_t = 128
	var recv C.size_t
	var msg [128]byte

	for {
		rc = C.zmq_recv(zock_in, unsafe.Pointer(&msg[0]), len, 0)
		recv = C.size_t(rc)
		rc = C.zmq_send(zock_out, unsafe.Pointer(&msg[0]), recv, 0)
	}

}
