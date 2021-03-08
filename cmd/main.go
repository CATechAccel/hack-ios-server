package main

import (
	"github.com/ari1021/hack-ios-server/pkg/presentation"
)

func main() {
	e := presentation.NewEcho()
	e.Logger.Fatal(e.Start(":8080"))
}
