# textstats [![Go Report Card](https://goreportcard.com/badge/github.com/darkliquid/textstats)](https://goreportcard.com/report/github.com/darkliquid/textstats) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/darkliquid/textstats/blob/master/LICENSE) [![GoDoc](https://godoc.org/github.com/darkliquid/textstats?status.svg)](https://godoc.org/github.com/darkliquid/textstats) [![Build Status](https://travis-ci.org/darkliquid/textstats.svg?branch=master)](https://travis-ci.org/darkliquid/textstats)

Generate information about text including syllable counts and Flesch-Kincaid,
Gunning-Fog, Coleman-Liau, Dale-Chall, SMOG and Automated Readability scores.

Initially a more or less direct port of [TextStatistics.js][1] to Go, this
supports analysing an io.Reader as well as strings.

[1]:https://github.com/cgiffard/TextStatistics.js
