package src_test

import (
	"go_sync/src"
	"sync"
	"testing"
	"time"
)

type opType uint8

const (
	get opType = iota
	put
)

const defaultTimeout = time.Second
const shortDely = time.Millisecond * 500
const longDely = time.Millisecond * 1500

type testOp struct {
	op   opType
	k    int
	v    int
	out  bool
	dely time.Duration
}

func TestConcurrent(t *testing.T) {
	cm := src.NewConcurrentMap()
	wg := sync.WaitGroup{}
	testCases := []struct {
		name string
		tops []testOp
	}{
		{
			name: "withoutBlock",
			tops: []testOp{
				{op: put, k: 1, v: 11},
				{op: get, k: 1, v: 11},
				{op: put, k: 2, v: 22},
				{op: get, k: 2, v: 22},
				{op: put, k: 3, v: 33},
				{op: get, k: 3, v: 33},
				{op: put, k: 4, v: 44},
			},
		},
		{
			name: "withBlock",
			tops: []testOp{
				{op: put, k: 1, v: 11, dely: shortDely},
				{op: get, k: 1, v: 11},
				{op: put, k: 2, v: 22, dely: shortDely},
				{op: get, k: 2, v: 22},
				{op: put, k: 3, v: 33},
				{op: get, k: 3, v: 33},
			},
		},
		{
			name: "withTimeout",
			tops: []testOp{
				{op: put, k: 1, v: 11},
				{op: get, k: 1, v: 11},
				{op: put, k: 2, v: 22, dely: longDely},
				{op: get, k: 2, v: 22, out: true},
				{op: put, k: 3, v: 33},
				{op: get, k: 3, v: 33},
				{op: get, k: 4, v: 44, out: true},
			},
		},
	}

	for _, testCase := range testCases {
		wg.Add(len(testCase.tops))
		t.Run(testCase.name, func(t *testing.T) {
			for _, tc := range testCase.tops {
				// put
				if tc.op == put {
					go func(tc testOp) {
						// 模拟延迟
						if tc.dely != 0 {
							time.Sleep(tc.dely)
						}
						cm.Put(tc.k, tc.v)
						wg.Done()
					}(tc)
				}
				// get
				if tc.op == get {
					go func(k, v int, out bool) {
						actual, err := cm.Get(k, defaultTimeout)
						if actual != v {
							t.Errorf("the value of key(%d) is %d, expected: %d", k, actual, v)
						}
						if err != nil && !out {
							t.Errorf("get with unexpected timeout, k: %d", k)
						}
						wg.Done()
					}(tc.k, tc.v, tc.out)
				}
			}
		})
	}
	wg.Wait()
}
