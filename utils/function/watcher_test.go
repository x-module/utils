package function

import (
	"testing"

	"github.com/x-module/utils/utils/internal"
)

func TestWatcher(t *testing.T) {
	assert := internal.NewAssert(t, "TestWatcher")

	w := NewWatcher()
	w.Start()

	longRunningTask()

	assert.Equal(true, w.excuting)

	w.Stop()

	eapsedTime := w.GetElapsedTime().Milliseconds()
	t.Log("Elapsed Time (milsecond)", eapsedTime)

	assert.Equal(false, w.excuting)

	w.Reset()

	assert.Equal(int64(0), w.startTime)
	assert.Equal(int64(0), w.stopTime)
}

func longRunningTask() []int64 {
	var data []int64
	for i := 0; i < 10000000; i++ {
		data = append(data, int64(i))
	}
	return data
}
