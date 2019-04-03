package uuid

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func UUID() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
