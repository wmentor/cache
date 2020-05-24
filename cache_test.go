package cache

import (
	"strconv"
	"testing"
)

func TestCache(t *testing.T) {

	c := New("")
	if c == nil {
		t.Fatal("New failed")
	}

	c = New("size=100")
	if c == nil {
		t.Fatal("New failed")
	}

	for i := 1; i <= 200; i++ {
		c.Set(strconv.Itoa(i), i)
	}

	for i := 1; i <= 70; i++ {
		if _, has := c.Get(strconv.Itoa(i)); has {
			t.Fatalf("Not expect key: %d", i)
		}
	}

	for i := 71; i <= 200; i++ {
		v, h := c.Get(strconv.Itoa(i))
		if !h {
			t.Fatalf("Not found key: %d", i)
		}

		if v.(int) != i {
			t.Fatalf("Invalid value for: %d", i)
		}
	}

	c.Set("190", 210)
	if v, h := c.Get("190"); !h || v != 210 {
		t.Fatal("rewrite key failed")
	}

	c.Flush()

	for i := 1; i <= 200; i++ {
		if _, has := c.Get(strconv.Itoa(i)); has {
			t.Fatalf("Not expect key: %d", i)
		}
	}
}
