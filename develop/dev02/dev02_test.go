package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type expected struct {
	longStr string
	err     error
}

type upsTest struct {
	name, shortStr string
	expected       expected
}

var upsTests = []upsTest{
	{
		name:     "test1",
		shortStr: "a4bc2d5e",
		expected: expected{
			longStr: "aaaabccddddde",
			err:     nil,
		},
	},
	{
		name:     "test2",
		shortStr: "abcd",
		expected: expected{
			longStr: "abcd",
			err:     nil,
		},
	},
	{
		name:     "test3",
		shortStr: "45",
		expected: expected{
			longStr: "",
			err:     ErrInvalid,
		},
	},
	{
		name:     "test4",
		shortStr: "",
		expected: expected{
			longStr: "",
			err:     nil,
		},
	},
	{
		name:     "test5",
		shortStr: "😂⌘👍4",
		expected: expected{
			longStr: "😂⌘👍👍👍👍",
			err:     nil,
		},
	},
	{
		name:     "test6",
		shortStr: `qwe\4\5`,
		expected: expected{
			longStr: "qwe45",
			err:     nil,
		},
	},
	{
		name:     "test7",
		shortStr: `qwe\45`,
		expected: expected{
			longStr: "qwe44444",
			err:     nil,
		},
	},
	{
		name:     "test8",
		shortStr: `qwe\\5`,
		expected: expected{
			longStr: `qwe\\\\\`,
			err:     nil,
		},
	},
	{
		name:     "test9",
		shortStr: `qwe\\5\`,
		expected: expected{
			longStr: "",
			err:     ErrInvalid,
		},
	},
	{
		name:     "test10",
		shortStr: `\`,
		expected: expected{
			longStr: "",
			err:     ErrInvalid,
		},
	},
	{
		name:     "test11",
		shortStr: `\\`,
		expected: expected{
			longStr: `\`,
			err:     nil,
		},
	},
	{
		name:     "test12",
		shortStr: `\\\`,
		expected: expected{
			longStr: "",
			err:     ErrInvalid,
		},
	},
	{
		name:     "test13",
		shortStr: `\😂⌘\👍4`,
		expected: expected{
			longStr: "😂⌘👍👍👍👍",
			err:     nil,
		},
	},
	{
		name:     "test14",
		shortStr: `0ghjc`,
		expected: expected{
			longStr: "",
			err:     ErrInvalid,
		},
	},
	{
		name:     "test15",
		shortStr: `\\\\`,
		expected: expected{
			longStr: `\\`,
			err:     nil,
		},
	},
	{
		name:     "test16",
		shortStr: `\\\\\`,
		expected: expected{
			longStr: "",
			err:     ErrInvalid,
		},
	},
	{
		name:     "test17",
		shortStr: `hjfh\`,
		expected: expected{
			longStr: "",
			err:     ErrInvalid,
		},
	},
}

func TestUnpackStr(t *testing.T) {
	for _, test := range upsTests {
		t.Run(test.name, func(t *testing.T) {
			res, err := UnpackStr(test.shortStr)
			require.Equal(t, test.expected.longStr, res)
			require.Equal(t, test.expected.err, err)
		})
	}
}
