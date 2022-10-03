package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysPassing(t *testing.T) {
	t.Parallel()
	assert.True(t, true)
}
