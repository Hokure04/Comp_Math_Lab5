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
	var interpolation float64
	for i := 0; i < len(l); i++ {
		interpolation += l[i]
	}
	fmt.Printf("Приближенное значение функции по Лагранжу: %f\n", interpolation)
	fmt.Println()
}

func Newton_polynomial_divided_differences(xValues, yValues []float64, argX float64) {
	n := len(xValues)
	//sum := yValues[0]
	f := dividedDifferences(xValues, yValues)
	fmt.Println(f)
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
	fmt.Printf("Приближённое значение функции по Ньютону с разделёнными разностями: %f\n", interpolation)
	fmt.Println()
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
	var t float64
	var h float64
	var differences float64
	var y0 float64
	differences = argX - xValues[0]
	y0 = yValues[0]
	h = (xValues[len(xValues)-1] - xValues[0]) / float64(len(xValues)-1)
	fmt.Printf("Значение шага h: %f\n", h)
	for i := 0; i < len(xValues); i++ {
		if math.Abs(argX-xValues[i]) < differences {
			differences = argX - xValues[i]
			y0 = yValues[i]
		}
	}

	t = differences / h
	fmt.Printf("Значение параметра t: %f\n", t)

	deltaY := finiteDifferences(yValues)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "№\ty\tΔyi\tΔ^2yi\tΔ^3yi\tΔ^4yi")
	fmt.Fprintln(w, "-----\t-----\t----------\t---------\t----------\t----------")

	for i := range deltaY {
		fmt.Fprintf(w, "%d", i)
		for _, value := range deltaY[i] {
			fmt.Fprintf(w, "\t%.4f", value)
		}
		fmt.Fprintln(w)
	}
	w.Flush()

	var yArray []float64
	if t > 0 {
		for _, row := range deltaY {
			//fmt.Printf("row: %f", row)
			//fmt.Println(y0)
			if row[0] == y0 {
				yArray = append(yArray, row...)
			}
		}
	} else {
		var indices []int
		for i, row := range deltaY {
			if row[0] == y0 {
				indices = append(indices, i)
			}
		}

		for _, rowIndex := range indices {
			yArray = append(yArray, deltaY[rowIndex][0])
			for j := 1; j < len(deltaY[0]); j++ {
				if rowIndex > 0 {
					rowIndex--
					yArray = append(yArray, deltaY[rowIndex][j])
				}
			}
		}
	}

	fmt.Printf("Массив y: %f\n", yArray)
	var Nx float64
	var tIteration float64 = 1
	var factorial int = 1
	Nx += y0
	for i := 0; i < len(yArray); i++ {
		if i != 0 {
			if i == 0 {
				tIteration *= t
				Nx += yArray[i] * tIteration
			} else {
				if t > 0 {
					tIteration = tIteration * (t - float64(i-1))
				} else if t < 0 {
					//fmt.Printf("titeration: %f", tIteration)
					tIteration = tIteration * (t + float64(i-1))
					//fmt.Printf("tIteration: %f", tIteration)
				}
				factorial *= i
				Nx += (yArray[i] * tIteration) / float64(factorial)
			}
		}
	}

	fmt.Printf("Приближённое значение функции по Ньютону для равноотстоящих узлов: %f\n", Nx)
	fmt.Println()
	return Nx
}

func finiteDifferences(y []float64) [][]float64 {
	n := len(y)
	deltaY := make([][]float64, n)

	for i := range deltaY {
		deltaY[i] = make([]float64, n)
		deltaY[i][0] = y[i]
	}

	for j := 1; j < n; j++ {
		for i := 0; i < n-j; i++ {
			deltaY[i][j] = deltaY[i+1][j-1] - deltaY[i][j-1]
		}
	}

	return deltaY
}
