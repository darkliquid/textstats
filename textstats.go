package textstats

import "strings"

// AverageLettersPerWord returns the average number of letters per word in the
// text
func AverageLettersPerWord(text string) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.AverageLettersPerWord()
}

// AverageSyllablesPerWord returns the average number of syllables per word in
// the text
func AverageSyllablesPerWord(text string) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.AverageSyllablesPerWord()
}

// AverageWordsPerSentence returns the avergae number of words per sentence in
// the text
func AverageWordsPerSentence(text string) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.AverageWordsPerSentence()
}

// WordsWithAtLeastNSyllables returns the number of words with at least N
// syllables, including or excluding proper nouns, in the text
func WordsWithAtLeastNSyllables(text string, n int, incProperNouns bool) int {
	res, _ := Analyse(strings.NewReader(text))
	return res.WordsWithAtLeastNSyllables(n, incProperNouns)
}

// PercentageWordsWithAtLeastNSyllables returns the percentage of words with at
// least N syllables, including or excluding proper nouns, in the text
func PercentageWordsWithAtLeastNSyllables(text string, n int, incProperNouns bool) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.PercentageWordsWithAtLeastNSyllables(n, incProperNouns)
}

// WordCount returns the number of words in a given string
func WordCount(text string) int {
	res, _ := Analyse(strings.NewReader(text))
	return res.Words
}

// SentenceCount returns the number of sentences in a given string
func SentenceCount(text string) int {
	res, _ := Analyse(strings.NewReader(text))
	return res.Sentences
}

// LetterCount returns the number of letters in a given string
func LetterCount(text string) int {
	res, _ := Analyse(strings.NewReader(text))
	return res.Letters
}

// SyllableCount returns the number of syllables in a given string
func SyllableCount(text string) int {
	res, _ := Analyse(strings.NewReader(text))
	return res.Syllables
}

// FleschKincaidReadingEase returns the Flesch-Kincaid reading ease score for
// given text
func FleschKincaidReadingEase(text string) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.FleschKincaidReadingEase()
}

// FleschKincaidGradeLevel returns the Flesch-Kincaid grade level for the given text
func FleschKincaidGradeLevel(text string) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.FleschKincaidGradeLevel()
}

// GunningFogScore returns the Gunning-Fog score for the given text
func GunningFogScore(text string) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.GunningFogScore()
}

// ColemanLiauIndex returns the Coleman-Liau index for the given text
func ColemanLiauIndex(text string) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.ColemanLiauIndex()
}

// SMOGIndex returns the SMOG index for the given text
func SMOGIndex(text string) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.SMOGIndex()
}

// AutomatedReadabilityIndex returns the Automated Readability index for the given text
func AutomatedReadabilityIndex(text string) float64 {
	res, _ := Analyse(strings.NewReader(text))
	return res.AutomatedReadabilityIndex()
}
