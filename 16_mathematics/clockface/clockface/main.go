package main

import (
	"os"
	"time"

	"github.com/riyoth/learn-go-with-tests-exercises/16_mathematics/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
