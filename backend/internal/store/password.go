package store

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type Password struct{}

func (Password) Hash(pw string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	// Argon2id params
	time := uint32(1)
	memory := uint32(64 * 1024) // 64MB
	threads := uint8(2)
	keyLen := uint32(32)

	hash := argon2.IDKey([]byte(pw), salt, time, memory, threads, keyLen)
	// format: argon2id$v=19$m=65536,t=1,p=2$<salt>$<hash>
	encoded := fmt.Sprintf("argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		memory, time, threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)
	return encoded, nil
}

func (Password) Verify(pw, encoded string) bool {
	var memory uint32
	var time uint32
	var threads uint8
	var saltB64, hashB64 string
	_, err := fmt.Sscanf(encoded, "argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		&memory, &time, &threads, &saltB64, &hashB64,
	)
	if err != nil { return false }

	salt, err := base64.RawStdEncoding.DecodeString(saltB64)
	if err != nil { return false }
	want, err := base64.RawStdEncoding.DecodeString(hashB64)
	if err != nil { return false }

	keyLen := uint32(len(want))
	got := argon2.IDKey([]byte(pw), salt, uint32(time), memory, threads, keyLen)

	// constant-time compare
	if len(got) != len(want) { return false }
	var v byte
	for i := range got { v |= got[i] ^ want[i] }
	return v == 0
}
