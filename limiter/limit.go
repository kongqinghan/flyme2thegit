package limiter

import "time"
import "fmt"

type Limiter struct {
	ch chan struct{}
	stop chan struct{}
	timeMs int64
}

func NewLimiter(maxQps int, speedMs int64) *Limiter {
	c := make(chan struct{}, maxQps)
	s := make(chan struct{})
	l := &Limiter{c, s, speedMs}
	go l.start()
	return l
}

func (l *Limiter) start() {
	ticker := time.NewTicker(time.Duration(l.timeMs*1e6))
	defer ticker.Stop()
	for {
		select {
		case <- l.stop:
			fmt.Println("done")
			return
		case <-ticker.C:
			l.ch <- struct{}{}
			fmt.Println("gen token")
		}
	}
}

func (l *Limiter) Stop() {
	l.stop <- struct{}{}
}

func (l *Limiter) Allow() bool {
	select {
	case <- l.ch:
		fmt.Println("allow")
		return true
	default:
		fmt.Println("forbid")
		return false
	}
}
