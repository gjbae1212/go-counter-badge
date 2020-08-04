package badge

import "strings"

// flat Badge template
var flatBadgeTemplate = strings.TrimSpace(`
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{.dx}}" height="{{.dy}}">
 <linearGradient id="smooth" x2="0" y2="100%">
   <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
   <stop offset="1" stop-opacity=".1"/>
 </linearGradient>

 <mask id="round">
   <rect width="{{.dx}}" height="{{.dy}}" rx="{{.rx}}" ry="{{.ry}}" fill="#fff"/>
 </mask>

 <g mask="url(#round)">
   <rect width="{{.left.rect.bound.dx}}" height="{{.left.rect.bound.dy}}" fill="{{.left.rect.color}}"/>
   <rect x="{{.left.rect.bound.dx}}" width="{{.right.rect.bound.dx}}" height="{{.right.rect.bound.dy}}" fill="{{.right.rect.color}}"/>
   <rect width="{{.dx}}" height="{{.dy}}" fill="url(#smooth)"/>
 </g>

 <g fill="#fff" text-anchor="middle" font-family="{{.fontFamily}}" font-size="{{.fontSize}}"> 
   <text x="{{.left.text.bound.x}}" y="{{.left.text.bound.y}}" fill="#010101" fill-opacity=".3">{{.left.text.msg | html}}</text>
   <text x="{{.left.text.bound.x}}" y="{{.left.text.bound.y - 1}}" fill="{{.left.text.color}}">{{.left.text.msg | html}}</text>
   <text x="{{.right.text.bound.x}}" y="{{.right.text.bound.y}}" fill="#010101" fill-opacity=".3">{{.right.text.msg | html}}</text>
   <text x="{{.right.text.bound.x}}" y="{{.right.text.bound.y -1}}" fill="{{.right.text.color}}">{{.right.text.msg | html}}</text>
 </g>
</svg>
`)
