package query

type QueryResult struct {
  Success string `xml:"success,attr"`
  Pods []Pod `xml:"pod"`
}

type Pod struct {
  Title string `xml:"title,attr"`
  SubPod SubPod `xml:"subpod"`
}

type SubPod struct {
  Plaintext Plaintext `xml:"plaintext"`
}

type Plaintext struct {
  Message string `xml:",chardata"`
}
