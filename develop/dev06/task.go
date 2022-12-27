package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// Реализовать утилиту аналог консольной команды cut (man cut).
// Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

// Реализовать поддержку утилитой следующих ключей:
// -f - "fields" - выбрать поля (колонки)
// -d - "delimiter" - использовать другой разделитель
// -s - "separated" - только строки с разделителем

// Cuter - структура, хранящая флаги, стролбцы.
type Cuter struct {
	sl        []string
	Fields    []string
	Delim     string
	Separated bool
	Total     string
}

func (c *Cuter) split(text string) []string {
	return strings.Split(text, c.Delim)
}

// Cut - основная реализация функции cut
func (c *Cuter) Cut(text string) string {

	c.sl = c.split(text)
	// Если не нашлись разделители
	if len(c.sl) <= 1 {
		// Не выводим строки если нет разделителя
		if c.Separated {
			return ""
		}
		c.Total = c.sl[0]
		return c.Total
	}

	for _, v := range c.Fields {
		j, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}

		if len(c.sl)-1 > j {
			j--
			if j < 0 {
				j = 0
			}
			c.Total += c.sl[j] + " "
		}
	}

	return c.Total
}

func main() {
	var fields = flag.String("f", "", "выбрать поля (колонки)")
	var delimiter = flag.String("d", "\t", "использовать другой разделитель")
	var separated = flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()
	text := flag.Arg(0)

	c := Cuter{
		Fields:    strings.Split(*fields, ","),
		Delim:     *delimiter,
		Separated: *separated,
	}

	res := c.Cut(text)
	fmt.Println(res)
}
