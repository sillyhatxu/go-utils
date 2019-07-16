package gotime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatTimestamp(t *testing.T) {
	Time := FormatTimestamp(1438167001716)
	assert.NotNil(t, Time)
	fmt.Println(FormatLocation(Time))
}
