package worker

import "github.com/kongqinghan/flyme2thegit/limiter"
import "time"
import "sync"
import "fmt"

type Worker struct {
	totalTasks int
	timeMs int64
	task func(i int)
}

func Run(total int, speed int64, fn func(i int)) {
	w := &Worker{total, speed, fn}
	w.start()
}

func (w *Worker) start() {
	l := limiter.NewLimiter(5, w.timeMs*2)
	wg := sync.WaitGroup{}
	wg.Add(w.totalTasks)
	for i := 0; i < w.totalTasks; i++ {
		go func(id int) {
			defer wg.Done()
			if l.Allow() {
				w.task(id)
			}
		}(i)
		time.Sleep(time.Duration(w.timeMs*1e6))
	}
	wg.Wait()
	l.Stop()
	fmt.Println("all tasks finished")
}

