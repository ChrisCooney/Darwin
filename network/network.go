package network

import (
	"Darwin/output"
	"io/ioutil"
	"net/http"
)

func GetBody(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		output.PrintError(err.Error())
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}
