package encoder

import "github.com/russianbulbasaur/my-resp/constants"

func (encoder *MyRespEncoder) EncodeSimpleError(error string) []byte {
	error += "\r\n"
	response := constants.ErrorPrefix
	response += error
	return []byte(response)
}
