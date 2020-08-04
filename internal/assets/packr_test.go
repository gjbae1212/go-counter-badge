package assets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAsset(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		input string
		exist bool
		isErr bool
	}{
		"empty": {
			input: "empty",
			exist: false,
			isErr: true,
		},
		"vera_sans": {
			input: "vera_sans/Vera.ttf",
			exist: true,
			isErr: false,
		},
	}

	for _, t := range tests {
		bys, err := GetAsset(t.input)
		assert.Equal(t.isErr, err != nil)
		assert.Equal(t.exist, len(bys) != 0)
	}
}

func TestGetVeraSansFont(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		isErr bool
	}{
		"ok": {
			isErr: false,
		},
	}

	for _, t := range tests {
		bys, err := GetVeraSansFont()
		assert.Equal(t.isErr, err != nil)
		if err != nil {
			assert.NotEmpty(bys)
		}
	}
}

func TestGetVerdanaFont(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		isErr bool
	}{
		"ok": {
			isErr: false,
		},
	}

	for _, t := range tests {
		bys, err := GetVerdanaFont()
		assert.Equal(t.isErr, err != nil)
		if err != nil {
			assert.NotEmpty(bys)
		}
	}
}
