package start

import (
	"fmt"
	"github.com/fatih/color"
	support "go-reloaded/pkg/functions"
	"strconv"
)

// Основная функция, которая запускает процесс обработки текста
func Run(inputFilePath, outputFilePath string) error {
	// Чтение данных из входного файла
	inputContent, err := support.ReadAndSave(inputFilePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения входного файла: %v", err)
	}
	replaceNewLine := support.ReplaceNewlines(inputContent, "\n", "§")
	replaceAn := support.ReplaceAWithAn(replaceNewLine)
	withCommand := support.ApplyCommands(replaceAn)
	unCommand := support.DeleteCommand(withCommand)
	fixQuotes := support.FixQuotes(unCommand)
	formatPuncuation := support.FormatPunctuation(fixQuotes)
	spaceAfterCharter := support.SpaceAfterCharter(formatPuncuation)
	deleteSpace := support.DeleteSpace(spaceAfterCharter)
	deleteSpaceStartAndFinish := support.DeleteSpaceStartAndFinish(deleteSpace)
	correctNewLine := support.CorrectNewLine(deleteSpaceStartAndFinish)
	replaceBackNewLine := support.ReplaceNewlines(correctNewLine, "§", "\n")
	removeEmptyLines := support.ReplaceSymbolToNewLine(replaceBackNewLine)
	correctComma := support.CorrectComma(removeEmptyLines)
	// Запись обработанного текста в выходной файл
	err = support.WriteToFile(outputFilePath, correctComma)
	if err != nil {
		return fmt.Errorf("ошибка записи в выходной файл: %v", err)
	}

	return nil
}

func Test() {
	count := 1
	//Test #1
	sample := "it (cap) was the best of times, it was the worst of times (up)"
	correctResult := "It was the best of times, it was the worst of TIMES"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #2
	sample = "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	correctResult = "Simply add 66 and 2 and you will see the result is 68."
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #3
	sample = "I am exactly how they describe me: ' awesome '"
	correctResult = "I am exactly how they describe me: 'awesome'"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #4
	sample = "There it was. A amazing rock!"
	correctResult = "There it was. An amazing rock!"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #5
	sample = "As Elton John said: ' I am the most well-known homosexual in the world '"
	correctResult = "As Elton John said: 'I am the most well-known homosexual in the world'"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #6
	sample = "Punctuation tests are ... kinda boring ,don't you think !?"
	correctResult = "Punctuation tests are... kinda boring, don't you think!?"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #7
	sample = "hello world word word2222 ! (up, 2)"
	correctResult = "hello world WORD WORD2222!"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #8
	sample = "word (cap, 100)"
	correctResult = "Word"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #9
	sample = "(cap)"
	correctResult = ""
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #10
	sample = "word (cap, 100) word, word (up, 2) word word (bin)"
	correctResult = "Word WORD, WORD word word"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #11
	sample = "word (cap) (low)"
	correctResult = "word"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #12
	sample = "12354 (bin)"
	correctResult = "12354"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #13
	sample = "Hello 10 (bin)"
	correctResult = "Hello 2"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #14
	sample = "1010 (bin)"
	correctResult = "10"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #15
	sample = "1234adbc (hex)"
	correctResult = "305442236"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #16
	sample = "1234adbcz (hex)"
	correctResult = "1234adbcz"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #17
	sample = "Sultan's        friend        told me:'     hello ?    ' (up, 9)"
	correctResult = "SULTAN'S FRIEND TOLD ME: 'HELLO?'"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #18
	sample = "123 (bin)"
	correctResult = "123"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #19
	sample = "(up)"
	correctResult = ""
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #20
	sample = "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
	correctResult = "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #21
	sample = "There is no greater agony than bearing a untold story inside you."
	correctResult = "There is no greater agony than bearing an untold story inside you."
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #22
	sample = "Punctuation tests are ... kinda boring ,don't you think !?"
	correctResult = "Punctuation tests are... kinda boring, don't you think!?"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #23
	sample = "There it was. A amazing rock!"
	correctResult = "There it was. An amazing rock!"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #24
	sample = "As Elton John said: ' I am " +
		"the most well-known " +
		"homosexual in the world '"
	correctResult = "As Elton John said: 'I am " +
		"the most well-known " +
		"homosexual in the world'"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #25
	sample = "I am exactly how " +
		"they describe me: ' awesome '"
	correctResult = "I am exactly how " +
		"they describe me: 'awesome'"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #26
	sample = "I was thinking ... You were right"
	correctResult = "I was thinking... You were right"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #27
	sample = "I was sitting over there ,and then BAMM !!"
	correctResult = "I was sitting over there, and then BAMM!!"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #28
	sample = "This is so exciting (up, 2)"
	correctResult = "This is SO EXCITING"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #29
	sample = "Welcome to the Brooklyn bridge (cap)"
	correctResult = "Welcome to the Brooklyn Bridge"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #30
	sample = "I should stop SHOUTING (low)"
	correctResult = "I should stop shouting"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #31
	sample = "Ready, set, go (up) !"
	correctResult = "Ready, set, GO!"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #32
	sample = "It has been 10 (bin) years"
	correctResult = "It has been 2 years"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #33
	sample = "1E (hex) files were added"
	correctResult = "30 files were added"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #34
	sample = "word (hex) "
	correctResult = "word"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #35
	sample = " a a a a apple a a"
	correctResult = "a a a an apple a a"
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

	//Test #35
	sample = "hello WORLD (cap)" +
		"(low)" +
		""
	correctResult = "hello world" +
		"" +
		""
	fmt.Println("Test #" + strconv.Itoa(count))
	check(sample, correctResult)
	count++

}

func check(sample, correctResult string) {
	replaceNewLine := support.ReplaceNewlines(sample, "\n", "§")
	replaceAn := support.ReplaceAWithAn(replaceNewLine)
	withCommand := support.ApplyCommands(replaceAn)
	unCommand := support.DeleteCommand(withCommand)
	fixQuotes := support.FixQuotes(unCommand)
	formatPuncuation := support.FormatPunctuation(fixQuotes)
	spaceAfterCharter := support.SpaceAfterCharter(formatPuncuation)
	deleteSpace := support.DeleteSpace(spaceAfterCharter)
	deleteSpaceStartAndFinish := support.DeleteSpaceStartAndFinish(deleteSpace)
	correctNewLine := support.CorrectNewLine(deleteSpaceStartAndFinish)
	replaceBackNewLine := support.ReplaceNewlines(correctNewLine, "§", "\n")
	removeEmptyLines := support.ReplaceSymbolToNewLine(replaceBackNewLine)
	correctComma := support.CorrectComma(removeEmptyLines)
	result := correctComma

	if result == correctResult {
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("%s\n", green("Passed"))
	} else {
		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		fmt.Printf("%s -> Result\n", red(result))
		fmt.Printf("%s -> Correct\n", green(correctResult))
	}
}
