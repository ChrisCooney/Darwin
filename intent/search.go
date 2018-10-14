package intent

import (
	"Darwin/query"
	"strings"
)

type SearchIntent struct{}

func (intent SearchIntent) Name() string {
	return "Search Intent"
}

func (intent SearchIntent) Usage() string {
	return "Search for the speed of light"
}

func (intent SearchIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "search for") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), " search ") {
		return 50
	}

	return 0
}

func (intent SearchIntent) Response(input string) string {
	preparedInput := prepareSearchQuery(input)
	response := query.Search(preparedInput)

	if response == "" {
		return "Sorry. I couldn't find anything."
	}

	return response
}

func (intent SearchIntent) PostSpeech(input string) {
	// No op
}

func prepareSearchQuery(input string) string {
	input = strings.ToLower(input)
	input = strings.Replace(input, "darwin", "", 1)
	input = strings.Replace(input, "search for", "", 1)
	input = strings.Replace(input, ".", "", 1)
	input = strings.Trim(input, " ")
	input = strings.Trim(input, "\n")
	return input
}
