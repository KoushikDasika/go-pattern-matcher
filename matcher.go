package main

import (
  "fmt"
)

func MatchAttributes(facts []map[string]interface{} , attributes map[string]map[string]interface{}) []map[string]interface{}  {
  var output []map[string]interface{}
  var factString string
  var valueString string
  for _, fact := range facts {
    flag := true
    for k, v := range attributes {
      factString = fmt.Sprint(fact[k])
      valueString = fmt.Sprint(v["value"])
      switch v["operator"] {
        case "==":
          if (factString != valueString) {
            flag = false
            break
          }
        case "!=":
          if (factString == valueString) {
            flag = false
            break
          }
        case "gt":
          if (factString <= valueString) {
            flag = false
            break
          }
        case "gte":
          if (factString < valueString) {
            flag = false
            break
          }
        case "lt":
          if (factString >= valueString) {
            flag = false
            break
          }
        case "lte":
          if (factString > valueString) {
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
