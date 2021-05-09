package split_test

import (
	"testing"

	"github.com/business_group/test_project/split"
)

func BenchmarkSplit(b *testing.B) {
    for i := 0; i < b.N; i++ {
        split.Split("a,b,c,d,e,f", ",")
    }
}