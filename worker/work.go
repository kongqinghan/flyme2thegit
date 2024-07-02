package worker

import "github.com/kongqinghan/flyme2thegit/limiter"
import "time"
import "sync"
import "fmt"

type Worker struct {
	totalTasks int
	timeMs int64
}

func Run(total int, speed int64) {
	w := &Worker{total, speed}
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
				fmt.Println(id, "allow")
			} else {
				fmt.Println(id, "forbid")
			}
		}(i)
		time.Sleep(time.Duration(w.timeMs*1e6))
	}
	wg.Wait()
	l.Stop()
	fmt.Println("all tasks finished")
}

