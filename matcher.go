package main

import (
  "fmt"
)

func MatchAttributes(facts []map[string]interface{}, attributes map[string]map[string]interface{}) []map[string]interface{} {
  var output []map[string]interface{}
  for _, fact := range facts {
    flag := true
    for k, v := range attributes {
      switch v["operator"] {
        case "==":
          if (fact[k] != v["value"]) {
            fmt.Println(v)
            flag = false
            break
          }
        default:
          fmt.Println("Pick a better operator, foo")
      }
    }
    if (flag) {
      output = append(output, fact)
    }
  }
  return output
}
