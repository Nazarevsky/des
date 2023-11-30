package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var opSys string

// Flag to show entropy
// Test DES
// Encrypt
// Decrypt
// Exit

func Run() {
	opSys = runtime.GOOS
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter something: ")

	var errorText string
	for {
		clearConsole()
		if errorText != "" {
			println(errorText)
			errorText = ""
		}

		println("Choose the variant:")
		fmt.Printf("1. Print entropy: %v\n", allowPrintEntropy)
		fmt.Printf("2. Test DES\n")
		fmt.Printf("0. Exit\n")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			switchEntropy()
		case "2":
			desTesting()
		case "0":
			println("Have a good day!")
			return
		}
	}
}

func clearConsole() {

	if opSys == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if opSys == "linux" || opSys == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		for i := 0; i < 3; i++ {
			println()
		}
	}
}
