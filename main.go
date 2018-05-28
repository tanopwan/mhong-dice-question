package main

import "fmt"

func main() {
	cc := 0
	i := 0
	for _, a := range []uint8{1, 2, 3, 4, 5, 6} {
		for _, b := range []uint8{1, 2, 3, 4, 5, 6} {
			for _, c := range []uint8{1, 2, 3, 4, 5, 6} {
				for _, d := range []uint8{1, 2, 3, 4, 5, 6} {
					i = i + 1
					if isDouble(i, [4]uint8{a, b, c, d}) {
						cc = cc + 1
					}
				}
			}
		}
	}

	max := 1296.0
	fmt.Printf("โอกาสที่จะเกิด: %d\n", cc)
	fmt.Printf("โอกาสทั้งหมด: 1296\n")
	fmt.Printf("prop: %f\n", float64(cc)/max)
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
