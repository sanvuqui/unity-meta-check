package meta

import (
	"crypto/rand"
	"hash/fnv"
	mathrand "math/rand"
)

type GUIDGen func() (*GUID, error)
type GUIDGenByName func(name string) (*GUID, error)

func RandomGUIDGenerator() GUIDGen {
	return func() (*GUID, error) {
		bytes := make([]byte, GUIDByteLength)
		_, err := rand.Read(bytes)
		if err != nil {
			return nil, err
		}
		return NewGUID(bytes)
	}
}

func GUIDGeneratorByName() GUIDGenByName {
	return func(name string) (*GUID, error) {
		h := fnv.New32()
		_, err := h.Write([]byte(name))
		if err != nil {
			return nil, err
		}
		r := mathrand.New(mathrand.NewSource(int64(h.Sum32())))

		bytes := make([]byte, GUIDByteLength)
		_, err = r.Read(bytes)
		if err != nil {
			return nil, err
		}

		return NewGUID(bytes)
	}
}
