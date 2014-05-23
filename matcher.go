package main

import (
  "fmt"
  //"encoding/json"
)

func MatchAttributes(facts []map[string]interface{}, attributes map[string]interface{}) {
  fmt.Println(facts)
  fmt.Println(attributes)
}
