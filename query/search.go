package query

import (
	"Darwin/network"
	"Darwin/secret"
	"encoding/xml"
	"fmt"
	"net/url"
	"regexp"
)

// Returns a plaintext response from the Wolfram Alpha API based on the input query.
func Search(query string) string {
	apiUrl := buildApiUrl(query)
	body := network.GetBody(apiUrl)
	return getPlaintextAnswer(body)
}

func buildApiUrl(input string) string {
	apiTemplate := "http://api.wolframalpha.com/v2/query?input=%s&appid=" + secret.WolframAlphaAppId() + "&format=plaintext"
	escapedInput := url.QueryEscape(input)
	return fmt.Sprintf(apiTemplate, escapedInput)
}

func getPlaintextAnswer(xmlBody string) string {
	queryResult := new(QueryResult)
	xml.Unmarshal([]byte(xmlBody), queryResult)

	if queryResult.Success == "false" {
		return ""
	}

	resultPod := getResultPod(queryResult)

	response := getShortPlaintextAnswer(resultPod.SubPod.Plaintext.Message)
	return response
}

func getResultPod(queryResult *QueryResult) Pod {
	var resultPod Pod

	for _, pod := range queryResult.Pods {
		if pod.Title == "Result" {
			resultPod = pod
			break
		}
	}

	return resultPod
}

// Wolfram alpha returns a single answer, then in brackets a bit description. We only need the short one.
func getShortPlaintextAnswer(fullAnswer string) string {
	re := regexp.MustCompile(`(?s)\((.*)\)`)
	return string(re.ReplaceAll([]byte(fullAnswer), []byte("")))
}
