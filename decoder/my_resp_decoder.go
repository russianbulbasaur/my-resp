package decoder

type MyRespDecoder struct {
}

func (decoder *MyRespDecoder) decode([]byte) (string, error) {
	return "", nil
}
