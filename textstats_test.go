package textstats

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StringSuite struct {
	suite.Suite
}

func (s *StringSuite) TestAverageLettersPerWord() {
	s.Equal(3.888888888888889, AverageLettersPerWord(qbf))
}

func (s *StringSuite) TestAverageSyllablesPerWord() {
	s.Equal(1.2222222222222223, AverageSyllablesPerWord(qbf))
}

func (s *StringSuite) TestAverageWordsPerSentence() {
	s.Equal(9, AverageWordsPerSentence(qbf))
	s.Equal(17.25, AverageWordsPerSentence(lorem))
}

func (s *StringSuite) TestWordsWithAtLeastNSyllables() {
	s.Equal(6, WordsWithAtLeastNSyllables(hw, 0, true))
	s.Equal(6, WordsWithAtLeastNSyllables(hw, 1, true))
	s.Equal(3, WordsWithAtLeastNSyllables(hw, 2, true))
	s.Equal(2, WordsWithAtLeastNSyllables(hw, 3, true))
	s.Equal(1, WordsWithAtLeastNSyllables(hw, 4, true))
	s.Equal(0, WordsWithAtLeastNSyllables(hw, 5, true))

	s.Equal(4, WordsWithAtLeastNSyllables(hw, 0, false))
	s.Equal(4, WordsWithAtLeastNSyllables(hw, 1, false))
	s.Equal(2, WordsWithAtLeastNSyllables(hw, 2, false))
	s.Equal(2, WordsWithAtLeastNSyllables(hw, 3, false))
	s.Equal(1, WordsWithAtLeastNSyllables(hw, 4, false))
	s.Equal(0, WordsWithAtLeastNSyllables(hw, 5, false))
}

func (s *StringSuite) TestPercentageWordsWithAtLeastNSyllables() {
	s.Equal(100, PercentageWordsWithAtLeastNSyllables(hw, 0, true))
	s.Equal(100, PercentageWordsWithAtLeastNSyllables(hw, 1, true))
	s.Equal(50, PercentageWordsWithAtLeastNSyllables(hw, 2, true))
	s.Equal(33.33333333333333, PercentageWordsWithAtLeastNSyllables(hw, 3, true))
	s.Equal(16.666666666666664, PercentageWordsWithAtLeastNSyllables(hw, 4, true))
	s.Equal(0, PercentageWordsWithAtLeastNSyllables(hw, 5, true))

	s.Equal(66.66666666666666, PercentageWordsWithAtLeastNSyllables(hw, 0, false))
	s.Equal(66.66666666666666, PercentageWordsWithAtLeastNSyllables(hw, 1, false))
	s.Equal(33.33333333333333, PercentageWordsWithAtLeastNSyllables(hw, 3, false))
	s.Equal(33.33333333333333, PercentageWordsWithAtLeastNSyllables(hw, 2, false))
	s.Equal(16.666666666666664, PercentageWordsWithAtLeastNSyllables(hw, 4, false))
	s.Equal(0, PercentageWordsWithAtLeastNSyllables(hw, 5, false))
}

func (s *StringSuite) TestWordCount() {
	s.Equal(9, WordCount(qbf))
}

func (s *StringSuite) TestLetterCount() {
	s.Equal(35, LetterCount(qbf))
}

func (s *StringSuite) TestSentenceCount() {
	s.Equal(4, SentenceCount(lorem))
}

func (s *StringSuite) TestSyllableCount() {
	words := map[string]int{
		"advertisement": 4,
		"bath":          1,
		"data":          2,
		"direct":        2,
		"diverse":       2,
		"due":           1,
		"economics":     4,
		"either":        2,
		"employee":      3,
		"exquisite":     3,
		"finance":       2,
		"forest":        2,
		"forever":       3,
		"glass":         1,
		"herb":          1,
		"juvenile":      3,
		"kilometer":     4,
		"laugh":         1,
		"neither":       2,
		"orange":        2,
		"pajamas":       3,
		"past":          1,
		"roof":          1,
		"route":         1,
		"shoreline":     2,
		"simile":        3,
		"tomatoe":       3,
		"volatile":      3,
		"when":          1,
		"why":           1,
	}

	for word, count := range words {
		s.Equal(count, SyllableCount(word), fmt.Sprintf("%q should have %d syllables", word, count))
	}
}

func (s *StringSuite) TestFleschKincaidReadingEase() {
	s.Equal(7.865380434782622, FleschKincaidReadingEase(lorem))
}

func (s *StringSuite) TestFleschKincaidGradeLevel() {
	s.Equal(16.447644927536235, FleschKincaidGradeLevel(lorem))
}

func (s *StringSuite) TestGunningFogScore() {
	s.Equal(19.073913043478264, GunningFogScore(lorem))
}

func (s *StringSuite) TestColemanLiauIndex() {
	s.Equal(15.681304347826085, ColemanLiauIndex(lorem))
}

func (s *StringSuite) TestSMOGIndex() {
	s.Equal(13.524018386038225, SMOGIndex(lorem))
}

func (s *StringSuite) TestAutomatedReadabilityIndex() {
	s.Equal(12.383260869565213, AutomatedReadabilityIndex(lorem))
}

func (s *StringSuite) TestDaleChallReadabilityScore() {
	s.Equal(5.837344444444444, DaleChallReadabilityScore(qbf))
}

func TestStringMethods(t *testing.T) {
	suite.Run(t, new(StringSuite))
}
