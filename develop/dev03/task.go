package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortingFlags struct {
	column  int
	num     bool
	reverse bool
	unique  bool
}

func fileRead(buf *bufio.Scanner) []string {
	s := make([]string, 0)

	for buf.Scan() {
		s = append(s, buf.Text())
	}
	return s
}

func standartSort(msg []string) []string {
	sort.SliceStable(msg, func(i, j int) bool { return strings.ToLower(msg[i]) < strings.ToLower(msg[j]) })
	return msg
}

func uniqueSort(msg []string) []string {
	s := make(map[string]bool)
	ret := make([]string, 0)
	for i := 0; i < len(msg); i++ {
		if !s[msg[i]] {
			s[msg[i]] = true
			ret = append(ret, msg[i])
		}
	}
	return ret
}

func reverseSort(msg []string) []string {
	for i, j := 0, len(msg)-1; i < j; i, j = i+1, j-1 {
		msg[i], msg[j] = msg[j], msg[i]
	}
	return msg
}

func columnSort(msg []string, fl *SortingFlags) []string {
	s := make([][]string, 0)

	k := fl.column - 1
	if k < 0 {
		k = 0
	}

	for _, str := range msg {
		s = append(s, strings.Split(str, " "))
	}

	if fl.num {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				x, err := strconv.Atoi(s[i][k])
				if err != nil {
					fmt.Println(err)
					return false
				}
				y, err := strconv.Atoi(s[j][k])
				if err != nil {
					fmt.Println(err)
					return false
				}
				return x < y
			}
			return false
		})
	} else {
		sort.SliceStable(s, func(i, j int) bool {
			if len(s[i]) > k && len(s[j]) > k {
				return strings.ToLower(s[i][k]) < strings.ToLower(s[j][k])
			}
			return false
		})
	}

	ret := make([]string, 0)
	// строка файла которая была разделена пробелом, джониться обратно c пробелом
	for i := 0; i < len(s); i++ {
		str := strings.Join(s[i], " ")
		ret = append(ret, str)
	}

	// возвращаем отсортированный слайс
	return ret
}

func sortFile(msg []string, fl *SortingFlags) []byte {
	msg = standartSort(msg)

	if fl.unique {
		msg = uniqueSort(msg)
	}

	if fl.column > -1 {
		msg = columnSort(msg, fl)
	}

	if fl.reverse {
		msg = reverseSort(msg)
	}

	return []byte(strings.Join(msg, "\n"))
}

func concat(x, y string) string {
	var builder strings.Builder
	builder.Grow(len(x) + len(y)) // Эта строка выделяет память
	builder.WriteString(x)        //Записывает в builder строку.
	builder.WriteString(y)
	return builder.String()
}

var (
	column  int
	num     bool
	reverse bool
	unique  bool
)

func main() {
	flag.IntVar(&column, "k", -1, "колонка для сортировки")
	flag.BoolVar(&num, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "убрать повторяющиеся значения")
	flag.Parse()

	fl := &SortingFlags{unique: unique, column: column, reverse: reverse, num: num}
	// fl := &SortingFlags{unique: true, column: 2, reverse: true, num: true}
	fmt.Println(fl)
	filename := flag.Arg(0)
	// filename := "text.txt"
	fmt.Println(filename)
	if filename == "" {
		fmt.Println("error: enter the file name")
		os.Exit(1)
	}
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}
	buf := bufio.NewScanner(f)
	msg := fileRead(buf)
	err = ioutil.WriteFile(concat("Sorted", f.Name()), sortFile(msg, fl), fs.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}
