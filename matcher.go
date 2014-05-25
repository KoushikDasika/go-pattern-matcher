package main

import (
  //"fmt"
)

func MatchAttributes(facts []map[string]interface{}, attributes map[string]interface{}) []map[string]interface{} {
  var output []map[string]interface{}
  for _, fact := range facts {
    flag := true
    for k, value := range attributes {
      if (fact[k] != value) {
        flag = false
        break
      }
    }
    if (flag) {
      output = append(output, fact)
    }
  }
  return output
}
