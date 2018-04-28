package main

import (
	"ForestModel/util"
	"ForestModel/world"
	"log"
)

func main() {
	world, err := world.InitializeWorld(util.Rect{10000, 10000}, 5)
	if err != nil {
		log.Fatal(err)
	}
}
