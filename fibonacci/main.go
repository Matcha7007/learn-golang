package main

import (
	"fmt"
)

func main() {
	// Example usage of each pattern function
	rightHalfPyramid(5)
	// leftHalfPyramid(5)
	// fullPyramid(5)
	// invertedRightHalfPyramid(5)
	// invertedLeftHalfPyramid(5)
	// invertedFullPyramid(5)
	// rhombusPattern(5)
	// diamondPattern(5)
	// hourglassPattern(5)
	// hollowSquarePattern(5)
	// hollowFullPyramid(5)
	// hollowInvertedFullPyramid(5)
	// hollowDiamondPattern(5)
	// hollowHourglassPattern(5)
	// floydsTriangle(5)
	// pascalsTriangle(5)
}

func rightHalfPyramid(n int) {
	fmt.Println("Right Half Pyramid:")
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func leftHalfPyramid(n int) {
	fmt.Println("Left Half Pyramid:")
	for i := 1; i <= n; i++ {
		for j := n; j > i; j-- {
			fmt.Print("  ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func fullPyramid(n int) {
	fmt.Println("Full Pyramid:")
	for i := 1; i <= n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
}

func invertedRightHalfPyramid(n int) {
	fmt.Println("Inverted Right Half Pyramid:")
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func invertedLeftHalfPyramid(n int) {
	fmt.Println("Inverted Left Half Pyramid:")
	for i := n; i >= 1; i-- {
		for j := n; j > i; j-- {
			fmt.Print("  ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func invertedFullPyramid(n int) {
	fmt.Println("Inverted Full Pyramid:")
	for i := n; i >= 1; i-- {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
}

func rhombusPattern(n int) {
	fmt.Println("Rhombus Pattern:")
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= n; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func diamondPattern(n int) {
	fmt.Println("Diamond Pattern:")
	for i := 1; i <= n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := n - 1; i >= 1; i-- {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
}

func hourglassPattern(n int) {
	fmt.Println("Hourglass Pattern:")
	for i := n; i >= 1; i-- {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 2; i <= n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()
}

func hollowSquarePattern(n int) {
	fmt.Println("Hollow Square Pattern:")
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 || i == n || j == 1 || j == n {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func hollowFullPyramid(n int) {
	fmt.Println("Hollow Full Pyramid:")
	for i := 1; i <= n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func hollowInvertedFullPyramid(n int) {
	fmt.Println("Hollow Inverted Full Pyramid:")
	for i := n; i >= 1; i-- {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func hollowDiamondPattern(n int) {
	fmt.Println("Hollow Diamond Pattern:")
	for i := 1; i <= n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for i := n - 1; i >= 1; i-- {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func hollowHourglassPattern(n int) {
	fmt.Println("Hollow Hourglass Pattern:")
	for i := n; i >= 1; i-- {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	for i := 2; i <= n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 || i == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func floydsTriangle(n int) {
	fmt.Println("Floyd's Triangle:")
	num := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println()
	}
	fmt.Println()
}

func pascalsTriangle(n int) {
	fmt.Println("Pascal's Triangle:")
	for i := 0; i < n; i++ {
		// Print leading spaces
		for j := 0; j <= n-i-2; j++ {
			fmt.Print(" ")
		}
		// Calculate and print binomial coefficients
		val := 1
		for k := 0; k <= i; k++ {
			fmt.Printf("%d ", val)
			val = val * (i - k) / (k + 1)
		}
		fmt.Println()
	}
	fmt.Println()
}
