package gif

import (
	"testing"

	"github.com/tumarsal/demo-go/img-micro-service"
	imageserver_image_test "github.com/tumarsal/demo-go/img-micro-service/image/_test"
	_ "github.com/tumarsal/demo-go/img-micro-service/image/jpeg"
	"github.com/tumarsal/demo-go/files"
)

func BenchmarkEncoder(b *testing.B) {
	enc := &Encoder{}
	params := imageserver.Params{}
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
			imageserver_image_test.BenchmarkEncoder(b, enc, tc.im, params)
		})
	}
}
