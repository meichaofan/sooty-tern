package util

import (
	"github.com/google/uuid"
)

func MustUUID() string {
	u, err := newUUID()
	if err != nil {
		panic(err)
	}
	return u
}

func newUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
