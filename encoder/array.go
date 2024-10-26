package encoder

import (
	"bytes"
	"fmt"
	"my_resp/constants"
)

func (encoder *MyRespEncoder) EncodeBooleanArray(array []bool) []byte {
	if len(array) == 0 {
		return emptyArray()
	}
	response := bytes.NewBuffer(make([]byte, 0))
	response.WriteString(fmt.Sprintf("%s%d\r\n", constants.ArrayPrefix, len(array)))
	for _, element := range array {
		response.Write(encoder.EncodeBoolean(element))
	}
	return response.Bytes()
}

func (encoder *MyRespEncoder) EncodeIntegerArray(array []int) []byte {
	if len(array) == 0 {
		return emptyArray()
	}
	response := bytes.NewBuffer(make([]byte, 0))
	response.WriteString(fmt.Sprintf("%s%d\r\n", constants.ArrayPrefix, len(array)))
	for _, element := range array {
		response.Write(encoder.EncodeInteger(element))
	}
	return response.Bytes()
}

func (encoder *MyRespEncoder) EncodeStringArray(array []string) []byte {
	if len(array) == 0 {
		return emptyArray()
	}
	response := bytes.NewBuffer(make([]byte, 0))
	response.WriteString(fmt.Sprintf("%s%d\r\n", constants.ArrayPrefix, len(array)))
	for _, element := range array {
		response.Write(encoder.EncodeSimpleString(element))
	}
	return response.Bytes()
}

func emptyArray() []byte {
	return []byte("*0\r\n")
}