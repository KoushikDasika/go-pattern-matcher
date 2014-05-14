package main

import (
  "github.com/garyburd/redigo/redis"
  "fmt"
  "encoding/json"
)

func getFactsFromRedis(conn redis.Conn, order_id string) string {
  facts, err := redis.String(conn.Do("GET", order_id))
  if err != nil {
        // handle error
  }
  return facts
}

func main() {
  conn, err := redis.Dial("tcp", ":6379")
  if err != nil {
        // handle error
  }
  defer conn.Close()

  conn.Do("SET", 1, "[{\"type\":\"A\",\"id\":5,\"sku\":\"E031\",\"amount\":100},{\"type\":\"B\",\"id\":6,\"sku\":\"E032\",\"amount\":200}]")
  facts_json := getFactsFromRedis(conn, "1")

  fmt.Printf("facts are %s\n", facts_json)
  var facts []map[string]interface{}
  err = json.Unmarshal([]byte(facts_json), &facts)
  if err != nil {
    fmt.Println("error:", err)
  }

  for i := range facts {
    fmt.Println(facts[i])
  }


  /*
  blah := []map[string]interface{} {
    map[string]interface{} {
      "id": 1,
      "type": "CheckoutItemSelected",
    },

    map[string]interface{} {
      "id": 2,
      "type": "CheckoutItemRemoved",
    },


  }
  for i := range blah {
    for k, v := range blah[i] {
      fmt.Printf("for %dth fact, for key %s value %s\n", i, k, v)
    }
  }
  */

}
