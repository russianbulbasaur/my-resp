package decoder

import (
	"github.com/russianbulbasaur/my-resp/constants"
	"log"
)

func decodeBoolean(input []byte) (bool, string) {
	if !isValid(input) {
		log.Println("Invalid cr lf bytes")
		return false, ""
	}
	respBool := input[1]
	var response bool
	if respBool == constants.True {
		response = true
	} else if respBool == constants.False {
		response = false
	} else {
		panic("invalid boolean type")
	}
	return response, constants.PurpleBoolean
}
