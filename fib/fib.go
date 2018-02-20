package fib

//Prove for base only and rest assume it all works for others
// 2 Base case

func Fib(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return Fib(n-2) + Fib(n-1)
}

var fib map[int]int

func init() {
	fib = map[int]int{0: 1, 1: 1}

}

func FibEfficient(n int) int {
	if v, ok := fib[n]; ok {
		return v
	}
	a := Fib(n-2) + Fib(n-1)
	fib[n] = a // Memoize ... caches
	return a
}
