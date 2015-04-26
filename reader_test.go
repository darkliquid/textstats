package textstats

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AnalyseSuite struct {
	suite.Suite
}

type badReader struct{}

func (b *badReader) Read(buf []byte) (int, error) {
	return 0, errors.New("fake error")
}

func (s *AnalyseSuite) TestBadReader() {
	_, err := Analyse(&badReader{})
	s.Error(err)
}

func (s *AnalyseSuite) TestAverageLettersPerWord() {
	res, _ := Analyse(strings.NewReader(qbf))
	s.Equal(3.888888888888889, res.AverageLettersPerWord())
}

func (s *AnalyseSuite) TestAverageSyllablesPerWord() {
	res, _ := Analyse(strings.NewReader(qbf))
	s.Equal(1.2222222222222223, res.AverageSyllablesPerWord())
}

func (s *AnalyseSuite) TestAverageWordsPerSentence() {
	res, _ := Analyse(strings.NewReader(qbf))
	s.Equal(9, res.AverageWordsPerSentence())

	res, _ = Analyse(strings.NewReader(lorem))
	s.Equal(17.25, res.AverageWordsPerSentence())
}

func (s *AnalyseSuite) TestWordsWithAtLeastNSyllables() {
	res, _ := Analyse(strings.NewReader(hw))
	s.Equal(6, res.WordsWithAtLeastNSyllables(0, true))
	s.Equal(6, res.WordsWithAtLeastNSyllables(1, true))
	s.Equal(3, res.WordsWithAtLeastNSyllables(2, true))
	s.Equal(2, res.WordsWithAtLeastNSyllables(3, true))
	s.Equal(1, res.WordsWithAtLeastNSyllables(4, true))
	s.Equal(0, res.WordsWithAtLeastNSyllables(5, true))

	s.Equal(4, res.WordsWithAtLeastNSyllables(0, false))
	s.Equal(4, res.WordsWithAtLeastNSyllables(1, false))
	s.Equal(2, res.WordsWithAtLeastNSyllables(2, false))
	s.Equal(2, res.WordsWithAtLeastNSyllables(3, false))
	s.Equal(1, res.WordsWithAtLeastNSyllables(4, false))
	s.Equal(0, res.WordsWithAtLeastNSyllables(5, false))
}

func (s *AnalyseSuite) TestPercentageWordsWithAtLeastNSyllables() {
	res, _ := Analyse(strings.NewReader(hw))
	s.Equal(100, res.PercentageWordsWithAtLeastNSyllables(0, true))
	s.Equal(100, res.PercentageWordsWithAtLeastNSyllables(1, true))
	s.Equal(50, res.PercentageWordsWithAtLeastNSyllables(2, true))
	s.Equal(33.33333333333333, res.PercentageWordsWithAtLeastNSyllables(3, true))
	s.Equal(16.666666666666664, res.PercentageWordsWithAtLeastNSyllables(4, true))
	s.Equal(0, res.PercentageWordsWithAtLeastNSyllables(5, true))

	s.Equal(66.66666666666666, res.PercentageWordsWithAtLeastNSyllables(0, false))
	s.Equal(66.66666666666666, res.PercentageWordsWithAtLeastNSyllables(1, false))
	s.Equal(33.33333333333333, res.PercentageWordsWithAtLeastNSyllables(3, false))
	s.Equal(33.33333333333333, res.PercentageWordsWithAtLeastNSyllables(2, false))
	s.Equal(16.666666666666664, res.PercentageWordsWithAtLeastNSyllables(4, false))
	s.Equal(0, res.PercentageWordsWithAtLeastNSyllables(5, false))
}

func (s *AnalyseSuite) TestWordCount() {
	res, _ := Analyse(strings.NewReader(qbf))
	s.Equal(9, res.Words)
}

func (s *AnalyseSuite) TestLetterCount() {
	res, _ := Analyse(strings.NewReader(qbf))
	s.Equal(35, res.Letters)
}

func (s *AnalyseSuite) TestSentenceCount() {
	res, _ := Analyse(strings.NewReader(lorem))
	s.Equal(4, res.Sentences)
}

func (s *AnalyseSuite) TestSyllableCount() {
	res, _ := Analyse(strings.NewReader(qbf))
	s.Equal(11, res.Syllables)
}

func (s *AnalyseSuite) TestFleschKincaidReadingEase() {
	res, _ := Analyse(strings.NewReader(lorem))
	s.Equal(7.865380434782622, res.FleschKincaidReadingEase())
}

func (s *AnalyseSuite) TestFleschKincaidGradeLevel() {
	res, _ := Analyse(strings.NewReader(lorem))
	s.Equal(16.447644927536235, res.FleschKincaidGradeLevel())
}

func (s *AnalyseSuite) TestGunningFogScore() {
	res, _ := Analyse(strings.NewReader(lorem))
	s.Equal(19.073913043478264, res.GunningFogScore())
}

func (s *AnalyseSuite) TestColemanLiauIndex() {
	res, _ := Analyse(strings.NewReader(lorem))
	s.Equal(15.681304347826085, res.ColemanLiauIndex())
}

func (s *AnalyseSuite) TestSMOGIndex() {
	res, _ := Analyse(strings.NewReader(lorem))
	s.Equal(13.524018386038225, res.SMOGIndex())
}

func (s *AnalyseSuite) TestAutomatedReadabilityIndex() {
	res, _ := Analyse(strings.NewReader(lorem))
	s.Equal(12.383260869565213, res.AutomatedReadabilityIndex())
}

func (s *AnalyseSuite) TestDaleChallReadabilityScore() {
	res, _ := Analyse(strings.NewReader(qbf))
	s.Equal(5.837344444444444, res.DaleChallReadabilityScore())
}

func TestAnalyseMethods(t *testing.T) {
	suite.Run(t, new(AnalyseSuite))
}
