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

//斐波那契数列，又称黄金分割数列、神秘数列，因数学家列昂纳多·斐波那契（Leonardoda Fibonacci）以兔子繁殖为例子而引入，故又称为“兔子数列”。
//1、1、2、3、5、8、13、21、34、……
//F(1)=1，F(2)=1, F(n)=F(n-1)+F(n-2)（n>=3，n∈N*）
//第一位是1，每个数等于前两两个数相加；前后两个数相除接近0.618或1.618。
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
