package main

import (
  "testing"
  //"fmt"
)

var facts = []map[string]interface{} {
  {"id": 1, "type": "A", "sku": "E031", "amount": 100},
  {"id": 2, "type": "B", "sku": "E032", "amount": 200},
  {"id": 3, "type": "A", "sku": "E033", "amount": 300},
  {"id": 4, "type": "B", "sku": "E034", "amount": 400},
  {"id": 5, "type": "A", "sku": "E035", "amount": 500},
  {"id": 6, "type": "B", "sku": "E035", "amount": 600},
}

type matching_test_pair struct {
  matches map[string]map[string]interface{}
  output []map[string]interface{}
}

var matching_tests = []matching_test_pair{
  { map[string]map[string]interface{} {"type": {"operator": "==", "value": "A"}} ,  []map[string]interface{}  {
    {"id": 1, "type": "A", "sku": "E031", "amount": 100},
    {"id": 3, "type": "A", "sku": "E033", "amount": 300},
    {"id": 5, "type": "A", "sku": "E035", "amount": 500},
  }},
  { map[string]map[string]interface{} {"type": {"operator": "!=", "value": "A"}} ,  []map[string]interface{}  {
    {"id": 2, "type": "B", "sku": "E032", "amount": 200},
    {"id": 4, "type": "B", "sku": "E034", "amount": 400},
    {"id": 6, "type": "B", "sku": "E035", "amount": 600},
  }},
  { map[string]map[string]interface{} {"id": {"operator": "gt", "value": "3"}} ,  []map[string]interface{}  {
    {"id": 4, "type": "B", "sku": "E034", "amount": 400},
    {"id": 5, "type": "A", "sku": "E035", "amount": 500},
    {"id": 6, "type": "B", "sku": "E035", "amount": 600},
  }},
  { map[string]map[string]interface{} {"id": {"operator": "gte", "value": "3"}} ,  []map[string]interface{}  {
    {"id": 3, "type": "A", "sku": "E033", "amount": 300},
    {"id": 4, "type": "B", "sku": "E034", "amount": 400},
    {"id": 5, "type": "A", "sku": "E035", "amount": 500},
    {"id": 6, "type": "B", "sku": "E035", "amount": 600},
  }},
  { map[string]map[string]interface{} {"id": {"operator": "lt", "value": "3"}} ,  []map[string]interface{}  {
    {"id": 1, "type": "A", "sku": "E031", "amount": 100},
    {"id": 2, "type": "B", "sku": "E032", "amount": 200},
  }},
  { map[string]map[string]interface{} {"id": {"operator": "lte", "value": "3"}} ,  []map[string]interface{}  {
    {"id": 1, "type": "A", "sku": "E031", "amount": 100},
    {"id": 2, "type": "B", "sku": "E032", "amount": 200},
    {"id": 3, "type": "A", "sku": "E033", "amount": 300},
  }},
}

func TestMatchAttributes(t *testing.T) {
  for _, pair := range matching_tests {
    v := MatchAttributes(facts, pair.matches)
    for i, _ := range pair.output {
      for j, _ := range pair.output[i] {
        if v[i][j] != pair.output[i][j] {
          t.Error(
            "For", pair.matches,
            "expected", pair.output[i][j],
            "got", v[i][j],
          )
        }
      }
    }
  }
}
