package skiplist

import (
	"math/rand"
	"time"
)

type SkipList struct {
	cache  []*node
	length int
	level  int
	rand   *rand.Rand
}

type node struct {
	cache []*node
	score int
	value interface{}
	level int
}

func New() *SkipList {
	const defaultLevelMax = 32
	return &SkipList{
		cache: make([]*node, 32),
		level: defaultLevelMax,
		rand:  rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (s *SkipList) randLevel() int {
	const PROPABILITY = 0x3FFF
	l := 1

	for ((s.rand.Int63() >> 32) & 0xFFFF) < PROPABILITY {
		l++
	}

	if l > s.level {
		l = s.level
	}

	return l
}

func (s *SkipList) Add(score int, value interface{}) {
	level := s.randLevel() // [1, defaultLevelMax]

	n := &node{
		score: score,
		value: value,
		cache: make([]*node, level),
		level: level,
	}

	var prev *[]*node = &s.cache
	var nextnode *node
	for i := s.level - 1; i >= 0; i-- {
		nextnode = (*prev)[i]
		for nextnode != nil && nextnode.score <= score {
			prev = &nextnode.cache
			nextnode = nextnode.cache[i]
		}

		if i <= level-1 {
			n.cache[i] = (*prev)[i]
			(*prev)[i] = n
		}
	}

	s.length++
}

func (s *SkipList) Peek() (score int, value interface{}) {
	node := s.cache[0]
	if node == nil {
		return
	}

	return node.score, node.value
}

func (s *SkipList) Pop() (score int, value interface{}) {
	node := s.cache[0]
	if node == nil {
		return
	}

	score = node.score
	value = node.value

	// delete cache
	for i := 0; i < node.level; i++ {
		s.cache[i] = node.cache[i]
	}

	s.length--
	return
}

func (s *SkipList) Length() int {
	return s.length
}

func (s *SkipList) Free() {
	s = &SkipList{}
}
