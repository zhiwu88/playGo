package main

func main() {
	x := gcd(12, 15)
	println(x)
	fibn := fib(5)
	println(fibn)
}

//辗转相除法求最大公约数（又名欧几里德法）
//一个数除以另一个数，要是比另一个数小的话，商为0，余数就是它自己。
/*
x, y = 15, 12mod15=12
x, y = 12, 15mod12=3
x, y = 3, 12mod3=0
*/
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
