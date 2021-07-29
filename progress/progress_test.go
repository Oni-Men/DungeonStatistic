package progress_test

import (
	"jp/thelow/static/progress"
	"testing"
	"time"
)

func TestProgressBar(t *testing.T) {

	bar := progress.NewProgressBar("Test Bar")

	for i := 0; i < 10; i++ {
		bar.AdvanceProgress(0.1)
		time.Sleep(100 * time.Microsecond)
	}

	bar.CompleteProgress()
}
