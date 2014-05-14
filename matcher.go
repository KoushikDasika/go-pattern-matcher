package main

import (
  "github.com/garyburd/redigo/redis"
  "fmt"
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

  conn.Do("APPEND", "1",
  `[{"id":1, "type":"CheckoutItemSelected", "sku":"E031", "amount":100},
    {"id":2, "type":"CheckoutItemSelected", "sku":"E031", "amount":100}]`)

  facts := getFactsFromRedis(conn, "1")

  fmt.Println("facts are %s\n", facts)
  defer conn.Close()
}
