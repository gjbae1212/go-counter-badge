package assets

import (
	"github.com/gobuffalo/packr/v2"
)

var (
	box *packr.Box
)

func init() {
	box = packr.New("assets", "./")
}

// GetAsset returns asset in assets directory.
func GetAsset(path string) ([]byte, error) {
	return box.Find(path)
}

// GetVeraSansFont returns vera sans font.
func GetVeraSansFont() ([]byte, error) {
	return box.Find("vera_sans/Vera.ttf")
}
