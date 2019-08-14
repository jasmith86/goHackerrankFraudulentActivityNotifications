package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

type testdata struct {
	name        string
	expenditure []int
	d           int
	answer      int
}

func TestActivityNotifications(t *testing.T) {
	tests := []testdata{
		{
			name:        "custom case 0",
			expenditure: []int{2, 3, 4, 2, 3, 6, 8, 4, 5, 33, 44, 55, 66, 77, 88, 5, 33, 44, 55, 66, 77, 88, 5, 33, 44, 55, 66, 77, 88, 5, 33, 44, 55, 66, 77, 88},
			d:           5,
			answer:      6,
		},
		{
			name:        "custom case 1",
			expenditure: []int{4, 3, 2, 1, 5},
			d:           4,
			answer:      1,
		},
		{
			name:        "hr sample case 0",
			expenditure: []int{2, 3, 4, 2, 3, 6, 8, 4, 5},
			d:           5,
			answer:      2,
		},
		{
			name:        "hr sample case 1",
			expenditure: []int{1, 2, 3, 4, 4},
			d:           4,
			answer:      0,
		},
		{
			name:        "hr sample case 2",
			expenditure: []int{10, 20, 30, 40, 50},
			d:           3,
			answer:      1,
		},
	}
	tmp := readBigTestCase()
	tests = append(tests, tmp)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := activityNotifications(test.expenditure, test.d)
			if test.answer == result {
				t.Errorf("wanted %d, but got %d", test.answer, result)
			}
		})
	}
}
func readBigTestCase() testdata {
	f, err := os.Open("./fraudulent-activity-notifications-testcase01.txt")
	checkError(err)
	reader := bufio.NewReaderSize(f, 1024*1024)
	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	expenditureTemp := strings.Split(readLine(reader), " ")

	expenditure := []int{}
	for i := 0; i < int(nTemp); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}
	return testdata{
		name:        "HR Real Test Case 01 - long",
		expenditure: expenditure,
		d:           int(dTemp),
		answer:      633,
	}
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
