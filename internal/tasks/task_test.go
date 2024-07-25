package tasks

import (
	"testing"
)

func BenchmarkAddTask(b *testing.B) {

	for i := 0; i < b.N; i++ {
		//cache.AdvDays()
	}
}
