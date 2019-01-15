package renderer

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/devinmcgloin/sail/pkg/library"
	"github.com/devinmcgloin/sail/pkg/sketch"
	"github.com/fogleman/gg"
	pb "gopkg.in/cheggaaa/pb.v1"
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

// RenderBulk renders each sketch more efficiently
func RenderBulk(sketchID string, count, threads int) error {
	seeds := make(chan int64)
	done := make(chan bool)

	bar := pb.StartNew(count)

	renderable, err := library.Lookup(sketchID)
	if err != nil {
		return err
	}

	for i := 0; i < threads; i++ {
		go process(bar, sketchID, renderable, done, seeds)
	}

	start := time.Now().Unix()
	for elapsed := start; elapsed-start < int64(count); elapsed++ {
		seeds <- elapsed
	}

	close(seeds)
	for i := 0; i < threads; i++ {
		<-done
	}
	bar.FinishPrint("Rendering Completed")

	return nil
}
func process(bar *pb.ProgressBar, sketchID string, renderable sketch.Renderable, done chan bool, seeds chan int64) {
	context := gg.NewContext(renderable.Dimensions())
	for seed := range seeds {
		run(renderable, context, dir(sketchID), fmt.Sprintf("sketch-%d.png", seed), seed)
		bar.Increment()
	}
	done <- true
}

// Backup sends the sketch up to the cloud for storage
func Backup(path string) {

}

func clearBackground(context *gg.Context) {
	context.SetColor(color.White)
	context.DrawRectangle(0, 0, float64(context.Width()), float64(context.Height()))
	context.Fill()
}

func run(renderer sketch.Renderable, context *gg.Context, dir, filename string, seed int64) error {
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
	return context.SavePNG(path)
}
