//go:build go1.18
// +build go1.18

package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	as := assert.New(t)

	v0 := false
	v1 := true

	as.Equal(&v0, Ptr(v0))
	as.Equal(&v1, Bool(v1))
}
