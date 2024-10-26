package encoder

import "github.com/russianbulbasaur/my-resp/constants"

func (encoder *MyRespEncoder) EncodeBoolean(input bool) []byte {
	response := constants.BooleanPrefix
	if input {
		response += "t"
	} else {
		response += "f"
	}
	response += "\r\n"
	return []byte(response)
}
