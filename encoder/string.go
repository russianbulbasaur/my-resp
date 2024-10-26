package encoder

import (
	"github.com/russianbulbasaur/my-resp/constants"
	"strconv"
)

func (encoder *MyRespEncoder) EncodeBulkString(input string) []byte {
	input += "\r\n"
	response := constants.BulkStringPrefix
	stringLen := len(input)
	if stringLen == 0 {
		//null string
		response += "-1"
		response += input
		return []byte(response)
	}
	response += strconv.Itoa(stringLen)
	response += "\r\n"
	response += input
	return []byte(response)
}

func (encoder *MyRespEncoder) EncodeSimpleString(input string) []byte {
	input += "\r\n"
	response := constants.SimpleStringPrefix
	response += input
	return []byte(response)
}
