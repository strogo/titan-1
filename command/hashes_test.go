package command

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initHashes(t *testing.T, key string, n int) {
	args := []string{key}
	for i := n; i > 0; i-- {
		args = append(args, "init_field"+strconv.Itoa(i), "bar")
	}
	ctx := ContextTest("hmset", args...)
	Call(ctx)
	lines := ctxLines(ctx.Out)
	assert.Equal(t, "+OK", lines[0])
}

func clearHashes(t *testing.T, key string) {
	ctx := ContextTest("del", key)
	Call(ctx)
	lines := ctxLines(ctx.Out)
	assert.Equal(t, ":1", lines[0])
}

func setHashes(t *testing.T, args ...string) []string {
	ctx := ContextTest("hmset", args...)
	Call(ctx)
	return ctxLines(ctx.Out)
}

func TestHLen(t *testing.T) {
	// init
	key := "hash-hlen-key"
	initList(t, key, 3)

	// case 1
	ctx := ContextTest("hlen", key)
	Call(ctx)
	lines := ctxLines(ctx.Out)
	assert.Equal(t, ":3", lines[0])

	// case 2
	lines = setHashes(t, key, "a", "a", "b", "b")
	assert.Equal(t, "+OK", lines[0])
	ctx = ContextTest("hlen", key)
	Call(ctx)
	lines = ctxLines(ctx.Out)
	assert.Equal(t, ":5", lines[0])

	// case 3
	lines = setHashes(t, key, "c", "c", "c", "d")
	assert.Equal(t, "+OK", lines[0])
	ctx = ContextTest("hlen", key)
	Call(ctx)
	lines = ctxLines(ctx.Out)
	assert.Equal(t, ":5", lines[0])

	// end
	clearHashes(t, key)
}

func TestHDel(t *testing.T) {
	// init
	key := "hash-hlen-key"
	initList(t, key, 5)

	// case 1
	ctx := ContextTest("hdel", key, "1")
	Call(ctx)
	lines := ctxLines(ctx.Out)
	assert.Equal(t, ":1", lines[0])
	ctx = ContextTest("hlen", key)
	Call(ctx)
	lines = ctxLines(ctx.Out)
	assert.Equal(t, ":4", lines[0])

	// case 2
	ctx = ContextTest("hdel", key, "2")
	Call(ctx)
	lines = ctxLines(ctx.Out)
	assert.Equal(t, ":1", lines[0])
	ctx = ContextTest("hlen", key)
	Call(ctx)
	lines = ctxLines(ctx.Out)
	assert.Equal(t, ":3", lines[0])

	// case 3
	ctx = ContextTest("hdel", key, "3", "4", "5")
	Call(ctx)
	lines = ctxLines(ctx.Out)
	assert.Equal(t, ":3", lines[0])
	ctx = ContextTest("hlen", key)
	Call(ctx)
	lines = ctxLines(ctx.Out)
	assert.Equal(t, ":0", lines[0])
	// then re-insert into hash
	lines = setHashes(t, key, "a", "a", "b", "b")
	assert.Equal(t, "+OK", lines[0])
	ctx = ContextTest("hlen", key)
	Call(ctx)
	lines = ctxLines(ctx.Out)
	assert.Equal(t, ":2", lines[0])

	// end
	clearHashes(t, key)
}

func TestHExists(t *testing.T)      {}
func TestHGet(t *testing.T)         {}
func TestHGetAll(t *testing.T)      {}
func TestHIncrBy(t *testing.T)      {}
func TestHIncrByFloat(t *testing.T) {}
func TestHKeys(t *testing.T)        {}
func TestHSet(t *testing.T)         {}
func TestHSetNX(t *testing.T)       {}
func TestHStrLen(t *testing.T)      {}
func TestHVals(t *testing.T)        {}
func TestHMGet(t *testing.T)        {}
func TestHMSet(t *testing.T)        {}
func TestHMSlot(t *testing.T)       {}