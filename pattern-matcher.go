package main

import "os"
import "github.com/codegangsta/cli"
import "github.com/garyburd/redigo/redis"

func main() {
	app := cli.NewApp()
	app.Name = "matcher"
	app.Usage = "Matching goodness"
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}

	app.Run(os.Args)
}
