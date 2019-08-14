package main

import (
	"sort"
)

func activityNotifications(exp []int, d int) int {
	dIsEven := d%2 == 0
	// fmt.Printf("even? %v\n", dIsEven)
	// fmt.Printf("%v\n", exp)
	alerts := 0
	mp := d / 2
	win := make([]int, d)
	copy(win, exp[:d])
	sort.Ints(win)
	// fmt.Printf("%v\n", win)
	for cur_ind, cur_val := range exp[d:] {
		win_elem_to_del := exp[cur_ind]
		// fmt.Printf("will delete win[%v] = %v\n", cur_ind, win_elem_to_del)

		// incr alerts
		if dIsEven {
			if cur_val >= win[mp]+win[mp-1] { //cur_val >= 2*midpoint
				alerts += 1
				// fmt.Println("alert + 1")
			}
		} else if cur_val >= win[mp]*2 { // odd AND cur_val >= 2*midpoint
			alerts += 1
			// fmt.Println("alert + 1")
		}
		//if win_elem_to_add == 77  && d == 5{
		//	continue
		//}
		// delete  - with memory leak fix f/ go wiki - slice tricks
		delIdx := sort.SearchInts(win, win_elem_to_del)

		//if delIdx == len(win) {
		//	delIdx--
		//}

		//	fmt.Printf("%v", len(win))
		//	fmt.Printf("%v", insIdx)
		//	copy(win[insIdx:], win[insIdx+1:])
		//	win[len(win)-1] = 0 // or the zero value of T
		//	win = win[:len(win)-1]
		win = append(win[:delIdx], win[delIdx+1:]...)

		// insert - with memory leak fix f/ go wiki - slice tricks
		insIdx := sort.SearchInts(win, cur_val)
		win = append(win, 0 /* use the zero value of the element type */)
		copy(win[insIdx+1:], win[insIdx:])
		win[insIdx] = cur_val
		//fmt.Printf("%v\n", win)

		// fmt.Printf("%v\n", win)
	}
	return alerts

}
