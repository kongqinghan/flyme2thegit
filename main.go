package main

import  "github.com/kongqinghan/flyme2thegit/worker"
import "fmt"

func main() {
	fn := func(i int) {
		fmt.Println(i, "work finished")
	}
	worker.Run(10, int64(200), fn)
}
