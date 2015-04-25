package textstats

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	qbf   = "The quick brown fox jumps over the lazy dog"
	lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
			eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim
			ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
			aliquip ex ea commodo consequat. Duis aute irure dolor in
			reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
			pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
			culpa qui officia deserunt mollit anim id est laborum.`
)

func TestWordCount(t *testing.T) {
	assert.Equal(t, 9, WordCount(qbf))
}

func TestLetterCount(t *testing.T) {
	assert.Equal(t, 35, LetterCount(qbf))
}

func TestSentenceCount(t *testing.T) {
	assert.Equal(t, 4, SentenceCount(lorem))
}

func TestSyllableCount(t *testing.T) {
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
		assert.Equal(t, count, SyllableCount(word), fmt.Sprintf("%q should have %d syllables", word, count))
	}
}

func TestFleschKincaidReadingEase(t *testing.T) {
	assert.Equal(t, 7.8653804347800005, FleschKincaidReadingEase(lorem))
}

func TestFleschKincaidGradeLevel(t *testing.T) {
	assert.Equal(t, 16.44764492754, FleschKincaidGradeLevel(lorem))
}

func TestGunningFogScore(t *testing.T) {
	assert.Equal(t, 19.073913043479997, GunningFogScore(lorem))
}

func TestColemanLiauIndex(t *testing.T) {
	assert.Equal(t, 15.681304347829998, ColemanLiauIndex(lorem))
}

func TestSMOGIndex(t *testing.T) {
	assert.Equal(t, 13.21893361077, SMOGIndex(lorem))
}

func TestAutomatedReadabilityIndex(t *testing.T) {
	assert.Equal(t, 12.38326086957, AutomatedReadabilityIndex(lorem))
}
