package news

import (
	"Darwin/network"
	"Darwin/secret"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Turn into environment variable
var apiUrl = "https://content.guardianapis.com/search?api-key=" + secret.GuardianApiKey()

func Subscribe(query string) string {
	apiUrl := buildSearchApiUrl(query)
	responseBody := network.GetBody(apiUrl)
	apiResponse := parseIntoModel(responseBody)
	liveBlog := findLiveBlog(apiResponse)

	if liveBlog.WebUrl != "" {
		beginMonitoring(liveBlog.WebUrl, liveBlog.WebTitle)
		return "Subscribed to news story."
	} else {
		return "No Live blog found for news story. Unable to subscribe."
	}
}

func NewsList() string {
	responseBody := network.GetBody(apiUrl)
	apiResponse := parseIntoModel(responseBody)
	newsTitles := grabTitles(apiResponse)
	return strings.Join(newsTitles, "\n")
}

func buildSearchApiUrl(query string) string {
	return fmt.Sprintf("%sq=%s", apiUrl, url.QueryEscape(query))
}

func parseIntoModel(responseBody string) *ApiResponse {
	apiResponse := new(ApiResponse)
	json.Unmarshal([]byte(responseBody), apiResponse)
	return apiResponse
}

func findLiveBlog(apiResponse *ApiResponse) Result {
	var liveBlog Result

	for _, result := range apiResponse.Response.Results {
		if result.Type == "liveblog" {
			liveBlog = result
			break
		}
	}

	return liveBlog
}

func grabTitles(apiResponse *ApiResponse) []string {
	var titles []string

	for _, result := range apiResponse.Response.Results {
		titles = append(titles, result.WebTitle)
	}

	return titles
}

func beginMonitoring(webUrl string, title string) {
	go StartNewWatcher(webUrl, title)
}
