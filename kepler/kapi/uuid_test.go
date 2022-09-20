package kapi

import (
	"testing"

	"github.com/google/uuid"
)

func TestUuid(t *testing.T) {
	t.Logf("%+v", uuid.NewString())
}
