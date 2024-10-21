package main

func main() {
	a := 3
	a = 4
	if a == 4 {
		println(a)
		a := 5
		println(a)
	}

	// This print 5 which is wrong, should print 4
	println(a)
}
