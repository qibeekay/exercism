package differenceofsquares

func SquareOfSum(n int) int {
	number := 0

    for i:=1; i<=n; i++ {
        number += i
    }

    return number * number
}

func SumOfSquares(n int) int {
	number := 0

    for i:=1; i<=n; i++ {
        number += i*i
    }

    return number
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n) 
}
