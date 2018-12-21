package sketch

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type Config struct {
	SketchID   string
	Seed       *uint64
	Iterations *int
	Height     int
	Width      int
	PathPrefix string
}
type Renderer interface {
	Draw(context *gg.Context, rand rand.Source)
}

func lookup(id string) Renderer {
	switch id {
	case "accrew:clouds":
		return accrew.Clouds
	case "sample:rectangle":
		return accrew.Clouds
	}

}

func clearBackground(context *gg.Context, width, height int) {
	context.SetColor(color.White)
	context.DrawRectangle(0, 0, config.Height, config.Width)
	context.Fill()
}

func Run(config Config) {
	renderer := lookup(config.SketchID)

	if config.Seed != nil {
		RunWithSeed(renderer, config)
	}
	return context.SavePNG(fmt.Sprintf("./sketches/%s/%d-sketch.png", sketchID, seed))
}

func RunWithSeed(renderer Renderer, config Config) {
	rand := rand.New(rand.NewSource(seed))
	context := gg.NewContext(config.Height, config.Width)
	clearBackground(context, config.Height, config.Width)

	renderer.Draw(context, rand)
}
