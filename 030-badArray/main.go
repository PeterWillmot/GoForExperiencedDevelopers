package main

func main() {
	var arr [5]int
	i := 6

	arr[6] = 3 // compile error : index 6 out of bounds [0:5]

	arr[i] = 3 // panic : exit status 2
}
