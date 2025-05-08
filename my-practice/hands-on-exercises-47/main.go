package main

import "fmt"

func main() {
	//
	//x := make([]int, 50)
	//fmt.Println(x)
	//fmt.Println(len(x))
	//fmt.Println(cap(x))
	//x = append(x, 99)
	//fmt.Println(x)
	//fmt.Println(len(x))
	//fmt.Println(cap(x))
	//
	//fmt.Println("--------------------------")
	//
	//y := make([]int, 0, 50)
	//fmt.Println(y)
	//fmt.Println(len(y))
	//fmt.Println(cap(y))
	//y = append(y, 99)
	//fmt.Println(y)
	//fmt.Println(len(y))
	//fmt.Println(cap(y))
	masood := make([]string, 0, 50)
	masood = append(masood, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(masood))
	fmt.Println(cap(masood))

	for i := 0; i < len(masood); i++ {
		fmt.Println(masood[i], i)
	}
}

/*

For this exercise, do the following:
● Create a slice to store the names of all of the states in the United States of America.
○ Use make and append to do this.
○ Goal: do not have the array that underlies the slice created more than once.
● Print out
○ the len
○ the cap
○ the values, along with their index position, without using the range clause.
● Here is a list of the 50 states:
`
,
,
` Alabama
` Alaska
`
` Arizona
`
` Arkansas
`
` California
`
` Colorado
`
` Connecticut`
`
,
,
,
,
,
,
Delaware
`
` Florida
`
,
` Georgia
`
` Hawaii`
` Idaho
`
` Illinois
`
` Indiana
`
` Iowa
`
` Kansas
`
`
,
,
,
,
,
,
,
Kentucky
`
` Louisiana
`
` Maine
`
,
,
,
` Maryland`
` Massachusetts
`
,
,
` Michigan
`
` Minnesota
`
`
,
,
Mississippi`
` Missouri`
` Montana
`
` Nebraska
`
` Nevada
`
,
,
,
,
,
` New Hampshire
`
,
` New Jersey
`
` New Mexico
`
` New York`
` North Carolina
`
` North Dakota
`
` Ohio
`
` Oklahoma
`
,
,
,
,
,
,
` Oregon
`
` Pennsylvania
`
` Rhode Island`
` South Carolina
`
` South Dakota
`
` T ennessee
`
` T exas
`
`
,
,
,
,
,
,
Utah`
` Vermont`
,
,
` Virginia
`
,
` Washington
`
,
` West Virginia
`
` Wisconsin
`
,
,
` Wyoming
`
,

*/
