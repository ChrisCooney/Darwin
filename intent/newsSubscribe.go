package intent

import (
	"Darwin/news"
	"strings"
)

type NewsSubscribeIntent struct{}

func (intent NewsSubscribeIntent) Name() string {
	return "News Subscribe Intention"
}

func (intent NewsSubscribeIntent) Usage() string {
	return "Subscribe to <name of article>"
}

func (intent NewsSubscribeIntent) Probability(input string) int {
	if strings.Contains(strings.ToLower(input), "subscribe me to") {
		return 100
	}

	if strings.Contains(strings.ToLower(input), "subscribe to") {
		return 50
	}

	if strings.Contains(strings.ToLower(input), " subscribe ") {
		return 25
	}

	return 0
}

func (intent NewsSubscribeIntent) Response(input string) string {
	preparedQuery := prepareNewsSubscribeQuery(input)
	message := news.Subscribe(preparedQuery)
	return message
}

func (intent NewsSubscribeIntent) PostSpeech(input string) {
	// No op.
}

func prepareNewsSubscribeQuery(input string) string {
	input = strings.ToLower(input)
	input = strings.Replace(input, "darwin", "", 1)
	input = strings.Replace(input, "subscribe me to", "", 1)
	input = strings.Replace(input, ".", "", 1)
	input = strings.Trim(input, " ")
	input = strings.Trim(input, "\n")
	return input
}
