package news

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "Darwin/output"
  "time"
)

type SiteWatcher struct {
  title string
  webUrl string
  lastMessage string
}

func (watcher *SiteWatcher) PollForLatestMessage() (string, bool) {
  fmt.Printf("Live url = %s\n", watcher.webUrl)
  doc, _ := goquery.NewDocument(watcher.webUrl)
  latestUpdateWrapper := doc.Find(".block.block--content").First()
  latestUpdateTextElement := latestUpdateWrapper.Find(".block-elements.block-elements--no-byline")

  if(latestUpdateWrapper.HasClass("is-summary")) {
    // No more live content.
    return fmt.Sprintf("There are no more updates left on %s", watcher.title), false
  }

  latestMessage := latestUpdateTextElement.Text()

  if latestMessage != watcher.lastMessage {
    output.PrintOutput("New updates.")
    watcher.lastMessage = latestMessage
    return latestMessage, true
  }

  return "", true
}

func StartNewWatcher(webUrl string, title string) {
  watcher := SiteWatcher{webUrl: webUrl, title: title}

  for true {
    message, continuePolling := watcher.PollForLatestMessage()

    if message != "" {
      completeMessage := message

      if continuePolling == true {
        completeMessage += ". Listening for more updates."
      }

      output.PrintOutput(completeMessage)
    }

    if(!continuePolling) {
      break
    }

    time.Sleep(60 * time.Second)
  }
}
