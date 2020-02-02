package main

import (
    "fmt"
	"github.com/tealeg/xlsx"
	"math"
	"strconv"
	"github.com/goml/gobrain"
	"math/rand"
)

func main() {
	// set the random seed to 0
	rand.Seed(0)

	
    excelFileName := "data.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        fmt.Printf("Error")
	}
	
	// create the XOR representation patter to train the network
	patterns := [][][]float64

    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
				text := cell.String()
				if text == "Бензин АИ-92" {
					fmt.Printf("%s\n", text)
					var intValue, intError = sheet.Cell(rowIndex, 28).Int()
					var floatQValue, floatQError = sheet.Cell(rowIndex, 30).Float()
					var floatWValue, floatWError = sheet.Cell(rowIndex, 29).Float()
					patterns += float64{{intValue,floatQValue }, {floatQValue}}
				}
            }
			rowIndex = rowIndex + 1
        }
    }
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
}
