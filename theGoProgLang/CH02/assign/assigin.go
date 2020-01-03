package main

func main() {
	x := gcd(3, 5)
	println(x)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
