package functions

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadAndSave(filePath string) (string, error) {
	// Чтение данных из файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	fileContetnt := string(data)
	return fileContetnt, nil
}

func WriteToFile(filePath string, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Функция для замены шестнадцатеричных чисел на десятичные
func HexToDecimal(hex string) (string, error) {
	decimal := new(big.Int)
	decimal, success := decimal.SetString(hex, 16)
	if !success {
		return "", fmt.Errorf("невозможно преобразовать шестнадцатеричное число %s в десятичное", hex)
	}
	return decimal.String(), nil
}

// Функция для замены двоичных чисел на десятичные
func BinaryToDecimal(binary string) (int64, error) {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return 0, err
	}
	return decimal, nil
}

func ReplaceNewlines(text, search, replacement string) string {
	// Создаем регулярное выражение для поиска последовательностей "\n"
	re := regexp.MustCompile(search)
	// Находим все вхождения последовательностей "\n"
	matches := re.FindAllStringIndex(text, -1)
	// Создаем новую строку для результата
	var replacedText strings.Builder
	// Индекс начала текущего вхождения
	startIdx := 0
	// Перебираем все найденные вхождения
	for _, match := range matches {
		// Получаем подстроку между текущим и предыдущим вхождением
		substr := text[startIdx:match[0]]
		// Заменяем все вхождения "\n" на "§" и добавляем в результирующую строку
		replacedText.WriteString(strings.ReplaceAll(substr, search, replacement))
		// Добавляем соответствующее количество "§" в зависимости от количества найденных вхождений
		replacedText.WriteString(strings.Repeat(replacement, match[1]-match[0]))
		// Обновляем индекс начала текущего вхождения
		startIdx = match[1]
	}
	// Добавляем оставшуюся часть исходной строки
	replacedText.WriteString(strings.ReplaceAll(text[startIdx:], search, replacement))
	return replacedText.String()
}

// Функция для применения команд к тексту
func ApplyCommands(content string) string {
	// Регулярное выражение для поиска команд в тексте
	re := regexp.MustCompile(`((\w+)([\s[:punct:]\s]*)[^\d[:punct:]]?)((|[А-Яа-я()\[\][:punct:]]+)(\s*)\((cap|up|low|hex|bin)(?:, (-?\d+))?\))`)

	// Функция для обработки найденных команд
	replacer := func(match string) string {
		// Извлекаем команду и ее аргументы из совпадения
		matches := re.FindStringSubmatch(match)
		if len(matches) < 7 {
			// Если не удалось извлечь команду, возвращаем исходную строку
			return match
		}
		wordWithSymbol := matches[1]
		word := matches[2]
		symbol := matches[3]
		command := matches[7]
		argStr := matches[8]

		// Если аргумент указан, преобразуем его в число
		var arg int
		if argStr != "" {
			arg, _ = strconv.Atoi(argStr)
			if arg <= 0 {
				return word
			}
			arg--
		}
		commandWithArg := ""
		if arg > 0 {
			commandWithArg = "(" + command + "," + " " + strconv.Itoa(arg) + ")"

		}

		// Применяем команду к слову в зависимости от ее типа
		switch command {
		case "up":
			if arg > 0 && argStr != "" {
				if symbol == "'" && len(symbol) == 1 {
					return commandWithArg + " " + strings.ToUpper(wordWithSymbol)
				}
				return commandWithArg + " " + strings.ToUpper(word) + symbol
			}
			if symbol == "'" && len(symbol) == 1 {
				return commandWithArg + " " + strings.ToUpper(wordWithSymbol)
			}
			return strings.ToUpper(word) + symbol
		case "low":
			if arg > 0 && argStr != "" {
				if symbol == "'" && len(symbol) == 1 {
					return commandWithArg + " " + strings.ToLower(wordWithSymbol)
				}
				return commandWithArg + " " + strings.ToLower(word) + symbol
			}
			if symbol == "'" && len(symbol) == 1 {
				return commandWithArg + " " + strings.ToLower(wordWithSymbol)
			}
			return strings.ToLower(word) + symbol
		case "cap":
			if arg > 0 {
				return commandWithArg + " " + strings.Title(strings.ToLower(word)) + symbol
				if symbol == "'" && len(symbol) == 1 {
					return commandWithArg + " " + strings.Title(strings.ToLower(wordWithSymbol))
				}
			}
			if symbol == "'" && len(symbol) == 1 {
				return commandWithArg + " " + strings.Title(strings.ToLower(wordWithSymbol))
			}
			return strings.Title(strings.ToLower(word)) + symbol
		case "hex":
			hexNum, err := HexToDecimal(word)
			if err != nil {
				return word
			}
			return hexNum + symbol
		case "bin":
			decimalNum, err := BinaryToDecimal(word)
			if err != nil {
				return word
			}
			return strconv.FormatInt(decimalNum, 10)
		default:
			return match + symbol
		}
	}

	// Применяем функцию замены ко всем совпадениям
	result := re.ReplaceAllStringFunc(content, replacer)
	if re.MatchString(result) {
		result = ApplyCommands(result)
	}
	return result
}

func DeleteCommand(content string) string {
	commandForDel := regexp.MustCompile(`\((cap|up|low|CAP|UP|LOW|Cap|Up|Low)(?:,\s*(\d+))?\)`)
	content = commandForDel.ReplaceAllString(content, "")
	// Заменяем все последовательности пробелов, кроме пробелов между переносами строк, на один пробел.
	content = regexp.MustCompile(`(?m)[^\S\n]+`).ReplaceAllString(content, " ")

	return content
}

func ReplaceAWithAn(text string) string {
	re := regexp.MustCompile(`\b([Aa])\s+([aeiouhAEIOUH]\w+)`)
	replacer := func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 2 {
			return match
		}
		wordA := matches[1]
		word := matches[2]
		if wordA == "A" {
			correctWord := "An" + " " + word
			return correctWord
		} else {
			correctWord := "an" + " " + word
			return correctWord
		}
	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

func FixQuotes(text string) string {
	// Регулярное выражение для добавления пробела после запятой, если его нет
	re := regexp.MustCompile(`([^\w]+)' *([^']+?) *'`)
	// Функция для обработки найденных команд
	replacer := func(match string) string {
		// Извлекаем команду и ее аргументы из совпадения
		matches := re.FindStringSubmatch(match)
		if len(matches) < 2 {
			// Если не удалось извлечь команду, возвращаем исходную строку
			return match
		}
		symbol := matches[1]
		word := matches[2]
		newWord := symbol + "'" + word + "'"
		return newWord

	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

func FormatPunctuation(text string) string {
	// Регулярное выражение для добавления пробела после запятой, если его нет
	re := regexp.MustCompile(`(\w+)(\s*)([:,;.!?]+)(?:\s*)(\w*)`)
	// Функция для обработки найденных команд
	replacer := func(match string) string {
		// Извлекаем команду и ее аргументы из совпадения
		matches := re.FindStringSubmatch(match)
		if len(matches) < 4 {
			// Если не удалось извлечь команду, возвращаем исходную строку
			return match
		}
		word_01 := matches[1]
		space := matches[2]
		symbol := matches[3]
		word_02 := matches[4]
		if word_02 != "" {
			correctWord := word_01 + symbol + " " + word_02
			return correctWord
		} else if word_01 != "" && space == " " && symbol != "" && word_02 == "" {
			correctWord := word_01 + symbol
			return correctWord
		} else {
			correctWord := match
			return correctWord
		}
	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

func SpaceAfterCharter(text string) string {
	re := regexp.MustCompile(`(\:+|\,+)`)
	// Функция для обработки найденных команд
	replacer := func(match string) string {
		// Извлекаем команду и ее аргументы из совпадения
		matches := re.FindStringSubmatch(match)
		if len(matches) < 1 {
			// Если не удалось извлечь команду, возвращаем исходную строку
			return match
		}
		symbol := matches[1]
		correctWord := symbol + " "
		return correctWord

	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

func DeleteSpace(text string) string {
	re := regexp.MustCompile(`( +)`)
	// Функция для обработки найденных команд
	replacer := func(match string) string {
		// Извлекаем команду и ее аргументы из совпадения
		matches := re.FindStringSubmatch(match)
		if len(matches) < 1 {
			// Если не удалось извлечь команду, возвращаем исходную строку
			return match
		}
		correctWord := " "
		return correctWord

	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

func DeleteSpaceStartAndFinish(text string) string {
	re := regexp.MustCompile(`(^\s+|\s+$)`)
	// Функция для обработки найденных пробелов
	replacer := func(match string) string {
		// Заменяем найденные пробелы на пустую строку
		return ""
	}
	// Заменяем пробелы в начале и в конце строки на пустую строку
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

func CorrectNewLine(text string) string {
	re := regexp.MustCompile(`(\ *\n+\ *)`)
	replacer := func(match string) string {
		return "\n"
	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

func ReplaceSymbolToNewLine(text string) string {
	// Создаем регулярное выражение для поиска пустых строк
	re := regexp.MustCompile(`([\w[:punct:]]*)(\n+)`)
	replacer := func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 2 {
			return match
		}
		group_01 := matches[1]
		group_02 := matches[2]
		newGrou_02 := ""
		if len(group_02) > 1 {
			for i := 0; i < len(group_02)/2; i++ {
				newGrou_02 += "\n"
			}
		}
		correctWord := group_01 + newGrou_02
		return correctWord

	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}

func CorrectComma(text string) string {
	re := regexp.MustCompile(`( +)(\,|\.|\:|\!|\?|\%|\$|\@)`)
	replacer := func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 2 {
			return match
		}

		symbol := matches[2]
		correctWord := symbol
		return correctWord

	}
	result := re.ReplaceAllStringFunc(text, replacer)
	return result
}
