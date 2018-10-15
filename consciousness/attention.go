package consciousness

import (
	"Darwin/cognition"
	"Darwin/intent"
	"Darwin/output"
)

func PayAttention() {
	keepAlive := true

	for keepAlive {
		input := GetTextInput()
		HandleInput(input)
	}
}

func HandleInput(input string) {
	intent, err := cognition.ResolveIntent(input)

	if intent == nil {
		HandleMissingIntent(intent)
	} else {
		Respond(err, input, intent, intent.Response(input))
	}
}

func HandleMissingIntent(intent intent.Intent) {
	output.PrintError("Sorry, I don't understand.")
}

func Respond(err error, input string, intent intent.Intent, response string) {
	if err != nil {
		output.PrintError(err.Error())
	} else {
		if response != "" && response != " " {
			output.PrintOutput(response)
		}
	}

	intent.PostSpeech(input)
}
