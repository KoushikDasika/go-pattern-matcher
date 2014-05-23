package main

import (
  "testing"
)

var facts = []map[string]interface{}{
  {"id": 1, "type": "A", "sku": "E031", "amount": 100},
  {"id": 2, "type": "B", "sku": "E032", "amount": 200},
  {"id": 3, "type": "C", "sku": "E033", "amount": 300},
  {"id": 4, "type": "D", "sku": "E034", "amount": 400},
  {"id": 5, "type": "E", "sku": "E035", "amount": 500},
}

type testpair struct {
  matches map[string]interface{}
  output []map[string]interface{}
}


func TestMatchAttributes(t *testing.T) {
  MatchAttributes(facts)
}
