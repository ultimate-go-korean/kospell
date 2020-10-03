package kospell

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCheck tests e2e
func TestCheck(t *testing.T) {
	out, err := Check("아버지가방에 들어가신다.")
	assert.NoError(t, err)
	assert.Equal(t, "아버지가 방에 들어가신다", out[0].ErrInfo[0].CandWord)
}
