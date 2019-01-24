package renderer

import (
	"bytes"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/devinmcgloin/sail/pkg/cloud"
	"github.com/devinmcgloin/sail/pkg/library"
	"github.com/devinmcgloin/sail/pkg/sketch"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
	pb "gopkg.in/cheggaaa/pb.v1"
)

func dir(sketchID string) string {
	return fmt.Sprintf("sketches/%s/", sketchID)
}

func path(sketchID string, seed int64) string {
	return fmt.Sprintf("sketches/%s/sketch-%d.png", sketchID, seed)
}

// Render renders the given sketch to the filessystem with the provided seed
func Render(sketchID string, backup bool, seed int64) (*bytes.Buffer, error) {
	renderable, err := library.Lookup(sketchID)
	if err != nil {
		return nil, err
	}

	xd, yd := renderable.Dimensions()
	context := gg.NewContext(xd, yd)
	slog.InfoPrintf("Rendering %T with dimesions (%d, %d) and seed %d\n", renderable, xd, yd, seed)
	bytes, err := run(renderable, context, dir(sketchID), fmt.Sprintf("sketch-%d.png", seed), seed)
	if err == nil && backup {
		cloud.Upload(bytes, path(sketchID, seed))
	}
	return bytes, nil
}

// RenderBulk renders each sketch more efficiently
func RenderBulk(sketchID string, backup bool, count, threads int) error {
	seeds := make(chan int64)
	done := make(chan bool)

	bar := pb.StartNew(count)

	renderable, err := library.Lookup(sketchID)
	if err != nil {
		return err
	}

	for i := 0; i < threads; i++ {
		go process(bar, sketchID, backup, renderable, done, seeds)
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
func process(bar *pb.ProgressBar, sketchID string, backup bool, renderable sketch.Renderable, done chan bool, seeds chan int64) {
	xd, yd := renderable.Dimensions()
	context := gg.NewContext(xd, yd)
	for seed := range seeds {
		slog.DebugPrintf("Rendering %T with dimesions (%d, %d) and seed: %d\n", renderable, xd, yd, seed)
		bytes, err := run(renderable, context, dir(sketchID), fmt.Sprintf("sketch-%d.png", seed), seed)
		if err == nil && backup {
			cloud.Upload(bytes, path(sketchID, seed))
		}
		bar.Increment()
	}
	done <- true
}

func clearBackground(context *gg.Context) {
	context.SetColor(color.White)
	context.DrawRectangle(0, 0, float64(context.Width()), float64(context.Height()))
	context.Fill()
}

func run(renderer sketch.Renderable, context *gg.Context, dir, filename string, seed int64) (*bytes.Buffer, error) {
	rand := rand.New(rand.NewSource(seed))
	clearBackground(context)

	renderer.Draw(context, rand)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
	}
	path := dir + filename

	bytes := new(bytes.Buffer)

	err := context.EncodePNG(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, context.SavePNG(path)
}
