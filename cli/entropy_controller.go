package cli

import "fmt"

var allowPrintEntropy bool

func switchEntropy() {
	if allowPrintEntropy {
		allowPrintEntropy = false
	} else {
		allowPrintEntropy = true
	}
}

func tryPrintEntropy(entrType string, blocksEntropy [][]float64) {
	if allowPrintEntropy {
		println(entrType)
		for i, blockEntropy := range blocksEntropy {
			fmt.Printf("Block %d:\n", i)
			for r, entropy := range blockEntropy {
				fmt.Printf("Round %d, entropy: %0.3f\n", r, entropy)
			}
			println()
		}
	}
}
