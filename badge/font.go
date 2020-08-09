package badge

import (
	"fmt"
	"sync"

	"github.com/gjbae1212/go-counter-badge/internal/assets"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type FontType int

const (
	VeraSans FontType = 1 + iota
	Verdana
)

const (
	fontDPI            = 72 // web standard
	fontSize           = 11 // font size
	extraVeraSansDx    = 13 // vera sans dx (text width margin)
	extraVerdanaDx     = 10 // verdana dx (text width margin)
	fontFamilyVeraSans = "DejaVu Sans,Verdana,Geneva,sans-serif"
	fontFamilyVerdana  = "Verdana,DejaVu Sans,Geneva,sans-serif"
)

// fontDrawer gets text size that is applied specified font.
type fontDrawer interface {
	measureString(string) float64
	getFontSize() int
	getFontFamily() string
}

type fontM struct {
	sync.Mutex
	fontSize   int
	extraDx    int
	fontFamily string
	drawer     *font.Drawer
}

// MeasureString returns width for text.
func (fd *fontM) measureString(s string) float64 {
	fd.Lock()
	p := fd.drawer.MeasureString(s)
	fd.Unlock()

	// must be more than 0.
	size := fd.fixedToPoint(p)
	if size <= 0 {
		return 0
	}

	// add extra margin.
	return size + float64(fd.extraDx)
}

// FixedToPoint converts fixed point to floored floating point.
func (fd *fontM) fixedToPoint(p fixed.Int26_6) float64 {
	// convert fixed point to floating point.
	// discard fractions.

	var result float64

	// 26 bit integer(with 1 sign)
	if p < 0 {
		reverse := -p
		result += float64(reverse>>6) * -1
	} else {
		result += float64(p >> 6)
	}
	return result
}

// getFontSize returns font size.
func (fd *fontM) getFontSize() int {
	return fd.fontSize
}

// getFontFamily returns font family.
func (fd *fontM) getFontFamily() string {
	return fd.fontFamily
}

// vera sans font drawer
var veraSansDrawer fontDrawer

// verdana font drawer
var verdanaDrawer fontDrawer

func init() {
	// initialize vera sans font drawer.
	veraTTF, err := assets.GetVeraSansFont()
	if err != nil {
		panic(err)
	}

	veraFont, err := truetype.Parse(veraTTF)
	if err != nil {
		panic(err)
	}

	veraSansDrawer = &fontM{
		fontSize:   fontSize,
		extraDx:    extraVeraSansDx,
		fontFamily: fontFamilyVeraSans,
		drawer: &font.Drawer{
			Face: truetype.NewFace(veraFont, &truetype.Options{
				Size:    fontDPI,
				DPI:     fontSize,
				Hinting: font.HintingFull,
			}),
		},
	}

	// initialize verdana font drawer.
	verdanaTTF, err := assets.GetVerdanaFont()
	if err != nil {
		panic(err)
	}

	verdanaFont, err := truetype.Parse(verdanaTTF)
	if err != nil {
		panic(err)
	}

	verdanaDrawer = &fontM{
		fontSize:   fontSize,
		extraDx:    extraVerdanaDx,
		fontFamily: fontFamilyVeraSans,
		drawer: &font.Drawer{
			Face: truetype.NewFace(verdanaFont, &truetype.Options{
				Size:    fontDPI,
				DPI:     fontSize,
				Hinting: font.HintingFull,
			}),
		},
	}
}

func getFontDrawer(fontType FontType) (fontDrawer, error) {
	switch fontType {
	case VeraSans:
		return veraSansDrawer, nil
	case Verdana:
		return verdanaDrawer, nil
	default:
		return nil, fmt.Errorf("[err] NewWriter empty params")
	}
}
