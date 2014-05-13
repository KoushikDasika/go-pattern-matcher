package main

import (
  "github.com/hoisie/web"
)

func handleFilter(ctx *web.Context, val string) string {
  for k,v := range ctx.Params {
    println(k, v)
  }
  return "Koushik is the best"
}

func main() {
  web.Post("(.*)", handleFilter)
  web.Run("0.0.0.0:9999")
}
