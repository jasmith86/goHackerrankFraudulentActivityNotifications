package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

func main() {
	/*
		Start My custom test cases
	*/

	// case custom 0
	expenditure := []int{2, 3, 4, 2, 3, 6, 8, 4, 5, 33, 44, 55, 66, 77, 88, 5, 33, 44, 55, 66, 77, 88, 5, 33, 44, 55, 66, 77, 88, 5, 33, 44, 55, 66, 77, 88}
	d := 5
	result := activityNotifications(expenditure, d)
	if result != 6 {
		log.Panicf("Result %v != 6", result)
	} else {
		fmt.Println("Case Custom 0: Ok")
	}
	// case custom 1
	expenditure = []int{4, 3, 2, 1, 5}
	d = 4
	result = activityNotifications(expenditure, d)
	if result != 1 {
		log.Fatalf("Result %v != 0", result)
	} else {
		fmt.Println("Case custom 1: Ok")
	}

	/*
		Start Hackerrank sample test cases
	*/

	// case 0
	expenditure = []int{2, 3, 4, 2, 3, 6, 8, 4, 5}
	d = 5
	result = activityNotifications(expenditure, d)
	if result != 2 {
		log.Fatalf("Result %v != 2", result)
	} else {
		fmt.Println("Case 0: Ok")
	}

	// case 1
	expenditure = []int{1, 2, 3, 4, 4}
	d = 4
	result = activityNotifications(expenditure, d)
	if result != 0 {
		log.Fatalf("Result %v != 0", result)
	} else {
		fmt.Println("Case 1: Ok")
	}

	// case 2
	expenditure = []int{10, 20, 30, 40, 50}
	d = 3
	result = activityNotifications(expenditure, d)
	if result != 1 {
		log.Fatalf("Result %v != 1", result)
	} else {
		fmt.Println("Case 2: Ok")
	}

	/*
		Start LONG hackerrank test case 01
	*/
	// test case 01 (real)
	f, err := os.Open("/home/james/Documents/projects/misc_python/fraudulent-activity-notifications-testcase01.txt")
	checkError(err)
	reader := bufio.NewReaderSize(f, 1024*1024)
	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d = int(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	expenditure = []int{}

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}
	result = activityNotifications(expenditure, d)
	ans := 633
	fmt.Println(result)
	if result != ans {
		log.Fatalf("Result %d != $d", result, result, ans)
	} else {
		fmt.Println("Case LONG 01: Ok")
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
