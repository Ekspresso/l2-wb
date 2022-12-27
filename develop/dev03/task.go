package main

// Отсортировать строки в файле по аналогии с консольной утилитой sort
// (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

// Реализовать поддержку утилитой следующих ключей:

// -k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
// -n — сортировать по числовому значению
// -r — сортировать в обратном порядке
// -u — не выводить повторяющиеся строки

// Дополнительно

// Реализовать поддержку утилитой следующих ключей:

// -M — сортировать по названию месяца
// -b — игнорировать хвостовые пробелы
// -c — проверять отсортированы ли данные
// -h — сортировать по числовому значению с учетом суффиксов

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

// Структура, хранящая в себе флаги.
type SortingFlags struct {
	column  int
	num     bool
	reverse bool
	unique  bool
}

// fileRead - функция построкового чтения из файла.
func fileRead(buf *bufio.Scanner) []string {
	s := make([]string, 0)

	for buf.Scan() {
		s = append(s, buf.Text())
	}
	return s
}

// standartSort - функция для начальной сортировки строк.
func standartSort(msg []string) []string {
	sort.SliceStable(msg, func(i, j int) bool { return strings.ToLower(msg[i]) < strings.ToLower(msg[j]) })
	return msg
}

// uniqueSort - функция удаляет из слайса строк дубликаты.
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

// reverseSort - функция переворачивает полученный слайс строк.
func reverseSort(msg []string) []string {
	for i, j := 0, len(msg)-1; i < j; i, j = i+1, j-1 {
		msg[i], msg[j] = msg[j], msg[i]
	}
	return msg
}

// columnSort - функция сортировки по заданной колонке. Сортирует как по числам, так и по символам.
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

// sortFile - основная функция сортировки. Последовательно вызывает необходимые функции в зависимости от флагов.
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

// concat - функция конкатенации 2 строк
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
	filename := flag.Arg(0)
	// filename := "text.txt"
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
