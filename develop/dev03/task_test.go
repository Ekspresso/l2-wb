package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type tableTest struct {
	name    string
	column  int
	reverse bool
	unique  bool
	num     bool
	file    string
	exp     []string
}

var tableTests = []tableTest{
	{
		name:    "test1",
		column:  2,
		reverse: true,
		unique:  true,
		num:     true,
		file:    "text.txt",
		exp: []string{
			"drwxrwxr-x 9 konstantin konstantin 4096 ноя 28 11:55 l1",
			"drwxrwxr-x 6 konstantin konstantin 4096 дек 23 05:21 l2",
			"drwxrwxr-x 4 konstantin konstantin 4096 ноя 18 07:54 l0",
			"drwxrwxr-x 2 konstantin konstantin 4096 ноя 17 02:45 .vscode",
			"-rw-rw-r-- 1 konstantin konstantin 1347 ноя 18 09:50 psql_commands_l0",
		},
	},
	{
		name:    "test2",
		column:  9,
		reverse: true,
		unique:  true,
		num:     false,
		file:    "text.txt",
		exp: []string{
			"-rw-rw-r-- 1 konstantin konstantin 1347 ноя 18 09:50 psql_commands_l0",
			"drwxrwxr-x 6 konstantin konstantin 4096 дек 23 05:21 l2",
			"drwxrwxr-x 9 konstantin konstantin 4096 ноя 28 11:55 l1",
			"drwxrwxr-x 4 konstantin konstantin 4096 ноя 18 07:54 l0",
			"drwxrwxr-x 2 konstantin konstantin 4096 ноя 17 02:45 .vscode",
		},
	},
}

func TestSortFile(t *testing.T) {
	for _, test := range tableTests {
		fl := &SortingFlags{
			column:  test.column,
			reverse: test.reverse,
			unique:  test.unique,
			num:     test.num,
		}
		f, err := os.Open(test.file)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		buf := bufio.NewScanner(f)
		s := make([]string, 0)
		for buf.Scan() {
			s = append(s, buf.Text())
		}
		t.Run(test.name, func(t *testing.T) {
			res := strings.Split(string(sortFile(s, fl)), "\n")
			require.Equal(t, test.exp, res)
		})
	}
}
