package thumbnails

import (
	"bytes"
	"image"
	"strings"
	"time"

	"github.com/nfnt/resize"
	"github.com/owncloud/ocis-thumbnails/pkg/thumbnails/cache"
)

// ThumbnailContext bundles information needed to generate a thumbnail for afile
type ThumbnailContext struct {
	Width     int
	Height    int
	ImagePath string
	Encoder   Encoder
}

// Manager is responsible for generating thumbnails
type Manager interface {
	// Get will return a thumbnail for a file
	Get(ThumbnailContext, image.Image) ([]byte, error)
	GetCached(ThumbnailContext) []byte
}

// SimpleManager is a simple implementation of Manager
type SimpleManager struct {
	Cache cache.Cache
}

// Get implements the Get Method of Manager
func (s SimpleManager) Get(ctx ThumbnailContext, img image.Image) ([]byte, error) {
	thumbnail := s.generate(ctx, img)

	key := buildCacheKey(ctx)
	s.Cache.Set(key, thumbnail)

	buf := new(bytes.Buffer)
	err := ctx.Encoder.Encode(buf, thumbnail)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GetCached tries to get the cached thumbnail and return it.
// If there is no cached thumbnail it will return nil
func (s SimpleManager) GetCached(ctx ThumbnailContext) []byte {
	key := buildCacheKey(ctx)
	cached := s.Cache.Get(key)
	if cached == nil {
		return nil
	}
	buf := new(bytes.Buffer)
	ctx.Encoder.Encode(buf, cached)
	return buf.Bytes()
}

func (s SimpleManager) generate(ctx ThumbnailContext, img image.Image) image.Image {
	// TODO: remove, just for demo purposes
	time.Sleep(time.Second * 2)

	thumbnail := resize.Thumbnail(uint(ctx.Width), uint(ctx.Height), img, resize.Lanczos2)
	return thumbnail
}

func buildCacheKey(ctx ThumbnailContext) string {
	parts := []string{
		ctx.ImagePath,
		string(ctx.Width) + "x" + string(ctx.Height),
		strings.Join(ctx.Encoder.Types(), ","),
	}
	return strings.Join(parts, "+")
}
