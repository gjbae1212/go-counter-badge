package badge

import (
	"bytes"
	"fmt"
	"html/template"
)

const (
	defaultBadgeHeight = 20
)

type Badge struct {
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
	fontDrawer
	tmplFlatBadge *template.Template // flat-badge template
}

// RenderFlatBadge renders Flat Badge formatted SVG to byte array.
func (fb *badgeWriter) RenderFlatBadge(b Badge) ([]byte, error) {
	buf := &bytes.Buffer{}

	// TODO:
	badge := &flatBadge{}
	//

	if err := fb.tmplFlatBadge.Execute(buf, badge); err != nil {
		return nil, fmt.Errorf("[err] RenderFlatBadge %w", err)
	}
	return buf.Bytes(), nil
}

// NewWriter returns Badge Writer.
func NewWriter(fontType FontType) (Writer, error) {
	var drawer fontDrawer
	switch fontType {
	case veraSans:
		drawer = veraSansDrawer
	case verdana:
		drawer = verdanaDrawer
	default:
		return nil, fmt.Errorf("[err] NewWriter empty params")
	}

	// make flat-badge template
	tmplFlatBadge, err := template.New("flat-badge").Parse(flatBadgeTemplate)
	if err != nil {
		return nil, fmt.Errorf("[err] NewWriter %w", err)
	}

	writer := &badgeWriter{
		fontDrawer:    drawer,
		tmplFlatBadge: tmplFlatBadge,
	}
	return writer, nil
}
