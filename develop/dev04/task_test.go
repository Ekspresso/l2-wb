package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type angTest struct {
	name     string
	arr      []string
	expected map[string][]string
}

func TestSearchAnagrams(t *testing.T) {
	m := make(map[string][]string)
	m["привет"] = []string{"ветпри", "привет"}
	m[""] = []string{""}
	var angTests = []angTest{
		{
			name:     "test1",
			arr:      []string{"привет", "ветпри", ""},
			expected: m,
		},
	}
	for _, test := range angTests {
		t.Run(test.name, func(t *testing.T) {
			res := searchAnagrams(test.arr)
			require.Equal(t, test.expected, res)
		})
	}
}
