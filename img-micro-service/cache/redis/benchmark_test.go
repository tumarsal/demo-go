package redis

import (
	"strconv"
	"testing"

	"github.com/tumarsal/demo-go/img-micro-service"
	cachetest "github.com/tumarsal/demo-go/img-micro-service/cache/_test"
	"github.com/tumarsal/demo-go/files"
)

func BenchmarkGetSize(b *testing.B) {
	for _, tc := range []struct {
		name string
		im   *imageserver.Image
	}{
		{"Small", testdata.Small},
		{"Medium", testdata.Medium},
		{"Large", testdata.Large},
		{"Huge", testdata.Huge},
	} {
		benchmarkGet(b, tc.name, tc.im, 1)
	}
}

func BenchmarkGetParallelism(b *testing.B) {
	for _, p := range []int{
		1, 2, 4, 8, 16, 32, 64, 128,
	} {
		benchmarkGet(b, strconv.Itoa(p), testdata.Medium, p)
	}
}

func benchmarkGet(b *testing.B, name string, image *imageserver.Image, parallelism int) {
	b.Run(name, func(b *testing.B) {
		cch := newTestCache(b)
		defer func() {
			_ = cch.Client.Close()
		}()
		cachetest.BenchmarkGet(b, cch, parallelism, image)
	})
}
