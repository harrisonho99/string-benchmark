package concat_test

import (
	"concat_string/concat"
	"testing"
)

const (
	num  = 100
	show = false
)

func BenchmarkConcatString(b *testing.B) {
	testing.Init()

	b.StartTimer()
	concat.UseConcat(num, show)
	b.StopTimer()

	b.StartTimer()
	concat.UseStringBuilder(num, show)
	b.StopTimer()

	b.StartTimer()
	concat.UseBytesBuffer(num, show)
	b.StopTimer()
}
