package fib

func FibFunc() func() uint64 {
	var a,b uint64 = 0, 1 // yes, it is wrong but we need a failing unit test
	return func () uint64 {
		a,b = b, a+b
		return a;
	}
}

func Mul(a uint64, b uint64) uint64 {
	return a*b;
}