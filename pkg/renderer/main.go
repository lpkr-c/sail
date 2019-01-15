package renderer

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"

	"github.com/devinmcgloin/sail/pkg/library"
	"github.com/devinmcgloin/sail/pkg/sketch"
	"github.com/fogleman/gg"
)

func dir(sketchID string) string {
	return fmt.Sprintf("./sketches/%s/", sketchID)
}

func path(sketchID string, seed int64) string {
	return fmt.Sprintf("./sketches/%s/sketch-%d.png", sketchID, seed)
}

// Render renders the given sketch to the filessystem with the provided seed
func Render(sketchID string, seed int64) error {
	renderable, err := library.Lookup(sketchID)
	if err != nil {
		return err
	}

	context := gg.NewContext(renderable.Dimensions())
	run(renderable, context, dir(sketchID), fmt.Sprintf("sketch-%d.png", seed), seed)
	return nil
}

func RenderBulk(sketchID string, count int)                  {}
func process(renderable sketch.Renderable, seeds <-chan int) {}

func Backup(path string) {

}

func clearBackground(context *gg.Context) {
	context.SetColor(color.White)
	context.DrawRectangle(0, 0, float64(context.Width()), float64(context.Height()))
	context.Fill()
}

func run(renderer sketch.Renderable, context *gg.Context, dir, filename string, seed int64) error {
	fmt.Printf("Rendering: %T\n", renderer)
	fmt.Printf("\tSeed: %d\n", seed)
	rand := rand.New(rand.NewSource(seed))
	clearBackground(context)

	renderer.Draw(context, rand)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	path := dir + filename
	fmt.Printf("\tSaving to: %s\n", path)
	return context.SavePNG(path)
}
