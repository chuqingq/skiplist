package skiplist

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

//------------------------------------------------------------------------------

func Test_test1(t *testing.T) {
	h := New()

	var a int = 10
	h.Add(a, &a)
	t.Logf("cache: %v\n", h.cache)
	t.Logf("h: %s\n", h)

	var b int = 100
	h.Add(b, &b)

	var c int = 1000
	h.Add(c, &c)

	var d int = 100
	h.Add(d, &d)

	if h.Length() != 4 {
		t.Fatalf("h.Length() unexpected: %d\n", h.Length())
	}

	vs, vv := h.Pop()
	if vs != a || vv != &a {
		t.Fatalf("vs unexpected a: %d\n", vs)
	}

	// 需要保证顺序，b在d前
	vs, vv = h.Peek()
	if vs != b || vv != &b {
		t.Fatalf("vs unexpected b1: %d\n", vs)
	}

	if h.Length() != 3 {
		t.Fatalf("h.Length() unexpected: %d\n", h.Length())
	}

	vs, vv = h.Pop()
	if vs != b || vv != &b {
		t.Fatalf("vs unexpected b: %d\n", vs)
	}

	if h.Length() != 2 {
		t.Fatalf("h.Length() unexpected: %d\n", h.Length())
	}

	vs, vv = h.Pop()
	if vs != d || vv != &d {
		t.Fatalf("vs unexpected d: %d\n", vs)
	}

	vs, vv = h.Pop()
	if vs != c || vv != &c {
		t.Fatalf("vs unexpected c: %d\n", vs)
	}
}

//------------------------------------------------------------------------------

func Benchmark_bench1(b *testing.B) {
	b.StopTimer()

	count := 100000

	// generate test data
	data := make([]int, count)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < count; i++ {
		data[i] = r.Intn(count)
	}

	b.StartTimer()
	h := New()
	for _, i := range data {
		h.Add(i, &i)
	}

	b.StopTimer()
	sort.Ints(data)

	for index, d := range data {
		s, _ := h.Pop()
		if d != s {
			b.Fatalf("benchmark unexpected: %d, %d, %d\n", index, d, s)
		}
	}
}
