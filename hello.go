package main

import (
	"fmt"
	"math/rand"

	"github.com/goml/gobrain"
)

func main() {
	// set the random seed to 0
	rand.Seed(0)

	// excelFileName := "data.xlsx"
	// xlFile, err := xlsx.OpenFile(excelFileName)
	// if err != nil {
	// 	fmt.Printf("Error")
	// }

	// create the XOR representation patter to train the network
	//patterns := make([][][]float64, 1000000)
	patterns := [][][]float64{
		{{500, 50.74}, {51}},
		{{2000, 45.12}, {48}},
		{{4000, 48.75}, {49}},
		{{300, 47.88}, {51}},
		{{900, 48.58}, {51.7}},
		{{450, 46.08}, {48}},
		{{1200, 44.76}, {51}},
	}

	// for _, sheet := range xlFile.Sheets {
	// 	var rowIndex = 0
	// 	for _, row := range sheet.Rows {
	// 		for _, cell := range row.Cells {
	// 			text := cell.String()
	// 			if text == "Бензин АИ-92" {
	// 				fmt.Printf("%s\n", text)
	// 				var floatPValue, intError = sheet.Cell(rowIndex, 28).Float()
	// 				var floatQValue, floatQError = sheet.Cell(rowIndex, 30).Float()
	// 				var floatWValue, floatWError = sheet.Cell(rowIndex, 29).Float()
	// 				fmt.Println(intError)
	// 				fmt.Println(floatQError)
	// 				fmt.Println(floatWError)

	// 				fmt.Println("floatPValue")
	// 				fmt.Println(floatPValue)
	// 				fmt.Println("floatQValue")
	// 				fmt.Println(floatQValue)
	// 				fmt.Println("floatWValue")
	// 				fmt.Println(floatWValue)
	// 			}
	// 		}
	// 		rowIndex = rowIndex + 1
	// 	}
	// }
	//

	// instantiate the Feed Forward
	ff := &gobrain.FeedForward{}

	// initialize the Neural Network;
	// the networks structure will contain:
	// 2 inputs, 2 hidden nodes and 1 output.
	ff.Init(2, 2, 1)

	// train the network using the XOR patterns
	// the training will run for 1000 epochs
	// the learning rate is set to 0.6 and the momentum factor to 0.4
	// use true in the last parameter to receive reports about the learning error
	ff.Train(patterns, 1000, 0.6, 0.4, true)

	inputs := []float64{600, 46.76} // 51.7
	result := ff.Update(inputs)
	fmt.Println(result)
}
