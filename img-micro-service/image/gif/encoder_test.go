package gif

import (
	"testing"

	"github.com/tumarsal/demo-go/img-micro-service"
	imageserver_image "github.com/tumarsal/demo-go/img-micro-service/image"
	imageserver_image_test "github.com/tumarsal/demo-go/img-micro-service/image/_test"
)

var _ imageserver_image.Encoder = &Encoder{}

func TestEncoder(t *testing.T) {
	imageserver_image_test.TestEncoder(t, &Encoder{}, "gif")
}

func TestEncoderChange(t *testing.T) {
	c := (&Encoder{}).Change(imageserver.Params{})
	if c {
		t.Fatal("not false")
	}
}
