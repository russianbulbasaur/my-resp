package encoder

import (
	"strconv"
	"strings"
)

type MyRespEncoder struct {
}

func (encoder *MyRespEncoder) encodeSimpleString(input string) []byte {
	input += "\r\n"
	response := "+"
	response += input
	return []byte(response)
}

func (encoder *MyRespEncoder) encodeSimpleError(error string) []byte {
	error += "\r\n"
	response := "-"
	response += error
	return []byte(response)
}

func (encoder *MyRespEncoder) encodeInteger(input int) []byte {
	response := ":"
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

func (encoder *MyRespEncoder) encodeBulkString(input string) []byte {
	input += "\r\n"
	response := "$"
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

func (encoder *MyRespEncoder) encodeBoolean(input bool) []byte {
	response := "#"
	if input {
		response += "t"
	} else {
		response += "f"
	}
	response += "\r\n"
	return []byte(response)
}
