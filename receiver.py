#!/usr/bin/python


import zmq
import sys


END_MSG = "This is the end message that has to be send to stop the benmark"
OK_MSG = "OK"


def main():
    ctx = zmq.Context()
    zock_in = ctx.socket(zmq.PAIR)
    zock_out = ctx.socket(zmq.PAIR)

    zock_in.connect(sys.argv[1])
    zock_out.bind(sys.argv[2])

    while True:
        try:
            msg = zock_in.recv()
            if msg == END_MSG:
                break
        except KeyboardInterrupt:
            break
    zock_out.send(OK_MSG)


if __name__ == "__main__":
    main()
