package memory

import (
	"testing"

	"github.com/tumarsal/demo-go/img-micro-service"
	cachetest "github.com/tumarsal/demo-go/img-micro-service/cache/_test"
	"github.com/tumarsal/demo-go/files"
)

func BenchmarkGet(b *testing.B) {
	for _, tc := range []struct {
		name string
		im   *imageserver.Image
	}{
		{"Small", testdata.Small},
		{"Medium", testdata.Medium},
		{"Large", testdata.Large},
		{"Huge", testdata.Huge},
	} {
		b.Run(tc.name, func(b *testing.B) {
			cch := newTestCache()
			cachetest.BenchmarkGet(b, cch, 1, tc.im) // more parallelism change nothing
		})
	}
}
