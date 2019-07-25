package verify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerify_Valid(t *testing.T) {
	type Department struct {
		ID   int    `valid:"Required"`
		Name string `valid:"Required;Length(20)"`
	}
	type User struct {
		ID                    int
		TestRequiredAndLength string `valid:"Required;Length(20)"`
		TestRangeRange        int64  `valid:"Range(20, 80)"`
		TestMax               int32  `valid:"Max(20)"`
		TestMin               int    `valid:"Min(20)"`
		Dept                  Department
		DeptRequired          Department `valid:"Required"`
	}
	u := &User{
		ID:                    1001,
		TestRequiredAndLength: "TestRequiredAndLength",
		TestRangeRange:        25,
		TestMax:               20,
		TestMin:               20,
		Dept:                  Department{ID: 1, Name: "DEPT"},
		DeptRequired:          Department{ID: 1, Name: "DEPT"},
	}
	verify := Verify{}
	b, err := verify.Valid(u)
	assert.Nil(t, err)
	assert.EqualValues(t, b, true)
}
