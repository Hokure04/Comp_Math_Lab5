package main

import (
	"bufio"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"os"
	"strconv"
)

func main() {
	//draw_function()
	input_selection()
}

func input_selection() {
	for {
		fmt.Println("1. Ввести набор точек")
		fmt.Println("2. Данные из файла")
		fmt.Println("3. Выбрать функцию")
		fmt.Println("Введите способ ввода:")
		choice := bufio.NewScanner(os.Stdin)
		choice.Scan()
		input := choice.Text()

		var choiceInt int
		_, err := fmt.Sscanf(input, "%d", &choiceInt)
		if err != nil {
			fmt.Println("Ошибка: Вы ввели некорректное значение")
		}
		if choiceInt == 1 {
			hand_input()
		} else if choiceInt == 2 {
			fmt.Println("Заглушка 2")
		} else if choiceInt == 3 {
			fmt.Println("Заглушка 3")
		} else {
			fmt.Println("Введите значение от 1 до 3")
		}
	}

}

func hand_input() {
	var xValues []float64
	var yValues []float64

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ведите значения x и y через пробел")

	for scanner.Scan() {
		input := scanner.Text()
		if input == "stop" {
			break
		}

		var x, y float64
		_, err := fmt.Sscanf(input, "%f %f", &x, &y)

		if err != nil {
			fmt.Println("Ошибка: Вы ввели x и y некорректно")
			continue
		}

		xValues = append(xValues, x)
		yValues = append(yValues, y)
	}

	fmt.Println("Введённые значения:")
	fmt.Println("X: ", xValues)
	fmt.Println("Y: ", yValues)

	var floatX float64
	var err error
	for {
		fmt.Print("Введите значение x для расчёта значаения: ")
		argumentX := bufio.NewScanner(os.Stdin)
		argumentX.Scan()
		inputX := argumentX.Text()
		floatX, err = strconv.ParseFloat(inputX, 64)
		if err == nil {
			break
		}
		fmt.Println("Ошибка: X должно быть целым числом")
	}
	fmt.Println()
	lagrange_polynominal(xValues, yValues, floatX)
	newton_polynomial_divided_differences(xValues, yValues, floatX)

}

func lagrange_polynominal(xValues, yValues []float64, argX float64) {
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

func newton_polynomial_divided_differences(xValues, yValues []float64, argX float64) {
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

func draw_function() {
	f := func(x float64) float64 {
		return x*x + 2*x
	}

	// Generate data points for the graph
	xVals := plotter.XYs{}
	for x := 0.0; x <= 100.0; x += 0.1 {
		xVals = append(xVals, plotter.XY{X: x, Y: f(x)})
	}

	// Create a plot
	p := plot.New()
	p.Title.Text = "Function graph"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Add the data points to the plot
	line, err := plotter.NewScatter(xVals)
	if err != nil {
		panic(err)
	}
	line.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	p.Add(line)

	// Save the plot to a file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "graph.png"); err != nil {
		panic(err)
	}
}
