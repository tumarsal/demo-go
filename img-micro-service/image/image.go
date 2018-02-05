// Package image provides a bridge to the Go "image" package.
package image

import "github.com/tumarsal/demo-go/img-micro-service"

// Changer returns true if the Image could change for the given Params.
type Changer interface {
	Change(imageserver.Params) bool
}
