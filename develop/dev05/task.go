package main

// Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

// Реализовать поддержку утилитой следующих ключей:
// -A - "after" печатать +N строк после совпадения
// -B - "before" печатать +N строк до совпадения
// -C - "context" (A+B) печатать ±N строк вокруг совпадения
// -c - "count" (количество строк)
// -i - "ignore-case" (игнорировать регистр)
// -v - "invert" (вместо совпадения, исключать)
// -F - "fixed", точное совпадение со строкой, не паттерн
// -n - "line num", напечатать номер строки

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type GrepFlags struct {
	after     int
	before    int
	context   int
	count     bool
	ignRegist bool
	invert    bool
	fixed     bool
	lineNum   bool
	msg       []string
	sub       string
	out       io.Writer
}

// func (g *GrepFlags) addNumLine() {
// 	for i := 0; i < len(g.sub); i++ {
// 		num := concat(strconv.Itoa(i+1), ": ")
// 		g.msg[i] = concat(num, g.msg[i])
// 	}
// }

func (g *GrepFlags) countLines(reg *regexp.Regexp) {
	c := 0

	for _, val := range g.msg {
		if reg.MatchString(val) {
			c++
		}
	}
	if g.invert {
		fmt.Println(len(g.msg) - c)
	} else {
		fmt.Println(c)
	}
}

func (g *GrepFlags) printResult(indMap map[int]bool) {
	if g.invert {
		for i := 0; i < len(g.msg); i++ {
			if !indMap[i] && g.lineNum {
				fmt.Printf("%d: %s\n", i+1, g.msg[i])
			} else if !indMap[i] {
				fmt.Println(g.msg[i])
			}
		}
	} else {
		for i := 0; i < len(g.msg); i++ {
			if indMap[i] && g.lineNum {
				fmt.Printf("%d: %s\n", i+1, g.msg[i])
			} else if indMap[i] {
				fmt.Println(g.msg[i])
			}
		}
	}
}

func (g *GrepFlags) Grep() {
	var reg *regexp.Regexp
	var err error
	var sub string

	if g.fixed {
		sub = concat(`^`, g.sub)
		sub = concat(sub, `$`)
	} else {
		sub = g.sub
	}
	if g.ignRegist {
		reg, err = regexp.Compile(concat("(?i)", sub))
	} else {
		reg, err = regexp.Compile(sub)
	}
	if err != nil {
		log.Println(err.Error())
		return
	}

	if g.count {
		g.countLines(reg)
		return
	} else {
		g.after = findMax(g.after, g.context)
		g.before = findMax(g.before, g.context)
	}
	indMap := make(map[int]bool)
	for ind, val := range g.msg {
		if reg.MatchString(val) {
			indMap[ind] = true
			for l, count := ind-1, g.before; l >= 0 && count > 0; l, count = l-1, count-1 {
				indMap[l] = true
			}
			for r, count := ind+1, g.after; r < len(g.msg) && count > 0; r, count = r+1, count-1 {
				indMap[r] = true
			}
		}
	}
	g.printResult(indMap)
}

func fileRead(buf *bufio.Scanner) []string {
	s := make([]string, 0)

	for buf.Scan() {
		s = append(s, buf.Text())
	}
	return s
}

func findMax(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
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
	afterFl     int
	beforeFl    int
	contextFl   int
	countFl     bool
	ignRegistFl bool
	invertFl    bool
	fixedFl     bool
	lineNumFl   bool
)

func main() {
	flag.IntVar(&afterFl, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&beforeFl, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&contextFl, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&countFl, "c", false, "подсчитать количество вхождений шаблона")
	flag.BoolVar(&ignRegistFl, "i", false, "игнорировать регистр")
	flag.BoolVar(&invertFl, "v", false, "инвертировать поиск, выдавать все строки кроме тех, что содержат шаблон")
	flag.BoolVar(&fixedFl, "F", false, "точное совпадение со строкой")
	flag.BoolVar(&lineNumFl, "n", false, "показывать номер строки в файле")
	flag.Parse()

	var in io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}
	sub := flag.Arg(1)

	buf := bufio.NewScanner(in)

	msg := fileRead(buf)

	g := &GrepFlags{
		after:     afterFl,
		before:    beforeFl,
		context:   contextFl,
		count:     countFl,
		ignRegist: ignRegistFl,
		invert:    invertFl,
		fixed:     fixedFl,
		lineNum:   lineNumFl,
		msg:       msg,
		sub:       sub,
		out:       os.Stdout,
	}
	g.Grep()
}
