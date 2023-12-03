package main

import (
	"errors"
	"fmt"
	"github.com/625tk/advent-coding/cacheflight"
	"time"
)

func main() {
	v := cacheflight.New()
	t := time.Now().Unix()

	for i := 0; i < 10; i++ {
		g, err := v.Get("key", func() (v any, err error) {
			if time.Now().Unix()-t > 2 {
				return nil, errors.New("some error")
			}
			return "val", nil
		}, 3*time.Second)
		if errors.Is(err, cacheflight.ErrExpired) {
			fmt.Println("expired val", g)
		}

		fmt.Println(g, err)
		time.Sleep(1 * time.Second)
	}
}
