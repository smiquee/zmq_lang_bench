#!/usr/bin/python


import zmq
import sys


def main():
    ctx = zmq.Context()
    zock_in = ctx.socket(zmq.PAIR)
    zock_out = ctx.socket(zmq.PAIR)

    zock_in.connect(sys.argv[1])
    zock_out.bind(sys.argv[2])

    while True:
        try:
            zock_out.send(zock_in.recv())
        except KeyboardInterrupt:
            break


if __name__ == "__main__":
    main()
