package main

import (
	"dushanbe_tour/internal"
	"dushanbe_tour/internal/router"
	"go.uber.org/fx"
	"runtime"
)

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	app := fx.New(
		internal.Modules,
		router.EntryPoint,
	)
	app.Run()
}
