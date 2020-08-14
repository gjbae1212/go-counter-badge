package badge

import "strings"

// flat Badge template
var flatBadgeTemplate = strings.TrimSpace(`
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.Dx}}" height="{{.Dy}}">
 <linearGradient id="smooth" x2="0" y2="100%">
   <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
   <stop offset="1" stop-opacity=".1"/>
 </linearGradient>

 <mask id="round">
   <rect width="{{.Dx}}" height="{{.Dy}}" rx="{{.Rx}}" ry="{{.Ry}}" fill="#fff"/>
 </mask>

 <g mask="url(#round)">
   <rect width="{{.Left.Rect.Bound.Dx}}" height="{{.Left.Rect.Bound.Dy}}" fill="{{.Left.Rect.Color}}"/>
   <rect x="{{.Left.Rect.Bound.Dx}}" width="{{.Right.Rect.Bound.Dx}}" height="{{.Right.Rect.Bound.Dy}}" fill="{{.Right.Rect.Color}}"/>
   <rect width="{{.Dx}}" height="{{.Dy}}" fill="url(#smooth)"/>
 </g>

 <g fill="#fff" text-anchor="middle" font-family="{{.FontFamily}}" font-size="{{.FontSize}}"> 
   <text x="{{.Left.Text.Bound.X}}" y="{{.Left.Text.Bound.Y}}" fill="#010101" fill-opacity=".3">{{.Left.Text.Msg | html}}</text>
   <text x="{{.Left.Text.Bound.X}}" y="{{.Left.Text.Bound.AddY -1}}" fill="{{.Left.Text.Color}}">{{.Left.Text.Msg | html}}</text>
   <text x="{{.Right.Text.Bound.X}}" y="{{.Right.Text.Bound.Y}}" fill="#010101" fill-opacity=".3">{{.Right.Text.Msg | html}}</text>
   <text x="{{.Right.Text.Bound.X}}" y="{{.Right.Text.Bound.AddY -1}}" fill="{{.Right.Text.Color}}">{{.Right.Text.Msg | html}}</text>
 </g>
</svg>
`)
