package main

import "fmt"

func main() {
	//data type
	//int, int32, int64, float32, float64,
	// string, bool,byte, date,slice
	// var a int = 10
	// b := 20
	// //- + * / % ++ -- **
	// fmt.Printf("a + b = %d", a+b)
	// var name string = "Reaksmey Kevin"
	// replacedName := strings.ReplaceAll(name, "Kevin", "Thkeam")
	// fmt.Printf("Hello, %s", replacedName)

	// arr := []int{1, 2, 3, 4, 5} // growable array
	// //arr.push(6)
	// //arr.add(6)

	// arr = append(arr, 6)
	// for _, v := range arr {
	// 	println(v)
	// }
	// arr := make([]int, 5) // create an array with length 5
	// arr[0] = 1
	// arr[1] = 2
	// arr[2] = 3
	// arr[3] = 4
	// arr[4] = 5
	//PointerA a x01234
	//reused variable
	// var a int = 20
	// var b = &a
	// *b = 30
	// var c = b
	person1 := Person{
		Name: "Reaksmey Kevin",
		Age:  30,
		Dob:  "1993-01-01",
	}
	//object , class
	person2 := Person{
		Name: "Thkeam Reaksmey",
		Age:  25,
		Dob:  "1998-01-01",
	}
	fmt.Printf("Person 1: %+v\n", person1)
	fmt.Printf("Person 2: %+v\n", person2)

}

type Person struct {
	Name string
	Age  int
	Dob  string
}

// class Student {}
//TEST SUIT
