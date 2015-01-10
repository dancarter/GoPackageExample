package greet

import (
  "log"
  "fmt"
  "time"
)

func init() {
  log.Println(time.Now())
}

// exported function that can be used outside the package
func SayHi(name string) {
  logger(fmt.Sprintf("Hello, %s!", name))
}

// not exported and can not be used outside the package
func logger(s string) {
  log.Println(s)
}
