package intent

import (
	"Darwin/news"
	"strings"
)

type NewsListIntent struct{}

func (intent NewsListIntent) Name() string {
	return "News List Intent"
}

func (intent NewsListIntent) Usage() string {
	return "List the news"
}

func (intent NewsListIntent) Response(input string) string {
	// Doesn't need to do anything with the string yet.
	return news.NewsList()
}

func (intent NewsListIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "news list") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "list the news") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "news") {
		return 50
	}

	return 0
}

func (intent NewsListIntent) PostSpeech(input string) {
	// No op.
}
