package main

import (
	"bufio"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
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
			read_from_file()
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
	newton_polynomial_equally_spaced_notes(xValues, yValues, floatX)
}

func read_from_file() {
	var xValues, yValues []float64
	for {
		fmt.Println("Введите путь к файлу: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		filename := scanner.Text()
		file, err := os.Open(filename)

		if err != nil {
			fmt.Println("Ошибка: Неверный путь к файлу")
			continue
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.Split(line, " ")
			if len(parts) != 2 {
				fmt.Println("Ошибка: Формат данных должен быть: x y")
				break
			}
			f1, err := strconv.ParseFloat(parts[0], 64)
			if err != nil {
				fmt.Println("Ошибка: В файле находятся неверные данные")
				break
			}
			f2, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				fmt.Println("Ошибка: В файле находятся неверные данные")
				break
			}

			xValues = append(xValues, f1)
			yValues = append(yValues, f2)

		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка: не получилось сканировать файл")
			break
		}

		fmt.Println("Введённые значения:")
		fmt.Println("X: ", xValues)
		fmt.Println("Y: ", yValues)
		fmt.Println()

		var floatX float64
		var inputX string
		for {
			fmt.Print("Введите значение x для расчёта значения: ")
			argumentX := bufio.NewScanner(os.Stdin)
			argumentX.Scan()
			inputX = argumentX.Text()
			floatX, err = strconv.ParseFloat(inputX, 64)
			if err == nil {
				break
			}
			fmt.Println("Ошибка: X должно быть числом")
		}
		lagrange_polynominal(xValues, yValues, floatX)
		newton_polynomial_divided_differences(xValues, yValues, floatX)
		newton_polynomial_equally_spaced_notes(xValues, yValues, floatX)
		break
	}
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

func newton_polynomial_equally_spaced_notes(xValues, yValues []float64, argX float64) {
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
			fmt.Println(xValues[i])
			y0 = xValues[i]
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
	for _, row := range deltaY {
		//fmt.Printf("row: %f", row)
		if row[0] == y0 {
			yArray = append(yArray, row...)
		}
	}
	fmt.Printf("Массив y: %f\n", yArray)
	var Nx float64
	var tIteration float64 = 1
	var factorial int = 1
	Nx += y0
	if t > 0 {
		for i := 0; i < len(yArray); i++ {
			if i != 0 {
				if i == 0 {
					tIteration *= t
					Nx += yArray[i] * tIteration
				} else {
					if t > 0 {
						tIteration = tIteration * (t - float64(i-1))
					} else if t < 0 {
						tIteration = tIteration * (t + float64(i-1))
					}
					factorial *= i
					Nx += (yArray[i] * tIteration) / float64(factorial)
				}
			}
		}
	}

	fmt.Printf("Приближённое значение функции по Ньютону для равноотстоящих узлов: %f\n", Nx)
	fmt.Println()
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
