package renderer

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/devinmcgloin/sail/pkg/library"
	"gopkg.in/cheggaaa/pb.v1"
)

// Render renders the given sketch to the filessystem with the provided seed
func Render(sketchID string, backup bool, seed int64) (*bytes.Buffer, error) {
	runner, err := library.Lookup(sketchID)
	if err != nil {
		return nil, err
	}
	bytes, err := runner.Render(seed)
	if err != nil {
		log.Fatal(err)
	}
	ensureDirectory(runner.Directory())
	runner.Write(seed)

	return bytes, nil
}

// RenderBulk renders each sketch more efficiently
func RenderBulk(sketchID string, backup bool, count, threads int) error {
	seeds := make(chan int64)
	done := make(chan bool)

	bar := pb.StartNew(count)

	runner, err := library.Lookup(sketchID)
	if err != nil {
		return err
	}

	for i := 0; i < threads; i++ {
		go process(bar, runner, done, seeds)
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

func process(bar *pb.ProgressBar, runner library.Runner, done chan bool, seeds chan int64) {
	for seed := range seeds {
		_, err := runner.Render(seed)
		if err != nil {
			continue
		}
		ensureDirectory(runner.Directory())
		runner.Write(seed)
		bar.Increment()
	}
	done <- true
}

func ensureDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
