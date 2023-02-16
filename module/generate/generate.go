package generate

import (
	"bytes"
	"encoding/base32"

	"github.com/google/uuid"
)

var (
	encoding = base32.NewEncoding("xGp46q4HEE0zqdV18t27Tk7j4tXOZK2I")
)

func RandomBytes() ([]byte, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return []byte(id[:]), nil
}

func UUID() (string, error) {
	rnd, err := RandomBytes()
	if err != nil {
		return "", err
	}

	var b bytes.Buffer

	encoder := base32.NewEncoder(encoding, &b)
	if _, err := encoder.Write(rnd); err != nil {
		return "", err
	}

	if err := encoder.Close(); err != nil {
		return "", err
	}

	b.Truncate(26)

	return b.String(), nil
}
