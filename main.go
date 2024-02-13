package main

import (
	"fmt"
	"math"
	"os"
)

var a = -6.3 //левая граница интервала локализации
var b = 13.  //правая граница интервала локализации

var e_ = 0.1
var arrList5 = [4]int{0, 0, 0, 0}

// коэффициенты уравнения вида: a0 + a1*x + a2*sin(a3*x)
const a0 = 4.07
const a1 = -0.66
const a2 = 3.96
const a3 = 1.07

func main() {
	list1()
	list2()
	list3_4()
	list_5()
}

func list1() {
	file, err := os.Create("лист1.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Коэффициенты уравнения вида: a0 + a1*x + a2*sin(a3*x)\n")
	fmt.Fprintf(file, "a0 = %.2f\na1 = %.2f\na2 = %.2f\na3 = %.2f\n~~~~~~~~~~~~~~~~~~~~~~\n", a0, a1, a2, a3)
	file.WriteString("Интервал локализации решения\n")
	fmt.Fprintf(file, "a = %.1f\nb = %.1f", a, b)
	fmt.Println("list 1 is done!")
}

func list2() {
	const step = 0.5
	var derivValues []float64
	var x = a
	for x < b {
		derivValues = append(derivValues, derivative(x, e_))
		x += step
	}
	derivValues = append(derivValues, derivative(b, e_))

	file, err := os.Create("лист2.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Координаты глобальных и локальных минимумов\nи максимумов (с точностью до 0.5) в указанном интервале локализации\n\n")
	file.WriteString("  x   |   R(x)  |   R'(x)  \n" + "----------------------\n")

	x = a

	for i := 1; i < len(derivValues)-1; i++ {
		if derivValues[i]*derivValues[i-1] < 0 {
			fmt.Fprintf(file, " %.2f | %.2f | %.2f\n", x+step, R(x+step), derivValues[i])
		}
		x += step
	}

	fmt.Println("list 2 is done!")
}

func list3_4() {
	e := [4]float64{0.1, 0.01, 0.001, 0.0001} //точности

	for j := 0; j < len(e); j++ {
		//создание файла
		fileName := fmt.Sprintf("лист3_4_%g.txt", e[j])
		file, err := os.Create(fileName)

		if err != nil {
			fmt.Println("Unable to create file:", err)
			return
		}
		defer file.Close()
		file.WriteString("\t\t\t\tМетод чисел Фибоначчи\n")
		file.WriteString("  a   |  x1   |  dx   |   x2  |   b   | F_s_k_1| F_s_k | R(x1) |  R(x2)\n")
		file.WriteString("-----------------------------------------------------------------------\n")

		var a_ = a
		var b_ = b

		var N = int((b_ - a_) / e[j])
		var fibSeq = fibonacciUntilN(N)

		var F_s_k = fibSeq[len(fibSeq)-2]
		var F_s_k_1 = fibSeq[len(fibSeq)-3]
		var F_s = fibSeq[len(fibSeq)-1]
		var dx = (b_ - a_) / float64(F_s)
		var x1 = a_ + dx*float64(F_s_k_1)
		var x2 = a_ + dx*float64(F_s_k)

		for i := len(fibSeq) - 2; i >= 4; i-- {
			var Rx1 = R(x1)
			var Rx2 = R(x2)
			fmt.Fprintf(file, " %.2f | %.3f | %f | %.3f | %.2f |  %d  | %d | %.3f | %.3f\n", a_, x1, dx, x2, b_, F_s_k_1, F_s_k, Rx1, Rx2)

			F_s_k = fibSeq[i-1]
			F_s_k_1 = fibSeq[i-2]

			if Rx1 < Rx2 {
				b_ = x2
				x2 = x1
				x1 = a_ + dx*float64(F_s_k_1)
			}
			if Rx2 < Rx1 {
				a_ = x1
				x1 = x2
				x2 = a_ + dx*float64(F_s_k)
			}
			if Rx2 == Rx1 {
				a_ = x1
				b_ = x2
				x1 = a_ + dx*float64(F_s_k_1)
				x2 = a_ + dx*float64(F_s_k)
			}
		}

		var a2 = b_
		var b2 = b

		F_s_k = fibSeq[len(fibSeq)-2]
		F_s_k_1 = fibSeq[len(fibSeq)-3]
		F_s = fibSeq[len(fibSeq)-1]
		dx = (b2 - a2) / float64(F_s)
		x1 = a2 + dx*float64(F_s_k_1)
		x2 = a2 + dx*float64(F_s_k)

		for i := len(fibSeq) - 2; i >= 4; i-- {
			arrList5[j] += 1
			var Rx1 = R(x1)
			var Rx2 = R(x2)
			fmt.Fprintf(file, " %.2f | %.3f | %f | %.3f | %.2f |  %d  | %d | %.3f | %.3f\n", a2, x1, dx, x2, b2, F_s_k_1, F_s_k, Rx1, Rx2)

			F_s_k = fibSeq[i-1]
			F_s_k_1 = fibSeq[i-2]

			if Rx1 < Rx2 {
				b2 = x2
				x2 = x1
				x1 = a2 + dx*float64(F_s_k_1)
			}
			if Rx2 < Rx1 {
				a2 = x1
				x1 = x2
				x2 = a2 + dx*float64(F_s_k)
			}
			if Rx2 == Rx1 {
				a2 = x1
				b2 = x2
				x1 = a2 + dx*float64(F_s_k_1)
				x2 = a2 + dx*float64(F_s_k)
			}
		}
	}
	fmt.Println("list 3_4 is done!")
}

func list_5() {
	//создание txt файла
	file, err := os.Create("лист5.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		return
	}
	defer file.Close()

	//запись в файл
	file.WriteString("Зависимость количества шагов от выбранной точности\n--------------------------------------------------\n")
	file.WriteString("\tТочность | Количество шагов\n")
	e := [4]float64{0.1, 0.01, 0.001, 0.0001} //точности
	for i := 0; i < len(e); i++ {
		fmt.Fprintf(file, "\t%.4f   | %d\n", e[i], arrList5[i])
	}
	fmt.Println("list 5 is done!")
}

func R(x float64) float64 {
	//return math.Pow(x-6, 2)
	return a0 + a1*x + a2*math.Sin(a3*x)
}

func derivative(x float64, e float64) float64 {
	return (R(x+e) - R(x-e)) / (2 * e)
}

func fibonacciUntilN(N int) []int {
	seq := []int{0, 1}

	for seq[len(seq)-1] <= N {
		next := seq[len(seq)-1] + seq[len(seq)-2]
		if next > N {
			break
		}
		seq = append(seq, next)
	}
	// Добавляем еще одно число в последовательность Фибоначчи, если последнее число меньше или равно N
	if seq[len(seq)-1] <= N {
		seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
	}

	return seq
}
