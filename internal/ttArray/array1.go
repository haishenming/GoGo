package ttArray

import "fmt"

func ShowSomeArray() {
	var a [2]int8

	a[0] = 1
	a[1] = 2

	fmt.Println(a)

	var b = [3]int8{1, 2, 3}

	fmt.Println(b)
}

func ShowSomeSlice() {
	//var a = [10]int8{1,2,3,4,5,6,7,8,9,0}
	//
	//var s []int8 = a[0: 4]
	//
	//fmt.Println(s)
	//
	//s[1] = 1
	//s[2] = 2
	//s[3] = 2
	//
	//fmt.Println(a)
	//fmt.Println(s)

	// int切片
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// bool切片
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	printSlice(q)

	var q2 []int
	printSlice(q2)

	if q2 == nil {
		fmt.Println("nil")
	}

	q3 := make([]int, 10)
	printSlice(q3)

	q4 := make([]int, 0, 10)
	printSlice(q4)

	for i := 0; i <= 20; i++ {
		q4 = append(q4, 1)
	}
	printSlice(q4)

	q4 = q4[:cap(q4)]
	printSlice(q4)

}

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
