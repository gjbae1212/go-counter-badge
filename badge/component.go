package badge

// color is a struct presenting color.
type color string

// bound is a struct presenting bound.
type bound struct {
	Dx float64
	Dy float64
	X  float64
	Y  float64
}

func (b bound) AddX(i float64) float64 {
	return b.X + i
}

func (b bound) AddY(i float64) float64 {
	return b.Y + i
}

// rect is a struct presenting rect.
type rect struct {
	Color color
	Bound bound
}

// text is a struct presenting text.
type text struct {
	Msg   string
	Color color
	Bound bound
}

// badge is a struct presenting badge.
type badge struct {
	Rect rect
	Text text
}

// flatBadge is a struct presenting flat badge.
type flatBadge struct {
	FontFamily string // font family
	FontSize   int
	Left       badge   // left Badge
	Right      badge   // right Badge
	Rx         string  // horizon radius
	Ry         string  // vertical radius
	Dx         float64 // Width
	Dy         float64 // Height
}
