package main

import (
	"Comp_Math_Lab5/modules"
	"bufio"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	input_selection()
}

func Function(choice int, x float64) float64 {
	if choice == 1 {
		return 2*math.Pow(x, 2) - 5*x
	} else if choice == 2 {
		return math.Sin(x)
	} else if choice == 3 {
		return math.Sqrt(x)
	}
	return 0
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
			input_from_function()
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
	modules.Lagrange_polynominal(xValues, yValues, floatX)
	modules.Newton_polynomial_divided_differences(xValues, yValues, floatX)
	modules.Newton_polynomial_equally_spaced_notes(xValues, yValues, floatX)
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
		modules.Lagrange_polynominal(xValues, yValues, floatX)
		modules.Newton_polynomial_divided_differences(xValues, yValues, floatX)
		modules.Newton_polynomial_equally_spaced_notes(xValues, yValues, floatX)
		break
	}
}

func input_from_function() {
	var xValues []float64
	var yValues []float64

	fmt.Println("Выберите функцию, которую хотите использовать:")
	fmt.Println("1. 2*x^2 - 5*x")
	fmt.Println("2. sin(x)")
	fmt.Println("3. √x")

	for {
		choice := bufio.NewScanner(os.Stdin)
		choice.Scan()
		input := choice.Text()

		var choiceInt int
		_, err := fmt.Sscanf(input, "%d", &choiceInt)
		if err != nil {
			fmt.Println("Ошибка: Вы ввели некорректное значение")
			continue
		}

		if choiceInt > 3 || choiceInt < 1 {
			fmt.Println("Введите значение от 1 до 3")
			continue
		}

		fmt.Print("Введите количество точек: ")
		pointsStr := bufio.NewScanner(os.Stdin)
		pointsStr.Scan()
		pointsInput := pointsStr.Text()
		points, err := strconv.Atoi(pointsInput)
		if err != nil || points <= 2 {
			fmt.Println("Ошибка: Количество точек должно быть больше 2")
			continue
		}

		var a, b float64
		for {
			fmt.Print("Введите интервал (a b): ")
			intervalStr := bufio.NewScanner(os.Stdin)
			intervalStr.Scan()
			intervalInput := intervalStr.Text()
			intervalParts := strings.Split(intervalInput, " ")
			if len(intervalParts) != 2 {
				fmt.Println("Ошибка: Вы ввели некорректный интервал")
				continue
			}
			a, err = strconv.ParseFloat(intervalParts[0], 64)
			if err != nil {
				fmt.Println("Ошибка: Вы ввели некорректный интервал")
				continue
			}
			b, err = strconv.ParseFloat(intervalParts[1], 64)
			if err != nil {
				fmt.Println("Ошибка: Вы ввели некорректный интервал")
				continue
			}
			break
		}

		h := (b - a) / float64(points-1)

		for i := 0; i < points; i++ {
			xValues = append(xValues, a+h*float64(i))
			yValues = append(yValues, Function(choiceInt, xValues[i]))
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
		modules.Lagrange_polynominal(xValues, yValues, floatX)
		modules.Newton_polynomial_divided_differences(xValues, yValues, floatX)
		modules.Newton_polynomial_equally_spaced_notes(xValues, yValues, floatX)
		DrawGraph(xValues, yValues, floatX, choiceInt)
		break
	}
}

func DrawGraph(xValues, yValues []float64, argX float64, choiceInt int) {
	dirPath := "graphs"
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	fileName := fmt.Sprintf("graph_%s.png", time.Now().Format("2006-01-02_15-04-05"))
	filePath := filepath.Join(dirPath, fileName)

	f := func(x float64) float64 {
		return Function(choiceInt, x)
	}

	xVals := plotter.XYs{}
	for x := xValues[0]; x <= xValues[len(xValues)-1]; x += 0.1 {
		xVals = append(xVals, plotter.XY{X: x, Y: f(x)})
	}

	p := plot.New()
	p.Title.Text = "График функции"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	line, err := plotter.NewLine(xVals)
	if err != nil {
		panic(err)
	}
	line.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	p.Add(line)

	pts := plotter.XYs{}
	for i := range xValues {
		pts = append(pts, plotter.XY{X: xValues[i], Y: yValues[i]})
	}

	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	scatter.GlyphStyle.Color = color.RGBA{G: 255, A: 255}
	scatter.GlyphStyle.Radius = vg.Points(5)
	p.Add(scatter)

	xVals = plotter.XYs{}
	for x := xValues[0]; x <= xValues[len(xValues)-1]; x += 0.1 {
		xVals = append(xVals, plotter.XY{X: x, Y: newtonPolynomial(xValues, yValues, x, choiceInt)})
	}

	line, err = plotter.NewLine(xVals)
	if err != nil {
		panic(err)
	}
	line.Color = color.RGBA{B: 255, A: 255}
	p.Add(line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, filePath); err != nil {
		panic(err)
	}
}

func newtonPolynomial(xValues, yValues []float64, argX float64, choiceInt int) float64 {
	n := len(xValues)
	f := divided(xValues, yValues, choiceInt)
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

func divided(x, y []float64, choiceInt int) []float64 {
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
