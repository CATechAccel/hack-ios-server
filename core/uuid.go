package core

import "github.com/google/uuid"

type UUIDGenerator interface {
	Generate() (string, error)
}

type UUID struct{}

// NewUUIDGenerator は，application.UUIDGeneratorを返します．
func NewUUIDGenerator() UUIDGenerator {
	return &UUID{}
}

// Generate は，uuidを返します．
func (u *UUID) Generate() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
