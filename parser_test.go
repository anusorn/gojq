package gojq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		src      string
		expected *Query
		err      string
	}{
		{
			src: "",
			err: ":0:0:",
		},
		{
			src: ".",
			expected: &Query{
				Pipe: &Pipe{
					Terms: []*Term{
						&Term{Identity: &Identity{}},
					},
				},
			},
		},
		{
			src: `.foo | . | .["bar"]?`,
			expected: &Query{
				Pipe: &Pipe{
					Terms: []*Term{
						&Term{ObjectIndex: &ObjectIndex{"foo", false}},
						&Term{Identity: &Identity{}},
						&Term{ObjectIndex: &ObjectIndex{"bar", true}},
					},
				},
			},
		},
		{
			src: "abc",
			err: ":1:1:",
		},
	}

	for _, tc := range testCases {
		p := NewParser()
		t.Run(tc.src, func(t *testing.T) {
			q, err := p.Parse(tc.src)
			assert.Equal(t, tc.expected, q)
			if err == nil {
				assert.Equal(t, "", tc.err)
			} else {
				assert.NotEqual(t, "", tc.err, err.Error())
				assert.Contains(t, err.Error(), tc.err)
			}
		})
	}
}