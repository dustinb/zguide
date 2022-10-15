// Clone pattern example
//
// Author: Dustin Butler <roundporch@gmail.com>
package main

import (
	"context"
	"math/rand"
	"time"

	kvsimple "myproject/lib"

	"github.com/go-zeromq/zmq4"
)

func main() {
	pub := zmq4.NewPub(context.Background())
	defer pub.Close()
	pub.Listen("tcp://localhost:5556")

	for {
		item := kvsimple.KVMsg{
			Key:     rand.Int63(),
			Body:    rand.Int63(),
			SubType: "tx",
		}
		item.Send(pub)

		// Example client is not subscribed to block messages
		item = kvsimple.KVMsg{
			Key:     rand.Int63(),
			Body:    rand.Int63(),
			SubType: "block",
		}
		item.Send(pub)

		time.Sleep(time.Second * 3)
	}
}
