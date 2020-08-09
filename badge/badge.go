package badge

import (
	"bytes"
	"fmt"
	"html/template"
)

const (
	defaultBadgeHeight = float64(20)
)

type Badge struct {
	FontType FontType

	LeftText            string
	LeftTextColor       string
	LeftBackgroundColor string

	RightText            string
	RightTextColor       string
	RightBackgroundColor string

	XRadius string
	YRadius string
}

// Writer is an interface generating Badge formatted SVG.
type Writer interface {
	RenderFlatBadge(b Badge) ([]byte, error)
}

type badgeWriter struct {
	tmplFlatBadge *template.Template // flat-badge template
}

// RenderFlatBadge renders Flat Badge formatted SVG to byte array.
func (fb *badgeWriter) RenderFlatBadge(b Badge) ([]byte, error) {
	drawer, err := getFontDrawer(b.FontType)
	if err != nil {
		return nil, fmt.Errorf("[err] RenderFlatBadge %w", err)
	}

	// default dy
	dy := defaultBadgeHeight

	// set x,y radius
	flatBadge := &flatBadge{FontFamily: drawer.getFontFamily(), FontSize: drawer.getFontSize()}
	flatBadge.Rx = b.XRadius
	flatBadge.Ry = b.YRadius

	// set left
	leftDx := drawer.measureString(b.LeftText)
	flatBadge.Left = badge{
		Rect: rect{Color: color(b.LeftBackgroundColor), Bound: bound{
			Dx: leftDx,
			Dy: dy,
			X:  0,
			Y:  0,
		}},
		Text: text{Msg: b.LeftText, Color: color(b.LeftTextColor), Bound: bound{
			Dx: 0, // not use
			Dy: 0, // not use
			X:  leftDx/2.0 + 1,
			Y:  15,
		}},
	}

	// set right
	rightDx := drawer.measureString(b.RightText)
	flatBadge.Right = badge{
		Rect: rect{Color: color(b.RightBackgroundColor), Bound: bound{
			Dx: rightDx,
			Dy: dy,
			X:  leftDx,
			Y:  0,
		}},
		Text: text{Msg: b.RightText, Color: color(b.RightTextColor), Bound: bound{
			Dx: 0, // not use
			Dy: 0, // not use
			X:  leftDx + rightDx/2.0 - 1,
			Y:  15,
		}},
	}

	// set dx, dy
	flatBadge.Dy = defaultBadgeHeight
	flatBadge.Dx = leftDx + rightDx

	buf := &bytes.Buffer{}
	if err := fb.tmplFlatBadge.Execute(buf, flatBadge); err != nil {
		return nil, fmt.Errorf("[err] RenderFlatBadge %w", err)
	}
	return buf.Bytes(), nil
}

// NewWriter returns Badge Writer.
func NewWriter() (Writer, error) {
	// make flat-badge template
	tmplFlatBadge, err := template.New("flat-badge").Parse(flatBadgeTemplate)
	if err != nil {
		return nil, fmt.Errorf("[err] NewWriter %w", err)
	}

	writer := &badgeWriter{
		tmplFlatBadge: tmplFlatBadge,
	}
	return writer, nil
}
