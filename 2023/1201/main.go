package main

import (
	"github.com/625tk/advent-coding/dtool"
	"os"
	"time"
)

func main() {
	dtool.PrintUnixtime(os.Stdout, time.Now().Unix())
}
