package modules

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

func Lagrange_polynominal(xValues, yValues []float64, argX float64) {
	var l []float64
	for i := 0; i < len(xValues); i++ {
		var numerator float64 = 1
		var denominator float64 = 1

		for j := 0; j < len(xValues); j++ {
			if j != i {
				numerator *= argX - xValues[j]
				denominator *= xValues[i] - xValues[j]
			}
		}

		//fmt.Println(numerator)
		//fmt.Println(denominator)
		l = append(l, numerator/denominator*yValues[i])
	}
	fmt.Println(l)
	fmt.Println()
	var interpolation float64
	for i := 0; i < len(l); i++ {
		interpolation += l[i]
	}
	fmt.Printf("Приближенное значение функции по Лагранжу: %f\n", interpolation)
	fmt.Println()
}

func Newton_polynomial_divided_differences(xValues, yValues []float64, argX float64) float64 {
	n := len(xValues)
	f := dividedDifferences(xValues, yValues)
	for i := 0; i < n; i++ {
		var finiteDifferences []float64
		var intermediateCalc float64 = 1
		for j := 0; j < i; j++ {
			intermediateCalc *= argX - xValues[j]
			finiteDifferences = append(finiteDifferences, intermediateCalc)
		}
		if i != 0 {
			f[i] = f[i] * finiteDifferences[i-1]
		}
	}
	var interpolation float64
	for i := 0; i < len(f); i++ {
		interpolation += f[i]
	}
	return interpolation
}

func dividedDifferences(x, y []float64) []float64 {
	n := len(x)
	f := make([]float64, n)

	for i := 0; i < n; i++ {
		f[i] = y[i]
	}
	for i := 1; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			f[j] = (f[j] - f[j-1]) / (x[j] - x[j-i])
		}
	}
	return f
}

func Newton_polynomial_equally_spaced_notes(xValues, yValues []float64, argX float64) float64 {
	var x0 int
	var isLessThanZero bool
	var factorial float64 = 1
	var result float64
	if len(xValues) == 0 {
		panic("xValues is empty")
	}

	n := len(xValues)
	h := xValues[1] - xValues[0]
	array := make([][]float64, n)

	for i := 0; i < n; i++ {
		array[i] = make([]float64, n)
		array[i][0] = yValues[i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			array[j][i] = array[j+1][i-1] - array[j][i-1]
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "№\ty\tΔyi\tΔ^2yi\tΔ^3yi\tΔ^4yi")
	fmt.Fprintln(w, "-----\t-----\t----------\t---------\t----------\t----------")

	for i := range array {
		fmt.Fprintf(w, "%d", i)
		for _, value := range array[i] {
			fmt.Fprintf(w, "\t%.4f", value)
		}
		fmt.Fprintln(w)
	}
	w.Flush()

	if argX <= xValues[n/2] {
		isLessThanZero = true
		x0 = n - 1
		for i := 0; i < n; i++ {
			if argX <= xValues[i] {
				x0 = i - 1
				break
			}
		}
		if x0 < 0 {
			x0 = 0
		}
		t := (argX - xValues[x0]) / h
		//fmt.Printf("Значение t: %f\n", t)
		result = array[x0][0]
		//fmt.Printf("Значение result= %f\n", result)
		for i := 1; i < n; i++ {
			factorial *= float64(i)
			result += (t_calculate(t, i, isLessThanZero) * array[x0][i]) / factorial
			//fmt.Printf("Значение result= %f\n", result)
		}
	} else {
		isLessThanZero = false
		t := (argX - xValues[n-1]) / h
		//fmt.Printf("значение t: %f\n", t)
		result = array[n-1][0]
		//fmt.Printf("значение result= %f\n", result)
		for i := 1; i < n; i++ {
			factorial *= float64(i)
			result += (t_calculate(t, i, isLessThanZero) * array[n-i-1][i]) / factorial
			//fmt.Printf("значение t result= %f\n", result)
		}
	}
	return result
}

func t_calculate(t float64, n int, isLessThanZero bool) float64 {
	result := t
	for i := 1; i < n; i++ {
		if isLessThanZero {
			result *= t - float64(i)
		} else {
			result *= t + float64(i)
		}
	}
	return result
}

func Stirling_polynomial(xValues, yValues []float64, argX float64) {
	var x0 int
	var factorial float64 = 1
	var pr_number int
	if len(xValues) == 0 {
		panic("xValues is empty")
	}
	n := len(xValues)
	h := xValues[1] - xValues[0]

	array := make([][]float64, n)
	for i := 0; i < n; i++ {
		array[i] = make([]float64, n)
		array[i][0] = yValues[i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			array[j][i] = array[j+1][i-1] - array[j][i-1]
		}
	}
	if n%2 == 0 {
		x0 = int(n/2 - 1)
	} else {
		x0 = int(n / 2)
	}

	t := (argX - xValues[x0]) / h
	if math.Abs(t) > 0.25 {
		fmt.Println("Так как t > 0.25, результат метода может быть неточным")
	}

	interpolation := array[x0][0]
	compT1 := t
	compT2 := math.Pow(t, 2)
	pr_number = 0
	for i := 1; i < n; i++ {
		factorial *= float64(i)
		if i%2 == 0 {
			if x0-(i/2) < 0 {
				break
			}
			interpolation += (compT2 / factorial) * array[x0-(i/2)][i]
			compT2 *= t*t - float64(pr_number*pr_number)
		} else {
			if x0-((i+1)/2) < 0 {
				break
			}
			if x0-(((i+1)/2)-1) < 0 {
				break
			}
			interpolation += (compT1 / factorial) * ((array[x0-((i+1)/2)][i] + array[x0-(((i+1)/2)-1)][i]) / 2)
			pr_number += 1
			compT1 *= t*t - float64(pr_number*pr_number)
		}
	}
	fmt.Printf("Приближенное значение функции по Стирлингу: %f\n", interpolation)
	fmt.Println()
}

func Bessel_polynomila(xValues, yValues []float64, argX float64) {
	var x0 int
	var lastNumber float64
	var factorial float64 = 1
	if len(xValues) == 0 {
		panic("xValues is empty")
	}
	n := len(xValues)
	h := xValues[1] - xValues[0]
	array := make([][]float64, n)
	for i := 0; i < n; i++ {
		array[i] = make([]float64, n)
		array[i][0] = yValues[i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			array[j][i] = array[j+1][i-1] - array[j][i-1]
		}
	}
	if n%2 == 0 {
		x0 = int(n/2 - 1)
	} else {
		x0 = int(n / 2)
	}
	t := (argX - xValues[x0]) / h
	if math.Abs(t) < 0.25 {
		fmt.Println("Так как t < 0.25, результат метода может быть неточным")
	} else if math.Abs(t) > 0.75 {
		fmt.Println("Так как t > 0.75, результат метода может быть неточным")
	}
	interpolation := ((array[x0][0] + array[x0+1][0]) / 2) + (t-0.5)*array[x0][1]
	compT := t
	lastNumber = 0
	for i := 2; i < n; i++ {
		factorial *= float64(i)
		if i%2 == 0 {
			lastNumber += 1
			compT *= t - lastNumber
			interpolation += (compT / factorial) * (array[x0-i/2][i] + array[x0-((i/2)-1)][i]) / 2
		} else {
			interpolation += (compT * (t - 0.5) / factorial) * array[x0-((i-1)/2)][i]
			compT *= t + lastNumber
		}
	}
	fmt.Printf("Приближенное значение функции по Бесселю: %f\n", interpolation)
	fmt.Println()
}
