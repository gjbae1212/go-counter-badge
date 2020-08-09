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
