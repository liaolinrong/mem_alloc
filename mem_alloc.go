package main

import (
	"flag"
	"fmt"
	"time"
)

const MAX_M = 512 * 1024
const ONE_M = 1024 * 1024

func main() {
	var num_m = flag.Int("n", 500, "alloc how many M")
	var time_Millisecond = flag.Int("t", 1000, "sleep how many milliseconds after alloc every 100M")
	var time_delay = flag.Int("d", 3600, "sleep how many seconds after alloc all memory")
	flag.Parse()

	var p [MAX_M]*[ONE_M]byte

	i := 0
	for i = 0; i < *num_m; i++ {
		p[i] = new([ONE_M]byte)
		for j := 0; j < ONE_M; j++ {
			p[i][j] = byte(j % 128)
		}

		fmt.Printf("alloc %dM, *num_m: %d\n", i, *num_m)
		if i > 0 && i%100 == 0 {
			fmt.Printf("alloc %dM\n", i)
			time.Sleep(time.Duration(*time_Millisecond) * time.Millisecond)
		}
	}

	fmt.Printf("alloc %dM\n", i)

	//	var b byte
	//	fmt.Scanf("%c", &b)
	time.Sleep(time.Duration(*time_delay) * time.Second)
}
