package encoder

type MyRespEncoder struct {
}

func (encoder *MyRespEncoder) encode(string) ([]byte, error) {
	return nil, nil
}
