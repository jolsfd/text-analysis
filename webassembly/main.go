package main

import (
	"strings"
	"syscall/js"
)

type result struct {
	words     int
	sentences int
	syllables int
	score     float64
}

var terminals = map[rune]bool{
	'.': true,
	';': true,
	'?': true,
	'!': true,
}

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func getWords(chars []rune) [][]rune {
	words := [][]rune{}
	currentWord := []rune{}
	prevCharTerminal := false

	for _, char := range chars {
		// latin letter
		if 97 <= char && char <= 122 {
			currentWord = append(currentWord, char)
			prevCharTerminal = false
			continue
		}

		if char == 180 || char == 96 || char == 697 || char == 700 || char == 8217 { // accents
			prevCharTerminal = false
			continue
		}

		// words seperated by ' are counted as two words

		if !prevCharTerminal {
			words = append(words, currentWord)
			currentWord = []rune{}
		}

		prevCharTerminal = true
	}

	return words
}

func countSentences(chars []rune) (numOfSentences int) {
	for _, char := range chars {
		if _, ok := terminals[char]; ok {
			numOfSentences += 1
		}

	}

	return numOfSentences
}

func countSyllablesInWord(chars []rune) int {
	numOfSyllables := 0

	if len(chars) <= 3 {
		return 1
	}

	if chars[len(chars)-1] == 'e' {
		chars = chars[0 : len(chars)-1]
	}

	prevCharVowel := false

	for _, char := range chars {
		if _, ok := vowels[char]; ok {
			if !prevCharVowel {
				numOfSyllables += 1
			}

			prevCharVowel = true
			continue
		}

		prevCharVowel = false
	}

	if chars[len(chars)-1] == 'y' {
		numOfSyllables += 1
	}

	return numOfSyllables
}

func countSyllables(words [][]rune) (numOfSyllables int) {
	for _, word := range words {
		numOfSyllables += countSyllablesInWord(word)
	}

	return
}

func computeReadability(text string) result {
	chars := []rune(strings.ToLower(text))
	words := getWords(chars)

	totalWords := len(words)
	totalSyllables := countSyllables(words)
	totalSentences := countSentences(chars)

	if totalSentences == 0 || totalWords == 0 {
		return result{words: 0, sentences: 0, syllables: 0, score: -1}
	}

	//fmt.Println(totalSentences, totalWords, totalSyllables)

	score := 206.835 - 1.015*(float64(totalWords)/float64(totalSentences)) - 84.6*(float64(totalSyllables)/float64(totalWords))

	//fmt.Println(score)

	return result{words: totalWords, sentences: totalSentences, syllables: totalSyllables, score: score}
}

func computeReadabilityJS(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return "Invalid number of arguments were given"
	}

	rawText := args[0].String()
	res := computeReadability(rawText)
	//fmt.Println(res)

	return []interface{}{res.score, res.words, res.sentences, res.syllables}
}

func main() {
	js.Global().Set("computeReadability", js.FuncOf(computeReadabilityJS))

	<-make(chan bool)
}
