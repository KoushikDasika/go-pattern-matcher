package main

import (
  "testing"
)

var facts = []map[string]string{
  {"id": "1", "type": "A", "sku": "E031", "amount": "100"},
  {"id": "2", "type": "B", "sku": "E032", "amount": "200"},
  {"id": "3", "type": "A", "sku": "E033", "amount": "300"},
  {"id": "4", "type": "B", "sku": "E034", "amount": "400"},
  {"id": "5", "type": "A", "sku": "E035", "amount": "500"},
  {"id": "6", "type": "B", "sku": "E035", "amount": "600"},
}

type matching_test_pair struct {
  matches map[string]map[string]string
  output []map[string]string
}

var matching_tests = []matching_test_pair{
  { map[string]map[string]string {"type": {"operator": "==", "value": "A"}} ,  []map[string]string {
    {"id": "1", "type": "A", "sku": "E031", "amount": "100"},
    {"id": "3", "type": "A", "sku": "E033", "amount": "300"},
    {"id": "5", "type": "A", "sku": "E035", "amount": "500"},
  }},
  { map[string]map[string]string {"type": {"operator": "!=", "value": "A"}} ,  []map[string]string {
    {"id": "2", "type": "B", "sku": "E032", "amount": "200"},
    {"id": "4", "type": "B", "sku": "E034", "amount": "400"},
    {"id": "6", "type": "B", "sku": "E035", "amount": "600"},
  }},
  { map[string]map[string]string {"id": {"operator": "gt", "value": "3"}} ,  []map[string]string {
    {"id": "4", "type": "B", "sku": "E034", "amount": "400"},
    {"id": "5", "type": "A", "sku": "E035", "amount": "500"},
    {"id": "6", "type": "B", "sku": "E035", "amount": "600"},
  }},
  { map[string]map[string]string {"id": {"operator": "gte", "value": "3"}} ,  []map[string]string {
    {"id": "3", "type": "A", "sku": "E033", "amount": "300"},
    {"id": "4", "type": "B", "sku": "E034", "amount": "400"},
    {"id": "5", "type": "A", "sku": "E035", "amount": "500"},
    {"id": "6", "type": "B", "sku": "E035", "amount": "600"},
  }},
  { map[string]map[string]string {"id": {"operator": "lt", "value": "3"}} ,  []map[string]string {
    {"id": "1", "type": "A", "sku": "E031", "amount": "100"},
    {"id": "2", "type": "B", "sku": "E032", "amount": "200"},
  }},
  { map[string]map[string]string {"id": {"operator": "lte", "value": "3"}} ,  []map[string]string {
    {"id": "1", "type": "A", "sku": "E031", "amount": "100"},
    {"id": "2", "type": "B", "sku": "E032", "amount": "200"},
    {"id": "3", "type": "A", "sku": "E033", "amount": "300"},
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
