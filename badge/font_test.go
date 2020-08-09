package badge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/image/math/fixed"
)

func TestFontM_FixedPointToPoint(t *testing.T) {
	assert := assert.New(t)

	drawer := &fontM{}

	tests := map[string]struct {
		input  fixed.Int26_6
		output float64
	}{
		"0":        {input: fixed.I(0), output: 0},
		"100":      {input: fixed.I(100), output: 100},
		"-255":     {input: fixed.I(-255), output: -255},
		"20.25":    {input: fixed.I(20), output: 20},
		"-100.875": {input: fixed.I(-100), output: -100},
	}

	for _, t := range tests {
		assert.Equal(t.output, drawer.fixedToPoint(t.input))
		fmt.Printf("%s = %f\n", t.input.String(), t.output)
	}
}

func TestFontM_MeasureString(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		drawer fontDrawer
		input  string
		output float64
	}{
		"vera_sans_1": {drawer: veraSansDrawer, input: "allan hi", output: 53},
		"verdana_1":   {drawer: verdanaDrawer, input: "allan hi", output: 51},
		"vera_sans_2": {drawer: veraSansDrawer, input: "hits", output: 34},
		"verdana_2":   {drawer: verdanaDrawer, input: "hits", output: 30},
	}

	for _, t := range tests {
		assert.Equal(t.output, t.drawer.measureString(t.input))
		fmt.Printf("%s = %f \n", t.input, t.drawer.measureString(t.input))
	}

}

func TestFontM_FontSize(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		drawer fontDrawer
		output int
	}{
		"ok": {drawer: &fontM{fontSize: fontSize}, output: 11},
	}

	for _, t := range tests {
		assert.Equal(t.output, t.drawer.getFontSize())
	}

}

func TestFontM_FontFamily(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		drawer fontDrawer
		output string
	}{
		"ok": {drawer: &fontM{fontFamily: fontFamilyVeraSans}, output: fontFamilyVeraSans},
	}

	for _, t := range tests {
		assert.Equal(t.output, t.drawer.getFontFamily())
	}
}

func TestGetFontDrawer(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		input  FontType
		output fontDrawer
		isErr  bool
	}{
		"fail":      {input: 0, isErr: true},
		"vera sans": {input: VeraSans, isErr: false, output: veraSansDrawer},
		"verdana":   {input: Verdana, isErr: false, output: verdanaDrawer},
	}

	for _, t := range tests {
		drawer, err := getFontDrawer(t.input)
		assert.Equal(t.isErr, err != nil)
		if err == nil {
			assert.Equal(t.output, drawer)
		}

	}
}
