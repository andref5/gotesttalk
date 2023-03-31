package collatz

import "fmt"

// The Collatz conjecture states that the hailstone
// sequence of every positive number reaches 1
// if we execute the Collatz process on `x = 10`, then we get the
// sequence of numbers `10, 5, 16, 8, 4, 2, 1`. This sequence of numbers is called
// the hailstone sequence of 10.
// Nobody knows if the Collatz conjecture is true
// it is one of the most famous unsolved problems in mathematics.
func Collatz(x int) (hailstone []int, err error) {
	for x != 1 {
		hailstone = append(hailstone, x)
		if x%2 == 0 {
			x, err = Even(x)
		} else {
			x, err = Odd(x)
		}
		if err != nil {
			return nil, err
		}
	}
	hailstone = append(hailstone, x)
	return hailstone, nil
}

func Even(x int) (int, error) {
	if x <= 0 {
		return 0, fmt.Errorf("%d is not positive", x)
	}
	if x%2 != 0 {
		return 0, fmt.Errorf("%d is not even", x)
	}
	return x / 2, nil
}

func Odd(x int) (int, error) {
	if x <= 0 {
		return 0, fmt.Errorf("%d is not positive", x)
	}
	if x%2 == 0 {
		return 0, fmt.Errorf("%d is not odd", x)
	}
	return 3*x + 1, nil
}
