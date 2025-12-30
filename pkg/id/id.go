package id

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() (ID, error) {
	u, err := uuid.NewRandom()
	return ID(u), err
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
