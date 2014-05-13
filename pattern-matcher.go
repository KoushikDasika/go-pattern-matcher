package main

import (
  "github.com/hoisie/web"
  "encoding/json"
)

func handleFilter(ctx *web.Context, val string) {
  for k,v := range ctx.Params {
    println(k, v)
  }

  enc := json.NewEncoder(ctx.ResponseWriter)
  d := map[string]int{"Koushik": 9001, "Charlie": 5}
  enc.Encode(d)
}

func main() {
  web.Post("(.*)", handleFilter)
  web.Run("0.0.0.0:9999")
}
