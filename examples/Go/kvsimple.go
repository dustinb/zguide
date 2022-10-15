// Clone pattern example
//
// Author: Dustin Butler <roundporch@gmail.com>
package kvsimple

import (
	"log"

	"github.com/go-zeromq/zmq4"
	"github.com/vmihailenco/msgpack/v5"
)

type KVMsg struct {
	Key     int64
	Body    int64
	SubType string
}

func (kv *KVMsg) Frames() [][]byte {
	b, _ := msgpack.Marshal(kv)
	frames := make([][]byte, 3)
	frames[0] = []byte(kv.SubType)
	frames[1] = b
	return frames
}

func (kv *KVMsg) Send(pub zmq4.Socket) {
	zmsg := zmq4.Msg{}
	zmsg.Frames = kv.Frames()
	log.Printf("Sending %v %v", kv, zmsg)
	pub.Send(zmsg)
}
