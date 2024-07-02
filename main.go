package main

import  "github.com/kongqinghan/flyme2thegit/worker"
func main() {
	worker.Run(10, int64(200))
}
