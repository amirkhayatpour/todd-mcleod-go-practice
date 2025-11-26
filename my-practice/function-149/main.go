package main

// closure

func main() {

	incrementor()
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
