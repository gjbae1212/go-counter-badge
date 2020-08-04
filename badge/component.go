package badge

// color is a struct presenting color.
type color string

// bound is a struct presenting bound.
type bound struct {
	dx float64
	dy float64
	x  float64
	y  float64
}

// rect is a struct presenting rect.
type rect struct {
	color color
	bound bound
}

// text is a struct presenting text.
type text struct {
	msg   string
	color color
	bound bound
}

// badge is a struct presenting badge.
type badge struct {
	rect rect
	text text
}

// flatBadge is a struct presenting flat badge.
type flatBadge struct {
	left  badge   // left Badge
	right badge   // right Badge
	rx    string  // horizon radius
	ry    string  // vertical radius
	dx    float64 // Width
	dy    float64 // Height
}
