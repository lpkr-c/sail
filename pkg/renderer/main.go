package renderer

import (
	"bytes"
	"errors"
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/devinmcgloin/sail/pkg/cloud"
	"github.com/devinmcgloin/sail/pkg/library"
	"github.com/devinmcgloin/sail/pkg/sketch"
	"github.com/devinmcgloin/sail/pkg/slog"
	"github.com/fogleman/gg"
	pb "gopkg.in/cheggaaa/pb.v1"
)

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
	slog.InfoPrintf("Rendering %T with dimesions (%d, %d) and seed %d\n", renderable, xd, yd, seed)
	slog.InfoValues("backup", backup)

	context := gg.NewContext(xd, yd)
	bytes, err := run(renderable, context, seed)
	if err != nil {
		slog.ErrorPrintf("error %+v\n", err)
		return nil, err
	}
	if err == nil && backup {
		cloud.Upload(bytes, path(sketchID, seed))
	}
	if bytes == nil {
		slog.ErrorPrintf("error %+v\n", err)
		return nil, errors.New("bytes was nil inside Render method")
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
	for seed := range seeds {
		context := gg.NewContext(xd, yd)
		bytes, err := run(renderable, context, seed)
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

func run(renderer sketch.Renderable, context *gg.Context, seed int64) (*bytes.Buffer, error) {
	rand := rand.New(rand.NewSource(seed))
	clearBackground(context)

	renderer.Draw(context, rand)

	bytes := new(bytes.Buffer)
	err := context.EncodePNG(bytes)
	if err != nil {
		slog.ErrorPrintf("Encoutered Error %s when encoding\n", err)
		return nil, err
	}

	if bytes == nil {
		slog.ErrorPrintf("bytes was nill when returning from run\n")
		return nil, errors.New("bytes was nill when returning from run")
	}

	return bytes, nil
}
