package main

func main() {
	if true {
		println("Test 3: i =", 0)
	}

	if false {
		println("Test 3: i =", 1)
	} else if true {
		println("Test 3: i =", 1)
	}

	if false {
		println("Test 3: i =", 2)
	} else if false {
		println("Test 3: i =", 2)
	} else {
		println("Test 3: i =", 2)
	}
}
