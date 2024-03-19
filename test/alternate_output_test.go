package src_test

import (
	"go_sync/src"
	"testing"
)

func TestAlternateOutput(t *testing.T) {
	str1 := src.AlternateOutput1(5, 3)
	t.Logf("双chan有缓冲: %s\n", str1)

	str2 := src.AlternateOutput2(5, 3)
	t.Logf("双chan无缓冲: %s\n", str2)

	str3 := src.AlternateOutput3(5, 3)
	t.Logf("单chan无缓冲: %s\n", str3)
}
