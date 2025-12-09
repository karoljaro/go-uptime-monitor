package id

import (
	"testing"
	"github.com/google/uuid"
)

func TestUUIDGeneratorGenerate(t *testing.T) {
	gen := NewUUIDGenerator()

	u1 := gen.Generate()
	u2 := gen.Generate()

	parsed1, err := uuid.Parse(u1)
	if err != nil {
		t.Fatalf("Generate() returned invalid UUID: %v", err)
	}

	if parsed1.Version() != 7 {
		t.Fatalf("expected UUID version 7, got %d", parsed1.Version())
	}

	if u1 == u2 {
		t.Fatalf("expected two different UUIDs, but got %q and %q", u1, u2)
	}
}
