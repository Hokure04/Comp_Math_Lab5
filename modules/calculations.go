package modules

import (
	"fmt"
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
