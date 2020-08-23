package badge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWriter(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		isErr bool
	}{
		"ok": {isErr: false},
	}

	for _, t := range tests {
		_, err := NewWriter()
		assert.Equal(t.isErr, err != nil)
	}
}

// svg parser https://www.rapidtables.com/web/tools/svg-viewer-editor.html
func TestBadgeWriter_RenderFlatBadge(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		input  Badge
		outupt []byte
		isErr  bool
	}{
		"verasans-flat": {input: Badge{
			FontType:             VeraSans,
			LeftText:             "verasans-flat",
			LeftTextColor:        "#fff",
			LeftBackgroundColor:  "#555",
			RightText:            "10 / 23234",
			RightTextColor:       "#fff",
			RightBackgroundColor: "#4c1",
			XRadius:              "3",
			YRadius:              "3",
		}},
		"verasans-round": {input: Badge{
			FontType:             VeraSans,
			LeftText:             "verasans-round",
			LeftTextColor:        "#1E9268",
			LeftBackgroundColor:  "#252050",
			RightText:            "1 / 202",
			RightTextColor:       "#fff",
			RightBackgroundColor: "#4c1",
			XRadius:              "auto",
			YRadius:              "auto",
		}},
		"verdana-flat": {input: Badge{
			FontType:             Verdana,
			LeftText:             "verdana-flat",
			LeftTextColor:        "#fff",
			LeftBackgroundColor:  "#555",
			RightText:            "10 / 23234",
			RightTextColor:       "#E25C9F",
			RightBackgroundColor: "#502038",
			XRadius:              "3",
			YRadius:              "3",
		}},
		"verdana-round": {input: Badge{
			FontType:             Verdana,
			LeftText:             "verdand-round",
			LeftTextColor:        "#fff",
			LeftBackgroundColor:  "#555",
			RightText:            "1 / 202",
			RightTextColor:       "#fff",
			RightBackgroundColor: "#4c1",
			XRadius:              "auto",
			YRadius:              "auto",
		}},
	}

	writer, err := NewWriter()
	assert.NoError(err)
	for _, t := range tests {
		result, err := writer.RenderFlatBadge(t.input)
		assert.Equal(t.isErr, err != nil)
		fmt.Println(string(result))
	}
}

// svg parser https://www.rapidtables.com/web/tools/svg-viewer-editor.html
func TestBadgeWriter_RenderIconBadge(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		input     Badge
		iconName  string
		iconColor string
		outupt    []byte
		isErr     bool
	}{
		"verasans-flat": {input: Badge{
			FontType:             VeraSans,
			LeftText:             "verasans-flat",
			LeftTextColor:        "#fff",
			LeftBackgroundColor:  "#555",
			RightText:            "10 / 23234",
			RightTextColor:       "#fff",
			RightBackgroundColor: "#4c1",
			XRadius:              "3",
			YRadius:              "3",
		}, iconName: "appveyor.svg", iconColor: "#00B3E0"},
		"verasans-round": {input: Badge{
			FontType:             VeraSans,
			LeftText:             "verasans-round",
			LeftTextColor:        "#1E9268",
			LeftBackgroundColor:  "#252050",
			RightText:            "1 / 202",
			RightTextColor:       "#fff",
			RightBackgroundColor: "#4c1",
			XRadius:              "auto",
			YRadius:              "auto",
		}, iconName: "appveyor.svg", iconColor: "#3c5688"},
		"verdana-flat": {input: Badge{
			FontType:             Verdana,
			LeftText:             "verdana-flat",
			LeftTextColor:        "#fff",
			LeftBackgroundColor:  "#555",
			RightText:            "10 / 23234",
			RightTextColor:       "#E25C9F",
			RightBackgroundColor: "#502038",
			XRadius:              "3",
			YRadius:              "3",
		}, iconName: "amazon.svg", iconColor: "#0000ff"},
		"verdana-round": {input: Badge{
			FontType:             Verdana,
			LeftText:             "verdand-round",
			LeftTextColor:        "#fff",
			LeftBackgroundColor:  "#555",
			RightText:            "1 / 202",
			RightTextColor:       "#fff",
			RightBackgroundColor: "#4c1",
			XRadius:              "auto",
			YRadius:              "auto",
		}, iconName: "babel.svg", iconColor: "#109556"},
		"verdana-hits": {input: Badge{
			FontType:             Verdana,
			LeftText:             "hits",
			LeftTextColor:        "#fff",
			LeftBackgroundColor:  "#555",
			RightText:            "1 / 202",
			RightTextColor:       "#fff",
			RightBackgroundColor: "#4c1",
			XRadius:              "auto",
			YRadius:              "auto",
		}, iconName: "aircall.svg", iconColor: "#ffffff"},
	}

	writer, err := NewWriter()
	assert.NoError(err)
	for _, t := range tests {
		result, err := writer.RenderIconBadge(t.input, t.iconName, t.iconColor)
		assert.Equal(t.isErr, err != nil)
		fmt.Println(string(result))
	}
}

func TestGetIconsMap(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
	}{
		"success": {},
	}

	for _, _ = range tests{
		icons := GetIconsMap()
		assert.Equal(len(icons), len(iconsMap))
	}

}

func BenchmarkBadgeWriter_RenderFlatBadge(b *testing.B) {
	m, _ := NewWriter()
	for i := 0; i < b.N; i++ {
		m.RenderFlatBadge(Badge{
			FontType:             Verdana,
			LeftText:             "verdand-round",
			LeftTextColor:        "#fff",
			LeftBackgroundColor:  "#555",
			RightText:            "1 / 202",
			RightTextColor:       "#fff",
			RightBackgroundColor: "#4c1",
			XRadius:              "auto",
			YRadius:              "auto",
		})
	}
}

func BenchmarkBadgeWriter_RenderIconBadge(b *testing.B) {
	m, _ := NewWriter()
	for i := 0; i < b.N; i++ {
		m.RenderIconBadge(Badge{
			FontType:             Verdana,
			LeftText:             "verdand-round",
			LeftTextColor:        "#fff",
			LeftBackgroundColor:  "#555",
			RightText:            "1 / 202",
			RightTextColor:       "#fff",
			RightBackgroundColor: "#4c1",
			XRadius:              "auto",
			YRadius:              "auto",
		}, "appveyor.svg", "#00B3E0")
	}
}
