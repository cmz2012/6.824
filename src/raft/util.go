package raft

import (
	"log"
	"sort"
	"sync/atomic"
)

// Debugging
const kDebug = 0

var Debug int32 = kDebug

func DPrintf(format string, a ...interface{}) (n int, err error) {
	if atomic.LoadInt32(&Debug) > 0 {
		log.Printf(format, a...)
	}
	return
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Middle(a []int) (int, bool) {
	if len(a) == 0 {
		return -1, false
	}
	mid := len(a) / 2
	b := make([]int, len(a))
	copy(b, a)
	sort.Ints(b)
	return b[mid], true
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

const MinIndex = 1
