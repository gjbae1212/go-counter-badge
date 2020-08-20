package badge

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"strings"

	"github.com/gjbae1212/go-counter-badge/internal/assets"
)

const (
	defaultBadgeHeight = float64(20)
	defaultIconWidth   = float64(14)
	defaultIconHeight  = float64(14)
	defaultIconX       = float64(3)
	defaultIconY       = float64(3)
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
	RenderIconBadge(b Badge, iconName, iconColor string) ([]byte, error)
}

type badgeWriter struct {
	tmplIconBadge *template.Template // icon-badge template
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

// RenderIconBadge renders Icon Badge formatted SVG to byte array.
func (fb *badgeWriter) RenderIconBadge(b Badge, iconName, iconColor string) ([]byte, error) {
	if iconName == "" {
		return nil, fmt.Errorf("[err] RenderIconBadge empty params")
	}
	iconBytes, err := assets.Asset("icons/" + iconName + ".svg")
	if err != nil {
		return nil, fmt.Errorf("[err] RenderIconBadge not found icons")
	}
	drawer, err := getFontDrawer(b.FontType)
	if err != nil {
		return nil, fmt.Errorf("[err] RenderFlatBadge %w", err)
	}

	// fill color
	iconsvg := string(iconBytes)
	iconsvg = strings.Replace(iconsvg, "<svg", fmt.Sprintf("<svg fill=\"%s\" ", iconColor), 1)

	// default dy
	dy := defaultBadgeHeight

	// set x,y radius
	iconBadge := &iconBadge{FontFamily: drawer.getFontFamily(), FontSize: drawer.getFontSize()}
	iconBadge.Rx = b.XRadius
	iconBadge.Ry = b.YRadius

	// set icon
	iconDx := defaultIconWidth + 2*defaultIconX
	iconBadge.Icon.Bound.X = defaultIconX
	iconBadge.Icon.Bound.Y = defaultIconY
	iconBadge.Icon.Bound.Dx = iconDx
	iconBadge.Icon.Bound.Dy = defaultIconHeight
	iconBadge.Icon.Base64 = base64.StdEncoding.EncodeToString([]byte(iconsvg))

	// set left
	leftDx := drawer.measureString(b.LeftText)
	iconBadge.Left = badge{
		Rect: rect{Color: color(b.LeftBackgroundColor), Bound: bound{
			Dx: leftDx + iconDx,
			Dy: dy,
			X:  0, // not use
			Y:  0, // not use
		}},
		Text: text{Msg: b.LeftText, Color: color(b.LeftTextColor), Bound: bound{
			Dx: 0, // not use
			Dy: 0, // not use
			X:  leftDx/2.0 - 1 + iconDx,
			Y:  15,
		}},
	}

	// set right
	rightDx := drawer.measureString(b.RightText)
	iconBadge.Right = badge{
		Rect: rect{Color: color(b.RightBackgroundColor), Bound: bound{
			Dx: rightDx,
			Dy: dy,
			X:  leftDx + iconDx,
			Y:  0,
		}},
		Text: text{Msg: b.RightText, Color: color(b.RightTextColor), Bound: bound{
			Dx: 0, // not use
			Dy: 0, // not use
			X:  iconDx + leftDx + rightDx/2.0,
			Y:  15,
		}},
	}

	// set dx, dy
	iconBadge.Dy = defaultBadgeHeight
	iconBadge.Dx = leftDx + rightDx + iconDx

	buf := &bytes.Buffer{}
	if err := fb.tmplIconBadge.Execute(buf, iconBadge); err != nil {
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

	tmplIconBadge, err := template.New("icon-badge").Parse(iconBadgeTemplate)
	if err != nil {
		return nil, fmt.Errorf("[err] NewWriter %w", err)
	}

	writer := &badgeWriter{
		tmplFlatBadge: tmplFlatBadge,
		tmplIconBadge: tmplIconBadge,
	}
	return writer, nil
}
