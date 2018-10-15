package intent

var IntentsList = []Intent{SleepIntent{}, SearchIntent{}, GoToWorkIntent{}, HomeTimeIntent{}, ListIntentionsIntent{}, OpenAppIntent{}, SetMeUpIntent{}}

// Intent interface defines some behaviour Darwin can perform when given the appropriate input.
type Intent interface {

	// Returns the name of this intent, i.e "Hello Intent"
	Name() string

	// Returns a usage example for this intent i.e Hello
	Usage() string

	// Returns the probability that this intent is the correct matchh for the input, as a percentage.
	Probability(input string) int

	// Returns the response to the input. This is where the meat of the Intent should live.
	Response(input string) string

	// A function that can execute after the speech is finished.
	PostSpeech(input string)
}
