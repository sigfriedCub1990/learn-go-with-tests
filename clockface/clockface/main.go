package main

import (
	"os"
	"time"

	"learninggo.com/m/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
