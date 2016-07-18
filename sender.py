#!/usr/bin/python


import zmq
import sys
import time

MSG = "This is a message that has to be transmitted across all the components"
END_MSG = "This is the end message that has to be send to stop the benmark"
OK_MSG = "OK"


def main():
    ctx = zmq.Context()
    zock_in = ctx.socket(zmq.PAIR)
    zock_out = ctx.socket(zmq.PAIR)

    zock_in.connect(sys.argv[1])
    zock_out.bind(sys.argv[2])

    try:
        nb_msg = int(sys.argv[3])
    except IndexError:
        nb_msg = 1000000

    start = time.time()
    for i in xrange(nb_msg):
        try:
            zock_out.send(MSG)
        except KeyboardInterrupt:
            break
    zock_out.send(END_MSG)

    nb = zock_in.poll(10000)
    if nb != 0:
        msg = zock_in.recv()
        if msg == OK_MSG:
            stop = time.time()

    print(stop - start)


if __name__ == "__main__":
    main()
