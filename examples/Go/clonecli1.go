// Clone pattern example
//
// Author: Dustin Butler <roundporch@gmail.com>
package main

import (
	"context"
	"log"
	"sync"

	kvsimple "myproject/lib"

	"github.com/go-zeromq/zmq4"
	"github.com/vmihailenco/msgpack/v5"
)

// Sort KVMsg by Key
var kvMap sync.Map

func main() {

	sub := zmq4.NewSub(context.Background())
	defer sub.Close()

	sub.Dial("tcp://localhost:5556")
	sub.SetOption(zmq4.OptionSubscribe, "tx")

	for {
		msg, err := sub.Recv()
		if err != nil {
			log.Fatalf("could not receive message: %v", err)
		}

		// Handling messages in a go routine is safe with sync.Map
		go func() {
			msgType := string(msg.Frames[0])
			b := []byte(msg.Frames[1])
			kvmsg := kvsimple.KVMsg{}
			msgpack.Unmarshal(b, &kvmsg)
			kvMap.Store(kvmsg.Key, kvmsg)
			log.Printf("Received item %s %v", msgType, kvmsg)
		}()
	}
}
