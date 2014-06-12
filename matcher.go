package main

import (
  "fmt"
)

func MatchAttributes(facts []map[string]interface{} , attributes map[string]map[string]interface{}) []map[string]interface{}  {
  var output []map[string]interface{}
  for _, fact := range facts {
    flag := true
    for k, v := range attributes {
      switch v["operator"] {
        case "==":
          if (fact[k] != v["value"]) {
            flag = false
            break
          }
        case "!=":
          if (fact[k] == v["value"]) {
            flag = false
            break
          }
        case "gt":
          if (fact[k] <= v["value"]) {
            flag = false
            break
          }
        case "gte":
          if (fact[k] < v["value"]) {
            flag = false
            break
          }
        case "lt":
          if (fact[k] >= v["value"]) {
            flag = false
            break
          }
        case "lte":
          if (fact[k] > v["value"]) {
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
