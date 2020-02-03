package main

import (
    "fmt"
	"github.com/tealeg/xlsx"
	"math"
	"strconv"
)

const eps = 0.01
const n = 2
const m = 20
var x[n]float64
var p[n]float64
var b[n]float64
var a[n][m]float64
const startPrice = 6767
const count = 80

func main() {
    excelFileName := "data.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        fmt.Printf("Error")
	}
	
    for _, sheet := range xlFile.Sheets {
		var rowIndex = 0
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
				text := cell.String()
				var i = 0
				if text == "Бензин АИ-92" {
					fmt.Printf("%s\n", text)
					var intValue, intError = sheet.Cell(rowIndex, 28).Int()
					var floatQValue, floatQError = sheet.Cell(rowIndex, 30).Float()
					var floatWValue, floatWError = sheet.Cell(rowIndex, 29).Float()
					fmt.Printf("%s\n", intError)
					fmt.Printf("%s\n", floatQError)
					fmt.Printf("%s\n", floatWError)
					a[i][0] = b[0] * float64(intValue) * floatQValue
					a[i][1] = b[1]
					b[i] = floatWValue
					i++
				}
            }
			rowIndex = rowIndex + 1
        }
    }
	iterate()
}



// Условие окончания
func converge(xk [n]float64, xkp [n]float64) bool {
    var norm float64 = 0;
    for i := 0; i < n; i++ {
		norm = norm + (xk[i] - xkp[i])*(xk[i] - xkp[i])
	}
    return (math.Sqrt(norm) < eps);
}

/*
    Ход метода, где:
    a[n][n] - Матрица коэффициентов
    x[n], p[n] - Текущее и предыдущее решения
    b[n] - Столбец правых частей
    Все перечисленные массивы вещественные и
    должны быть определены в основной программе,
    также в массив x[n] следует поместить начальное
    приближение столбца решений (например, все нули)
*/

func iterate() {
	for {
		fmt.Printf(strconv.FormatBool(converge(x, p)))
		fmt.Printf("\n")
		for i := 0; i < n; i++ {
			p[i] = x[i];
		}
		for i := 0; i < n; i++ {
			var varr float64 = 0;
			for j := 0; j < i; j++ {
				varr += varr + (a[i][j] * x[j])
			}
			for j := i + 1; j < n; j++{
				varr += varr + (a[i][j] * p[j])
			}
			x[i] = (b[i] - varr) / a[i][i]
		}
		if converge(x, p) {
			break
		}
	}
	var price = b[0] * startPrice * count + b[1]
	fmt.Printf("%s\n", price)
}






func solveE(startPrice float64, count int, soldedPrice float64) float64 {
	var y = fun(startPrice, count, soldedPrice)
	var delta = soldedPrice - y
	var e = delta * delta
	return e
}

func fun(startPrice float64, count int, soldedPrice float64) float64 {
	var comp = soldedPrice - b[0] * startPrice * float64(count) + b[1]
	return comp
}