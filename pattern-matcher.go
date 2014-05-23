package main

import (
  "github.com/garyburd/redigo/redis"
  "fmt"
  "github.com/hoisie/web"
  "encoding/json"
  "strconv"
)

var conn redis.Conn
var err error

func getFactsFromRedis(order_id int64) []map[string]interface{} {
  facts_json, err := redis.String(conn.Do("GET", order_id))
  if err != nil {
        // handle error
  }

  var facts []map[string]interface{}
  err = json.Unmarshal([]byte(facts_json), &facts)

  return facts
}

func handleFilter(ctx *web.Context, val string) {
  order_id, err := strconv.ParseInt(ctx.Params["id"], 10, 64)

  if err != nil {
     // handle error
  }

  facts := getFactsFromRedis(order_id)

  response, err := json.Marshal(facts)
  if err != nil {
     // handle error
  }

  ctx.WriteString(string(response[:]))
}

func main() {
  conn, err = redis.Dial("tcp", ":6379")
  if err != nil {
        // handle error
  }
  defer conn.Close()

  web.Post("(.*)", handleFilter)
  web.Run("0.0.0.0:9999")
}
