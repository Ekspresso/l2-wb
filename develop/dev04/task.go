package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

// Написать функцию поиска всех множеств анаграмм по словарю.

// Например:
// 'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
// 'листок', 'слиток' и 'столик' - другому.

// Требования:
// Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
// Выходные данные: ссылка на мапу множеств анаграмм
// Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
// слово из множества.
// Массив должен быть отсортирован по возрастанию.
// Множества из одного элемента не должны попасть в результат.
// Все слова должны быть приведены к нижнему регистру.
// В результате каждое слово должно встречаться только один раз.

// createMapWord - Функция создания карты символов из слова. Ключ - символ, значение - кол-во повторений символа.
func createMapWord(s string) map[rune]int {
	word := make(map[rune]int) // карта для хранения букв в слове
	for i := 0; i < len(s); {
		r, w := utf8.DecodeRuneInString(s[i:])
		word[r]++
		i += w
	}
	return word
}

// isAnagram - Функция проверки является ли слово анаграммой
func isAnagram(s string, wordMap map[rune]int) bool {
	word := make(map[rune]int)
	for key, val := range wordMap {
		word[key] = val
	}
	for i := 0; i < len(s); {
		r, w := utf8.DecodeRuneInString(s[i:])
		word[r]--
		if word[r] < 0 {
			return false
		}
		i += w
	}
	for _, val := range word {
		if val != 0 {
			return false
		}
	}
	return true
}

// searchAnagrams - Функция поиска анаграмм в переданном массиве
func searchAnagrams(arr []string) map[string][]string {
	if arr == nil {
		return nil
	}
	ret := make(map[string][]string)   // карта для хранения результата
	uniqWords := make(map[string]bool) // карта для хранения уникальных слов

	for i := 0; i < len(arr); i++ {
		if !uniqWords[strings.ToLower(arr[i])] {
			anagrams := make([]string, 0)
			s := strings.ToLower(arr[i])
			wordMap := createMapWord(s)
			anagrams = append(anagrams, s)
			uniqWords[s] = true
			for j := i + 1; j < len(arr); j++ {
				checkStr := strings.ToLower(arr[j])
				if !uniqWords[checkStr] && isAnagram(checkStr, wordMap) {
					uniqWords[checkStr] = true
					anagrams = append(anagrams, checkStr)
				}
			}
			sort.SliceStable(anagrams, func(i, j int) bool { return anagrams[i] < anagrams[j] })
			ret[s] = anagrams
		}
	}
	return ret
}

func main() {
	arr := []string{"привет", "пукчек", "Чекпук", "аБв", "пуКчек", "авб", "вбА", "бав", "ваб", "бва", "пРивет", ""}
	fmt.Println(searchAnagrams(arr))
}
