package main

import (
	"github.com/csolarz-ml/chain-of-responsibility-api-pattern/routes"
)

func main() {
	app := routes.Setup()
	app.Run()
}
