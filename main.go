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

	var intX float64
	var err error
	for {
		fmt.Print("Введите значение x для расчёта значаения: ")
		argumentX := bufio.NewScanner(os.Stdin)
		argumentX.Scan()
		inputX := argumentX.Text()
		intX, err = strconv.ParseFloat(inputX, 64)
		if err == nil {
			break
		}
		fmt.Println("Ошибка: X должно быть целым числом")
	}

	lagrange_polynominal(xValues, yValues, intX)

}

func lagrange_polynominal(xValues, yValues []float64, argX float64) {
	var i int
	var l []float64
	for i < len(xValues) {
		var numerator float64 = 1
		var denominator float64 = 1

		j := 0
		for j < len(xValues) {
			if j != i {
				numerator *= argX - xValues[j]
				denominator *= xValues[i] - xValues[j]
			}
			j += 1
		}

		//fmt.Println(numerator)
		//fmt.Println(denominator)
		l = append(l, numerator/denominator*yValues[i])
		i += 1
	}
	fmt.Println(l)
	var interpolation float64
	i = 0
	for i < len(l) {
		interpolation += l[i]
		i += 1
	}
	fmt.Printf("Приближенное значение функции по Лагранжу: %f\n", interpolation)

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
