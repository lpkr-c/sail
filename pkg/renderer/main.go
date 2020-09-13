package renderer

import (
	"os"
	"time"

	"gopkg.in/cheggaaa/pb.v1"
)

// RenderBulk renders each sketch more efficiently
func RenderBulk(runner Runner, backup bool, count, threads int) error {
	seeds := make(chan int64)
	done := make(chan bool)

	bar := pb.StartNew(count)

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

func process(bar *pb.ProgressBar, runner Runner, done chan bool, seeds chan int64) {
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
