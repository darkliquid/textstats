package main

import (
	"fmt"
	"os"

	termutil "github.com/andrew-d/go-termutil"
	"github.com/darkliquid/textstats"
)

func printStats(name string, res *textstats.Results) {
	fmt.Printf("Statistics for %q:\n", name)
	fmt.Printf(`
	Words              %d
	Sentences          %d
	Letters            %d
	Punctuation        %d
	Spaces             %d
	Syllables          %d
	Difficult Words    %d
	Avg Letters/Word   %f
	Avg Syllables/Word %f
	Avg Words/Sentence %f

Readability Scores:
	Flesch-Kincaid Reading Ease  %f
	Flesch-Kincaid Grade Level   %f
	Gunning-Fog Score            %f
	Coleman-Liau Index           %f
	SMOG Index                   %f
	Automated Readability Index  %f
	Dale-Chall Readability Score %f

`,
		res.Words,
		res.Sentences,
		res.Letters,
		res.Punctuation,
		res.Spaces,
		res.Syllables,
		res.DifficultWords,
		res.AverageLettersPerWord(),
		res.AverageSyllablesPerWord(),
		res.AverageWordsPerSentence(),
		res.FleschKincaidReadingEase(),
		res.FleschKincaidGradeLevel(),
		res.GunningFogScore(),
		res.ColemanLiauIndex(),
		res.SMOGIndex(),
		res.AutomatedReadabilityIndex(),
		res.DaleChallReadabilityScore(),
	)
}

func main() {
	if !termutil.Isatty(os.Stdin.Fd()) {
		res, err := textstats.Analyse(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		printStats("STDIN", res)
		return
	}

	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Println(os.Args[0])
		fmt.Println("Usage:", os.Args[0], "[filename]")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	res, err := textstats.Analyse(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printStats(os.Args[1], res)
}
