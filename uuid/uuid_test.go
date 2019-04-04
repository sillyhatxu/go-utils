package uuid

import (
	"github.com/stretchr/testify/assert"
	"go-utils/hashset"
	"testing"
)

func TestShortId(t *testing.T) {
	idSet := hashset.New()
	for i := 0; i < 100000; i++ {
		id := ShortId()
		//fmt.Println(id)
		idSet.Add(id)
	}
	assert.EqualValues(t, idSet.Size(), 100000)

}
