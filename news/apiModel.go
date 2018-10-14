package news

type ApiResponse struct {
  Response Response `json:"response"`
}

type Response struct {
  Status string `json:"status"`
  Results []Result `json:"results"`
}

type Result struct {
  ApiUrl string `json:"apiUrl"`
  WebUrl string `json:"webUrl"`
  WebTitle string `json:"webTitle"`
  Type string `json:"type"`
}
