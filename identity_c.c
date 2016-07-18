#include <zmq.h>
#include <string.h>
#include <stdio.h>

int main(int argc, char* argv[]) {

  if(argc != 3){
    printf("%s takes 2 arguments: url_socket_in url_socket_out", argv[0]);
    return 1;
  }
  void * ctx = zmq_ctx_new() ;

  void * zock_in = zmq_socket(ctx, ZMQ_PAIR);
  void * zock_out = zmq_socket(ctx, ZMQ_PAIR);

  zmq_connect(zock_in, argv[1]);
  zmq_bind(zock_out, argv[2]);

  char msg[128] ;
  int nb = -1;

  while(1) {
    nb = zmq_recv(zock_in, msg, 128, 0);
    if( nb == -1){break;}
    msg[nb] = '\0';
    nb = zmq_send(zock_out, msg, strlen(msg), 0);
    if( nb == -1){break;}
    }
  return 0;
}
