package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	cc := 0
	i := 0
	jobs := make(chan [4]uint8, 1296)
	results := make(chan int, 1296)
	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results)
	}

	for _, a := range []uint8{1, 2, 3, 4, 5, 6} {
		for _, b := range []uint8{1, 2, 3, 4, 5, 6} {
			for _, c := range []uint8{1, 2, 3, 4, 5, 6} {
				for _, d := range []uint8{1, 2, 3, 4, 5, 6} {
					i = i + 1
					jobs <- [4]uint8{a, b, c, d}
				}
			}
		}
	}
	close(jobs)

	for i := 0; i < 1296; i++ {
		r := <-results
		cc = cc + r
	}

	max := 1296.0
	elapsed := time.Since(start)
	log.Printf("เวลาที่ใช้คำนวน %s", elapsed)
	fmt.Printf("โอกาสที่จะเกิด: %d\n", cc)
	fmt.Printf("โอกาสทั้งหมด: 1296\n")
	fmt.Printf("prop: %f\n", float64(cc)/max)
}

func worker(id int, jobs <-chan [4]uint8, results chan<- int) {
	for j := range jobs {
		r := isDouble(id, j)
		if r {
			results <- 1
		} else {
			results <- 0
		}
	}
}

func isDouble(i int, dd [4]uint8) bool {
	dice := map[uint8]uint8{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}

	found := false
	for _, d := range dd {
		dice[uint8(d)] = dice[uint8(d)] + 1
		if dice[uint8(d)] == 1 {
			continue
		} else if dice[uint8(d)] == 2 && found == false {
			found = true
			continue
		}
		found = false
		break
	}
	fmt.Printf("(%d) %v ---- %v, ---> %v\n", i, dd, dice, found)
	return found
}
