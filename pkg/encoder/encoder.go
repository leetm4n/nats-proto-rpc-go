package encoder

type Encoder interface {
	Encode(subject string, v interface{}) ([]byte, error)
	Decode(subject string, data []byte, vPtr interface{}) error
}
