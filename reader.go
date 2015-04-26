package textstats

import (
	"bufio"
	"io"
	"math"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Results is a struct containing the results of an analysis
type Results struct {
	Words       int
	Sentences   int
	Letters     int
	Punctuation int
	Spaces      int
	Syllables   int

	syllableProperNouns map[int]int
	syllableWords       map[int]int
}

// AverageLettersPerWord returns the average number of letters per word in the
// text
func (r *Results) AverageLettersPerWord() float64 {
	return float64(r.Letters) / float64(r.Words)
}

// AverageSyllablesPerWord returns the average number of syllables per word in
// the text
func (r *Results) AverageSyllablesPerWord() float64 {
	return float64(r.Syllables) / float64(r.Words)
}

// AverageWordsPerSentence returns the avergae number of words per sentence in
// the text
func (r *Results) AverageWordsPerSentence() float64 {
	if r.Sentences == 0 {
		return float64(r.Words)
	}
	return float64(r.Words) / float64(r.Sentences)
}

// WordsWithAtLeastNSyllables returns the number of words with at least N
// syllables, including or excluding proper nouns, in the text
func (r *Results) WordsWithAtLeastNSyllables(n int, incProperNouns bool) int {
	var total int
	for sCount, wCount := range r.syllableWords {
		if sCount >= n {
			total += wCount
		}
	}

	if !incProperNouns {
		for sCount, wCount := range r.syllableProperNouns {
			if sCount >= n {
				total -= wCount
			}
		}
	}

	if total < 0 {
		return 0
	}

	return total
}

// PercentageWordsWithAtLeastNSyllables returns the percentage of words with at
// least N syllables, including or excluding proper nouns, in the text
func (r *Results) PercentageWordsWithAtLeastNSyllables(n int, incProperNouns bool) float64 {
	return (float64(r.WordsWithAtLeastNSyllables(n, incProperNouns)) / float64(r.Words)) * 100.0
}

// FleschKincaidReadingEase returns the Flesch-Kincaid reading ease score for
// given text
func (r *Results) FleschKincaidReadingEase() float64 {
	return 206.835 - (1.015 * r.AverageWordsPerSentence()) - (84.6 * r.AverageSyllablesPerWord())
}

// FleschKincaidGradeLevel returns the Flesch-Kincaid grade level for the given text
func (r *Results) FleschKincaidGradeLevel() float64 {
	return (0.39 * r.AverageWordsPerSentence()) + (11.8 * r.AverageSyllablesPerWord()) - 15.59
}

// GunningFogScore returns the Gunning-Fog score for the given text
func (r *Results) GunningFogScore() float64 {
	return (r.AverageWordsPerSentence() + r.PercentageWordsWithAtLeastNSyllables(3, false)) * 0.4
}

// ColemanLiauIndex returns the Coleman-Liau index for the given text
func (r *Results) ColemanLiauIndex() float64 {
	return (5.89 * (float64(r.Letters) / float64(r.Words))) - (0.3 * (float64(r.Sentences) / float64(r.Words))) - 15.8
}

// SMOGIndex returns the SMOG index for the given text
func (r *Results) SMOGIndex() float64 {
	return 1.0430 * math.Sqrt((float64(r.WordsWithAtLeastNSyllables(3, true))*(30/float64(r.Sentences)))+3.1291)
}

// AutomatedReadabilityIndex returns the Automated Readability index for the given text
func (r *Results) AutomatedReadabilityIndex() float64 {
	return (4.71 * (float64(r.Letters) / float64(r.Words))) + (0.5 * (float64(r.Words) / float64(r.Sentences))) - 21.43
}

func syllableCount(word string) (sCount int) {
	word = strings.ToLower(word)

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
	for _, wordPart := range consonantsRegexp.Split(word, -1) {
		if len(wordPart) > 0 {
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

func analyseWord(word string, res *Results) {
	res.Words++

	sCount := syllableCount(word)
	res.Syllables += sCount

	if _, ok := res.syllableWords[sCount]; ok {
		res.syllableWords[sCount]++
	} else {
		res.syllableWords[sCount] = 1
	}

	if l, _ := utf8.DecodeRuneInString(word); unicode.IsUpper(l) {
		if _, ok := res.syllableProperNouns[sCount]; ok {
			res.syllableProperNouns[sCount]++
		} else {
			res.syllableProperNouns[sCount] = 1
		}
	}
}

// Analyse scans a reader and outputs an analysis
func Analyse(r io.Reader) (res *Results, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	res = &Results{}
	res.syllableWords = make(map[int]int)
	res.syllableProperNouns = make(map[int]int)

	var word string
	var endWord bool
	for scanner.Scan() {
		str := scanner.Text()
		letter, _ := utf8.DecodeRuneInString(str)
		switch {
		case unicode.IsLetter(letter):
			res.Letters++
			word += str
			endWord = false
		case unicode.IsSpace(letter):
			endWord = true
			res.Spaces++
		case unicode.IsPunct(letter):
			endWord = true
			switch str {
			case ".", "!", "?":
				res.Sentences++
			}
			res.Punctuation++
		}

		if endWord && len(word) > 0 {
			analyseWord(word, res)
			endWord = false
			word = ""
		}
	}

	if len(word) > 0 {
		analyseWord(word, res)
	}

	// Return scanner error if any
	err = scanner.Err()

	return
}
