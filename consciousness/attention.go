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

		intent, err := cognition.ResolveIntent(input)

		Respond(err, input, intent, intent.Response(input))
	}
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
