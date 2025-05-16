// Text Analyzer - програма для аналізу тексту користувача.
// Функціонал: пошук слова, пошук слова за літерою, підрахунок слів та частоти, виведення найдовшого слова. Підтримує повторення аналізу.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func cleanWord(word string) string {
	// Видалення розділових знаків з початку і кінця слова
	return strings.TrimFunc(word, func(r rune) bool {
		return unicode.IsPunct(r)
	})
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// Запитування данних у користувача
	for {
		fmt.Println("Введіть текст для аналізу:")
		scanner.Scan()
		text := scanner.Text()
		// Якщо користувач нічого не ввів
		if strings.TrimSpace(text) == "" {
			fmt.Println("Текст не може бути порожнім.")
			continue
		}

		fmt.Println("Введіть слово для пошуку:")
		scanner.Scan()
		searchWord := strings.ToLower(scanner.Text())

		fmt.Println("Введіть літеру для пошуку першого слова:")
		scanner.Scan()
		letterInput := scanner.Text()

		if len(letterInput) == 0 {
			fmt.Println("Не введена літера.")
			continue
		}
		letter := []rune(strings.ToLower(letterInput))[0]

		// Аналіз тексту
		words := strings.Fields(text)
		totalWords := 0
		wordCount := 0
		firstWord := ""
		longestWord := ""

		for _, word := range words {
			cleaned := cleanWord(word)
			if cleaned == "" {
				continue
			}

			totalWords++

			if strings.ToLower(cleaned) == searchWord {
				wordCount++
			}

			if longestWord == "" || len(cleaned) > len(longestWord) {
				longestWord = cleaned
			}

			if firstWord == "" && strings.HasPrefix(strings.ToLower(cleaned), string(letter)) {
				firstWord = cleaned
			}
		}

		fmt.Printf("Слово \"%s\" зустрічається %d раз(и).\n", searchWord, wordCount)
		fmt.Printf("У тексті всього %d слів.\n", totalWords)

		if firstWord != "" {
			fmt.Printf("Перше слово, що починається на \"%c\": %s\n", letter, firstWord)
		} else {
			fmt.Printf("Слово, що починається на \"%c\", не знайдено.\n", letter)
		}

		fmt.Printf("Найдовше слово в тексті: %s\n", longestWord)

		// Повторення аналізу
		fmt.Println("Хочете проаналізувати інший текст? (yes/no)")
		scanner.Scan()
		answer := strings.ToLower(scanner.Text())

		switch answer {
		case "yes":
			continue
		case "no":
			fmt.Println("До побачення!")
			break
		default:
			fmt.Println("Незрозуміла відповідь. Завершення програми.")
			break
		}
		break
	}
}
