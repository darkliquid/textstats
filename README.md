# textstats
Generate information about text including syllable counts and Flesch-Kincaid,
Gunning-Fog, Coleman-Liau, Dale-Chall, SMOG and Automated Readability scores.

Initially a more or less direct port of [TextStatistics.js][1] to Go, this 
supports analysing an io.Reader as well as strings.

[1]:https://github.com/cgiffard/TextStatistics.js
[2]:http://en.wikipedia.org/wiki/Dale%E2%80%93Chall_readability_formula
