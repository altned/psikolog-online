package main

import "github.com/example/psikolog-online/internal/routes"

func main() {
	r := routes.SetupRouter()
	r.Run()
}
