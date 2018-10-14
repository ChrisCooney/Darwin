package output

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

func PrintOutputListen(printChannel chan string) {
	for {
		value := <-printChannel
		PrintOutput(value)
	}
}

func PrintErrorListen(printErrorChannel chan string) {
	for {
		value := <-printErrorChannel
		PrintError(value)
	}
}

func PrintOutput(output string) {
	fmt.Printf("Darwin > %s. \n", output)
}

func PrintError(output string) {
	color.Red("Darwin > %s \n", output)
}

type CountSummary struct {
	Count   int
	Total   int
	Message string
}

func PrintCountSummary(title string, countSummaries ...CountSummary) {
	fmt.Printf("Darwin > %s \n", title)

	for _, summary := range countSummaries {
		var colour = calculateGoodOrBadColour(summary.Count, summary.Total)
		fmt.Printf("-> %s/%s %s \n", colour(summary.Count), colour(summary.Total), summary.Message)
	}
}

func calculateGoodOrBadColour(count int, total int) func(...interface{}) string {
	if count < total {
		return red
	}

	return green
}
