package main

import (
	"flag"
	"fmt"

	"github.com/devinmcgloin/generative-go/pkg/sketch"
)

func main() {
	cfg := sketch.Config{}

	flag.StringVar(&cfg.SketchID, "sketch-id", "", "Sketch to render")
	flag.Int64Var(&cfg.Seed, "seed", 0, "Seed to render sketch with")
	flag.IntVar(&cfg.Iterations, "iterations", 1, "Number of times to run the sketch")
	flag.Parse()

	err := sketch.Run(cfg)
	if err != nil {
		fmt.Printf("Encountered Error: %+v\n", err)
	}
}
