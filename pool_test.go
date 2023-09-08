package pool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidOpts(t *testing.T) {
	assert.Panics(t, func() { New[int](0, 10) })
	assert.Panics(t, func() { New[int](10, 0) })
	assert.Panics(t, func() { New[int](0, 0) })
}

func TestIncrementWorkload(t *testing.T) {}

func TestPanicWorkload(t *testing.T) {}

func TestFilledQueue(t *testing.T) {}

func TestGracefulShutdown(t *testing.T) {}

func TestDoubleFree(t *testing.T) {}

func TestEnqueueAfterFree(t *testing.T) {}
