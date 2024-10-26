package encoder

import (
	"my_resp/constants"
	"strconv"
	"strings"
)

func (encoder *MyRespEncoder) EncodeInteger(input int) []byte {
	response := constants.IntegerPrefix
	inputString := ""
	if input < 0 {
		inputString = strings.Trim(strconv.Itoa(input), "-")
		response += "-"
	} else {
		inputString = strings.Trim(strconv.Itoa(input), "+")
		response += "+"
	}
	response += inputString
	return []byte(response)
}
