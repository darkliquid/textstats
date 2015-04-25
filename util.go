package textstats

import "math"

func wordsWithThreeSyllables(text string, countProperNouns bool) (count int) {
	for _, word := range whitespaceRegexp.Split(text, -1) {
		if !capitalRegexp.MatchString(word) || countProperNouns {
			if SyllableCount(word) > 2 {
				count++
			}
		}
	}

	return
}

func percentageWordsWithThreeSyllables(text string, countProperNouns bool) float64 {
	count := float64(wordsWithThreeSyllables(text, countProperNouns))
	return count / float64(WordCount(text)) * 100.0
}

func averageWordsPerSentence(text string) float64 {
	return float64(WordCount(text)) / float64(SentenceCount(text))
}

func averageSyllablesPerWord(text string) float64 {
	var sCount int
	for _, word := range whitespaceRegexp.Split(text, -1) {
		sCount += SyllableCount(word)
	}

	if sCount == 0 {
		sCount = 1
	}

	return float64(sCount) / float64(WordCount(text))
}

func roundTen(f float64) float64 {
	shift := math.Pow(10, float64(10))
	return math.Floor((f*shift)+0.5) / shift
}
