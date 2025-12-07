package id

import "github.com/google/uuid"

type UUIDGenerator struct{}

func (g *UUIDGenerator) Generate() string {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return id.String()
}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}