// package dev02

package main

// Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
// "a4bc2d5e" => "aaaabccddddde"
// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "" => ""

// Дополнительно
// Реализовать поддержку escape-последовательностей.
// Например:
// qwe\4\5 => qwe45 (*)
// qwe\45 => qwe44444 (*)
// qwe\\5 => qwe\\\\\ (*)

// В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// ErrInvalid - глобальная переменная для некорректной строки для возвращения из функции
var ErrInvalid = fmt.Errorf("invalid string")

// UnpackStr - Go-функция, осуществляющая примитивную распаковку строки, содержащую повторяющиеся символы/руны.
func UnpackStr(str string) (string, error) {
	// Если строка пустая, то вернуть пустую строку без ошибки.
	if len(str) == 0 {
		return "", nil
	}

	firstRun, width := utf8.DecodeRuneInString(str)
	// Если первый символ - число, то вернуть ошибку и пустую строку.
	// Если есть только символ обратного слэша, то вернуть ошибку и пустую строку.
	if firstRun >= '0' && firstRun <= '9' || len(str) == width && firstRun == '\\' {
		return "", ErrInvalid
	}

	w := width
	backSlash := false
	printBackSlash := false
	retStr := ""
	// Проверка на обратный слэш. Если его нет в считанном значении, то символ добавляется в строку.
	if firstRun != '\\' {
		retStr = concat(retStr, string(firstRun))
	}
	if firstRun == '\\' {
		backSlash = true
		firstRun, width = utf8.DecodeRuneInString(str[w:])
		w += width
		retStr = concat(retStr, string(firstRun))
	}

	// Цикл считывания переданной строки.
	for i, w := w, 0; i < len(str); i += w {
		secondRun, width := utf8.DecodeRuneInString(str[i:])
		w = width

		// Проверка на наличие лишнего \ в конце строки, на наличие нулевого счётчика символов и повторения чисел без \.
		if !backSlash && !printBackSlash && secondRun == '\\' && i+w == len(str) || firstRun != '\\' && secondRun == 0 ||
			!backSlash && secondRun >= '0' && secondRun <= '9' && firstRun >= '0' && firstRun <= '9' {
			return "", ErrInvalid
		} else if !backSlash && secondRun == '\\' {
			backSlash = true
		} else if backSlash && firstRun == '\\' && !printBackSlash {
			retStr = concat(retStr, string(secondRun))
			if secondRun == '\\' {
				printBackSlash = true
			}
		} else if secondRun >= '0' && secondRun <= '9' {
			k, _ := strconv.Atoi(string(secondRun))
			for i := 0; i < k-1; i++ {
				retStr = concat(retStr, string(firstRun))
			}
			if firstRun == '\\' {
				printBackSlash = false
			}
			backSlash = false
		} else if backSlash && printBackSlash {
			printBackSlash = false
		} else if backSlash && secondRun == '\\' {
		} else {
			backSlash = false
			retStr = concat(retStr, string(secondRun))
		}
		if printBackSlash && secondRun == '\\' && i+w == len(str) {
			return "", ErrInvalid
		}
		firstRun = secondRun
	}
	return retStr, nil
}

// concat выполняет эффективную конкатенацию строк
func concat(x, y string) string {
	var builder strings.Builder
	builder.Grow(len(x) + len(y)) // Эта строка выделяет память
	builder.WriteString(x)        //Записывает в builder строку.
	builder.WriteString(y)
	return builder.String()
}

func main() {
	fmt.Println(UnpackStr(`qw\\\e`))
}
