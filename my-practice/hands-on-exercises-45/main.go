package main

import "fmt"

func main() {

	masood := []int{
		42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}

	masood = append(masood, 52)
	fmt.Println(masood)

	masood = append(masood, 53, 54, 55)
	fmt.Println(masood)

	smallMasood := []int{
		56, 57, 58, 59, 60,
	}

	masood = append(masood, smallMasood...)
	fmt.Println(masood)
	/*
		● start with this slice
		○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		● append to that slice this value
		○ 52
		● print out the slice
		● in ONE STATEMENT append to that slice these values
		○ 53
		○ 54
		○ 55
		● print out the slice
		● append to the slice this slice
		○ y := []int{56, 57, 58, 59, 60}
		● print out the slice
	*/

}
