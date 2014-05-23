package main

import (
  "testing"
)

var facts = []map[string]interface{}{
  {"id": 1, "type": "A", "sku": "E031", "amount": 100},
  {"id": 2, "type": "B", "sku": "E032", "amount": 200},
  {"id": 3, "type": "A", "sku": "E033", "amount": 300},
  {"id": 4, "type": "B", "sku": "E034", "amount": 400},
  {"id": 5, "type": "A", "sku": "E035", "amount": 500},
  {"id": 6, "type": "B", "sku": "E035", "amount": 500},
}

type testpair struct {
  matches map[string]interface{}
  output []map[string]interface{}
}

var matching_tests = []testpair{
  { map[string]interface{} {"type": "A"} ,  []map[string]interface{} {
    {"id": 1, "type": "A", "sku": "E031", "amount": 100},
    {"id": 3, "type": "A", "sku": "E033", "amount": 300},
    {"id": 5, "type": "A", "sku": "E035", "amount": 500},
  }},
}

func TestMatchAttributes(t *testing.T) {
  for _, pair := range matching_tests {
    v := MatchAttributes(facts, pair.matches)
    if v != pair.output {
      t.Error(
        "For", pair.matches,
        "expected", pair.output,
        "got", v,
      )
    }
  }
}
