package cache

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Parallel()

	c := New(100)
	require.NotNil(t, c)

	for i := 1; i <= 200; i++ {
		c.Set(strconv.Itoa(i), i)
	}

	for i := 1; i <= 70; i++ {
		_, has := c.Get(strconv.Itoa(i))
		require.False(t, has)
	}

	for i := 71; i <= 200; i++ {
		v, h := c.Get(strconv.Itoa(i))
		require.True(t, h)
		require.True(t, v.(int) == i)
	}

	c.Set("190", 210)
	v, h := c.Get("190")
	require.True(t, h)
	require.Equal(t, 210, v.(int))

	c.Flush()

	for i := 1; i <= 200; i++ {
		_, has := c.Get(strconv.Itoa(i))
		require.False(t, has)
	}
}
