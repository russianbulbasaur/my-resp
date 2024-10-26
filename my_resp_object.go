package my_resp

import (
	"github.com/russianbulbasaur/my-resp/decoder"
	"github.com/russianbulbasaur/my-resp/encoder"
)

type MyRespObject struct {
	E *encoder.MyRespEncoder
	D *decoder.MyRespDecoder
}

func (ob *MyRespObject) newEncoder() {
	ob.E = &encoder.MyRespEncoder{}
}

func (ob *MyRespObject) newDecoder() {
	ob.D = &decoder.MyRespDecoder{}
}
