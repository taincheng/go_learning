package main

import "homework_4_blog/internal/router"

func main() {
	r := router.SetupRouter()
	r.Run()
}
