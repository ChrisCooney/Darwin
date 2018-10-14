package consciousness

import (
  "fmt"
  "bufio"
  "os"
)

func GetTextInput() string {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter command > ")
  text, _ := reader.ReadString('\n')
  return text
}
