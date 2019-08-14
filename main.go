package main

import "sort"

// Go solution to HackerRank Fraudulent Activity Notifications challenge
// https://www.hackerrank.com/challenges/fraudulent-activity-notifications/
func activityNotifications(exp []int, d int) int {
	dIsEven := d%2 == 0
	alerts := 0
	mp := d / 2
	win := make([]int, d)
	copy(win, exp[:d]) // otherwise it will overwrite exp[]
	sort.Ints(win)
	for cur_ind, cur_val := range exp[d:] {
		win_elem_to_del := exp[cur_ind]

		// increment alerts alerts
		if dIsEven {
			if cur_val >= win[mp]+win[mp-1] { //cur_val >= 2*midpoint
				alerts += 1
			}
		} else if cur_val >= win[mp]*2 { // odd AND cur_val >= 2*midpoint
			alerts += 1
		}

		// delete oldest value
		delIdx := sort.SearchInts(win, win_elem_to_del)
		win = append(win[:delIdx], win[delIdx+1:]...)

		// insert newest value in correct sorted location - with memory leak fix f/ go wiki - slice tricks
		insIdx := sort.SearchInts(win, cur_val)
		win = append(win, 0)
		copy(win[insIdx+1:], win[insIdx:])
		win[insIdx] = cur_val
	}
	return alerts
}
