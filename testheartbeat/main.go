package main

import (
	"time"

	"github.com/WiscEdgeCentralController/heartbeat_client"
)


func main() {

	address     := "localhost"
	var port uint16
	port = 50050
	DefaultName := "world"
	timeOut := 10 * time.Second
	interval := 1 * time.Second

	hb, err := heartbeat_client.NewHeartbeat(
		address, port, interval, timeOut, heartbeat_client.HeartbeatMessage{DefaultName},
	)

	if err != nil {
		return
	}

	hb.StartHeartbeat()
}