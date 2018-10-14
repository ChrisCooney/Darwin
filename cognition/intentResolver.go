package cognition

import (
	"Darwin/intent"
	"errors"
)

type IntentProbability struct {
	Intent      intent.Intent
	Probability int
}

func ResolveIntent(input string) (intent.Intent, error) {
	intentions := intent.IntentsList
	intention := getClosestMatch(intentions, input)

	if intention == nil {
		return nil, errors.New("I do not understand your command")
	}

	return intention, nil
}

func getClosestMatch(intentions []intent.Intent, input string) intent.Intent {
	var intentionChannel = make(chan IntentProbability)

	for _, intention := range intentions {
		go getProbability(input, intention, intentionChannel)
	}

	var probabilities = []IntentProbability{}

	for x := 0; x < len(intentions); x++ {
		probability := <-intentionChannel
		probabilities = append(probabilities, probability)
	}

	return getIntentWithHighestProbability(probabilities)
}

func getIntentWithHighestProbability(intentionProbabilities []IntentProbability) intent.Intent {
	highestProbability := 0
	var matchedIntention intent.Intent

	for _, probability := range intentionProbabilities {
		value := probability.Probability
		if value > highestProbability {
			highestProbability = value
			matchedIntention = probability.Intent
		}
	}

	return matchedIntention
}

func getProbability(input string, intention intent.Intent, responseChannel chan IntentProbability) {
	probability := intention.Probability(input)
	responseChannel <- IntentProbability{Intent: intention, Probability: probability}
}
