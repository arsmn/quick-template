package encoder

type EncoderFormat string

const (
	Unknown EncoderFormat = ""
	JSON    EncoderFormat = "json"
	MsgPack EncoderFormat = "msgpack"
)

type Encoder[T any] interface {
	Encode(v T) ([]byte, error)
	Decode(data []byte) (T, error)
}

func GetEncoder[T any](format EncoderFormat) Encoder[T] {
	switch format {
	case JSON:
		return JSONEncoder[T]{}
	case MsgPack:
		return MsgPackEncoder[T]{}
	default:
		return NopEncoder[T]{}
	}
}
