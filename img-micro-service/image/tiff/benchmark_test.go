package tiff

import (
	"testing"

	"github.com/tumarsal/demo-go/img-micro-service"
	imageserver_image_test "github.com/tumarsal/demo-go/img-micro-service/image/_test"
	_ "github.com/tumarsal/demo-go/img-micro-service/image/jpeg"
	"github.com/tumarsal/demo-go/files"
)

func Benchmark(b *testing.B) {
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
			imageserver_image_test.BenchmarkEncoder(b, &Encoder{}, tc.im, imageserver.Params{})
		})
	}
}
