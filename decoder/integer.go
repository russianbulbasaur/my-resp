package decoder

import (
	"github.com/russianbulbasaur/my-resp/constants"
	"log"
	"strconv"
)

func decodeInteger(input []byte) (int, string) {
	if !isValid(input) {
		log.Println("Invalid cr lf bytes")
		return -1, ""
	}
	sign := input[1]
	if sign != constants.PositiveSign && sign != constants.NegativeSign {
		log.Println("No sign")
		return -1, ""
	}
	parsedInteger, err := strconv.Atoi(string(input[2:(len(input) - 2)]))
	if err != nil {
		log.Println("Cannot parse the integer value")
		return -1, ""
	}
	return parsedInteger, constants.PurpleInteger
}
