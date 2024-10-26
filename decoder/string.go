package decoder

import (
	"bytes"
	"github.com/russianbulbasaur/my-resp/constants"
	"log"
	"strconv"
)

func decodeSimpleString(input []byte) (string, string) {
	if !isValid(input) {
		log.Println("simple string : invalid lf and cr bytes")
		return "", ""
	}

	return string(input[1 : len(input)-2]), constants.PurpleString
}

func decodeBulkString(input []byte) (string, string) {
	if !isValid(input) {
		log.Println("invalid lf and cr bytes")
		return "", ""
	}
	lengthAndPayload := bytes.Split(input[1:], []byte("\r\n"))
	if len(lengthAndPayload) < 2 {
		log.Println("Invalid bulk string")
		return "", ""
	}
	bulkLength, err := strconv.Atoi(string(lengthAndPayload[0]))
	if err != nil {
		log.Println(err)
		return "", ""
	}
	payload := string(lengthAndPayload[1])
	if len(payload) != bulkLength {
		log.Println("Payload length and specified length do not match")
		return "", ""
	}
	return payload, constants.PurpleString
}
