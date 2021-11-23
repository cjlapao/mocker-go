package help

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlugify(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		{"this is now slugify", "this-is-now-slugify"},
		{"th;s is now slugify", "th-s-is-now-slugify"},
		{"this-is now slugify", "this-is-now-slugify"},
		{"this%^is now slugify", "this-is-now-slugify"},
		{"This%^is now-slugify", "this-is-now-slugify"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v", tt.value, tt.result)
		t.Run(testname, func(t *testing.T) {
			result := Slugify(tt.value)
			assert.Equal(t, tt.result, result)
		})
	}
}
