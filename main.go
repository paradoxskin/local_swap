package main

import "local_swap/router"

func main() {
	r := router.Build()
	r.Run()
}
