package textstats

import (
	"math"
	"strings"
)

// WordCount returns the number of words in a given string
func WordCount(text string) int {
	return len(wordsRegexp.FindAllString(text, -1))
}

// SentenceCount returns the number of sentences in a given string
func SentenceCount(text string) int {
	return len(sentencesRegexp.ReplaceAllString(text, ""))
}

// LetterCount returns the number of letters in a given string
func LetterCount(text string) int {
	return len(lettersRegexp.ReplaceAllString(text, ""))
}

// SyllableCount returns the number of syllables in a word
func SyllableCount(word string) (sCount int) {
	word = letterRegexp.ReplaceAllString(strings.ToLower(word), "")

	// return early if we have a problem word
	sCount, ok := ProblemWords[word]
	if ok {
		return
	}

	var prefixSuffixCount int
	for _, regex := range PrefixSuffixes[:] {
		if regex.MatchString(word) {
			word = regex.ReplaceAllString(word, "")
			prefixSuffixCount++
		}
	}

	var wordPartCount int
	for _, wordPart := range consontantsRegexp.Split(word, -1) {
		if len(whitespaceRegexp.ReplaceAllString(wordPart, "")) > 0 {
			wordPartCount++
		}
	}

	sCount = wordPartCount + prefixSuffixCount

	for _, regex := range SubSyllables[:] {
		if regex.MatchString(word) {
			sCount--
		}
	}

	for _, regex := range AddSyllables[:] {
		if regex.MatchString(word) {
			sCount++
		}
	}

	return
}

// FleschKincaidReadingEase returns the Flesch-Kincaid reading ease score for
// given text
func FleschKincaidReadingEase(text string) float64 {
	return roundTen((206.835-(1.015*averageWordsPerSentence(text))-(84.6*averageSyllablesPerWord(text)))*10.0) / 10.0
}

// FleschKincaidGradeLevel returns the Flesch-Kincaid grade level for the given text
func FleschKincaidGradeLevel(text string) float64 {
	return roundTen(((0.39*averageWordsPerSentence(text))+(11.8*averageSyllablesPerWord(text))-15.59)*10.0) / 10.0
}

// GunningFogScore returns the Gunning-Fog score for the given text
func GunningFogScore(text string) float64 {
	return roundTen(((averageWordsPerSentence(text)+percentageWordsWithThreeSyllables(text, false))*0.4)*10.0) / 10.0
}

// ColemanLiauIndex returns the Coleman-Liau index for the given text
func ColemanLiauIndex(text string) float64 {
	return roundTen(((5.89*(float64(LetterCount(text))/float64(WordCount(text))))-(0.3*(float64(SentenceCount(text))/float64(WordCount(text))))-15.8)*10.0) / 10.0
}

// SMOGIndex returns the SMOG index for the given text
func SMOGIndex(text string) float64 {
	return roundTen(1.043*math.Sqrt((float64(wordsWithThreeSyllables(text, false))*(30/float64(SentenceCount(text))))+3.1291)*10.0) / 10.0
}

// AutomatedReadabilityIndex returns the Automated Readability index for the given text
func AutomatedReadabilityIndex(text string) float64 {
	return roundTen(((4.71*(float64(LetterCount(text))/float64(WordCount(text))))+(0.5*(float64(WordCount(text))/float64(SentenceCount(text))))-21.43)*10.0) / 10.0
}
