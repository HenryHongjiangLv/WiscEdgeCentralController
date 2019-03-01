package heartbeat_client

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

type Heartbeat struct {
	serverAddress  string
	clientAddress string
	interval time.Duration
	timeOut  time.Duration
	message  HeartbeatMessage
	stopCh   chan struct{}
}

type HeartbeatMessage struct {
	Message string
}


func NewHeartbeat(
	host string, port uint16, interval time.Duration, timeOut time.Duration, message HeartbeatMessage,
) (*Heartbeat, error) {

	address := fmt.Sprintf("%s:%d", host, port)
	stopCh := make(chan struct{})

	return &Heartbeat{
		address, address, interval, timeOut, message, stopCh,
	}, nil
}

func (hb *Heartbeat)StopHeartbeat(struct{}) {
	close(hb.stopCh)
}

func (hb *Heartbeat) StartHeartbeat() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(hb.serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	for {
		select {
		case m := <- hb.stopCh:
			hb.StopHeartbeat(m)
			return
		case <-time.After(hb.interval):
			ctx, cancel := context.WithTimeout(context.Background(), hb.timeOut)

			r, err := c.ReceiveAndReply(ctx, &pb.HelloRequest{Name: hb.message.Message})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", r.Message)
			cancel()
		}

	}
}


