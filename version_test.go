package main_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import "github.com/activatedio/go-release"

func TestParseVersion(t *testing.T) {

	type s struct {
		input  string
		assert func(got *main.Version, err error)
	}

	cases := map[string]s{
		"missing v": {
			input: "1.0.0",
			assert: func(got *main.Version, err error) {

				assert.EqualError(t, err, "invalid version string")
				assert.Nil(t, got)
			},
		},
		"not ending in number": {
			input: "v1.0.0.a",
			assert: func(got *main.Version, err error) {

				assert.EqualError(t, err, "invalid version string")
				assert.Nil(t, got)
			},
		},
		"prerelease": {
			input: "v0.0.1",
			assert: func(got *main.Version, err error) {

				assert.Nil(t, err)
				assert.Equal(t, &main.Version{
					Version:            "v0.0.1",
					Prefix:             "v0.0",
					IncrementingSuffix: 1,
				}, got)
			},
		},
		"with beta": {
			input: "v1.0.0-beta.2",
			assert: func(got *main.Version, err error) {

				assert.Nil(t, err)
				assert.Equal(t, &main.Version{
					Version:            "v1.0.0-beta.2",
					Prefix:             "v1.0.0-beta",
					IncrementingSuffix: 2,
				}, got)
			},
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {

			v.assert(main.ParseVersion(v.input))
		})
	}
}

func TestVersion_Increment(t *testing.T) {

	v, err := main.ParseVersion("v1.0.1")

	assert.Nil(t, err)

	assert.Equal(t, "v1.0.1", v.Version)
	assert.Equal(t, "v1.0", v.Prefix)
	assert.Equal(t, 1, v.IncrementingSuffix)

	v2 := v.Increment()

	assert.Equal(t, "v1.0.1", v.Version)
	assert.Equal(t, "v1.0", v.Prefix)
	assert.Equal(t, 1, v.IncrementingSuffix)

	assert.Equal(t, "v1.0.2", v2.Version)
	assert.Equal(t, "v1.0", v2.Prefix)
	assert.Equal(t, 2, v2.IncrementingSuffix)
}
