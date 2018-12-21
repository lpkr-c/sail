package sketch

import (
	"errors"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/devinmcgloin/sail/pkg/sketch/accrew"
	"github.com/devinmcgloin/sail/pkg/sketch/primitives"
	"github.com/devinmcgloin/sail/pkg/sketch/sampling"
	"github.com/fogleman/gg"
)

type Config struct {
	SketchID   string
	Seed       int64
	Iterations int
	PathPrefix string
}

type Renderer interface {
	Draw(context *gg.Context, rand *rand.Rand)
	Dimensions() (int, int)
	// AspectRatio() float64
}

func lookup(id string) (Renderer, error) {
	switch id {
	case "accrew/clouds":
		return accrew.Cloud{}, nil
	case "sampling/rectangle":
		return sampling.RectangleDot{}, nil
	case "primitive/rotated-lines":
		return primitives.RotatedLines{}, nil
	default:
		return nil, errors.New("SketchID not found")
	}

}

func clearBackground(context *gg.Context) {
	context.SetColor(color.White)
	context.DrawRectangle(0, 0, float64(context.Width()), float64(context.Height()))
	context.Fill()
}

func Run(config Config) error {
	renderer, err := lookup(config.SketchID)
	if err != nil {
		return err
	}
	context := gg.NewContext(renderer.Dimensions())

	if config.Seed != 0 {
		if err := RunWithSeed(renderer, context, config); err != nil {
			return err
		}
	} else {
		for x := 0; x < config.Iterations; x++ {
			config.Seed = time.Now().Unix()
			if err := RunWithSeed(renderer, context, config); err != nil {
				return err
			}
			time.Sleep(time.Second)
		}
	}
	return nil
}

func RunWithSeed(renderer Renderer, context *gg.Context, config Config) error {
	rand := rand.New(rand.NewSource(config.Seed))
	clearBackground(context)

	renderer.Draw(context, rand)

	dir := fmt.Sprintf("./sketches/%s", config.SketchID)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
	path := fmt.Sprintf("./sketches/%s/%d-sketch.png", config.SketchID, config.Seed)
	fmt.Printf("Saving to: %s\n", path)
	return context.SavePNG(path)
}
