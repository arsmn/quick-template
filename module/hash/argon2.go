package hash

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2Config struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var DefaultArgon2Config = Argon2Config{
	Memory:      64 * 1024,
	Iterations:  1,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

type Argon2 struct {
	config Argon2Config
}

func NewArgon2(cfg Argon2Config) *Argon2 {
	return &Argon2{cfg}
}

func (h *Argon2) Hash(ctx context.Context, data []byte) ([]byte, error) {
	salt := make([]byte, h.config.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	hash := argon2.IDKey(
		[]byte(data),
		salt,
		h.config.Iterations,
		h.config.Memory,
		h.config.Parallelism,
		h.config.KeyLength,
	)

	return h.encodeHash(h.config, salt, hash)
}

func (h *Argon2) Compare(ctx context.Context, data []byte, hash []byte) error {
	cfg, salt, hash, err := h.decodeHash(string(hash))
	if err != nil {
		return err
	}

	otherHash := argon2.IDKey([]byte(data), salt, cfg.Iterations, cfg.Memory, cfg.Parallelism, cfg.KeyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return nil
	}

	return ErrMismatchedHashAndData
}

func (h *Argon2) encodeHash(cfg Argon2Config, salt, hash []byte) ([]byte, error) {
	var b bytes.Buffer

	if _, err := fmt.Fprintf(
		&b,
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, h.config.Memory, h.config.Iterations, h.config.Parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (h *Argon2) decodeHash(encodedHash string) (cfg Argon2Config, salt, hash []byte, err error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return cfg, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return cfg, nil, nil, err
	}

	if version != argon2.Version {
		return cfg, nil, nil, fmt.Errorf("argon2 incompatible version. [hash: %d] [current: %d]", version, argon2.Version)
	}

	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &cfg.Memory, &cfg.Iterations, &cfg.Parallelism)
	if err != nil {
		return cfg, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return cfg, nil, nil, err
	}

	cfg.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return cfg, nil, nil, err
	}

	cfg.KeyLength = uint32(len(hash))

	return cfg, salt, hash, nil
}
