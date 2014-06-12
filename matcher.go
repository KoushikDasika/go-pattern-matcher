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
          for _, val := range v["value"] {
            if (fact[k] != val) {
              flag = false
              break
            }
          }
        case "!=":
          for _, val := range v["value"] {
            if (fact[k] == val) {
              flag = false
              break
            }
          }
        case "gt":
          for _, val := range v["value"] {
            if (fact[k] <= val) {
              flag = false
              break
            }
          }
        case "gte":
          for _, val := range v["value"] {
            if (fact[k] < val) {
              flag = false
              break
            }
          }
        case "lt":
          for _, val := range v["value"] {
            if (fact[k] >= val) {
              flag = false
              break
            }
          }
        case "lte":
          for _, val := range v["value"] {
            if (fact[k] > val) {
              flag = false
              break
            }
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
